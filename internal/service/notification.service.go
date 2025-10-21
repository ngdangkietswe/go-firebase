/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package service

import (
	"context"
	"errors"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/data/repository"
	"go-firebase/internal/firebase"
	"go-firebase/internal/mapper"
	"go-firebase/pkg/request"
	"go-firebase/pkg/response"
	"go-firebase/pkg/util"

	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type notificationSvc struct {
	cli                   *ent.Client
	logger                *logger.Logger
	fcmCli                firebase.FCMClient
	notificationRepo      repository.NotificationRepository
	notificationTopicRepo repository.NotificationTopicRepository
	deviceTokenRepo       repository.DeviceTokenRepository
	userRepo              repository.UserRepository
	notificationMapper    mapper.NotificationMapper
}

func (s *notificationSvc) SendNotification(ctx context.Context, request *request.SendNotificationRequest) (*response.SendNotificationResponse, error) {
	if request.UserID != "" {
		return s.sendNotificationToUser(ctx, request)
	} else if request.TopicID != "" || request.TopicName != "" {
		return s.sendNotificationToTopic(ctx, request)
	}
	return nil, errors.New("either user_id or topic_id/topic_name must be provided")
}

func (s *notificationSvc) sendNotificationToUser(ctx context.Context, request *request.SendNotificationRequest) (*response.SendNotificationResponse, error) {
	user, err := s.userRepo.FindByID(ctx, uuid.MustParse(request.UserID))
	if err != nil {
		s.logger.Error("Failed to find user by ID", zap.String("user_id", request.UserID), zap.String("error", err.Error()))
		return nil, err
	}

	deviceTokens, err := s.deviceTokenRepo.FindAllByUserID(ctx, user.ID)
	if err != nil {
		s.logger.Error("Failed to find device tokens by user ID", zap.String("user_id", request.UserID), zap.String("error", err.Error()))
		return nil, err
	}

	if len(deviceTokens) == 0 {
		s.logger.Info("No device tokens found for user", zap.String("user_id", request.UserID))
		return nil, nil
	}

	notification, err := repository.WithTxResult(ctx, s.cli, s.logger, func(tx *ent.Tx) (*ent.Notification, error) {
		return s.notificationRepo.Save(ctx, tx, request)
	})
	if err != nil {
		s.logger.Error("Failed to save notification", zap.String("user_id", request.UserID), zap.String("error", err.Error()))
		return nil, err
	}

	go func(deviceTokens []*ent.DeviceToken, title, body string, payload map[string]string) {
		tokens := lo.Map(deviceTokens, func(item *ent.DeviceToken, _ int) string {
			return item.Token
		})
		if firebaseErr := s.fcmCli.SendToTokens(tokens, title, body, payload); firebaseErr != nil {
			s.logger.Error("Failed to send notification to user", zap.String("user_id", request.UserID), zap.String("error", firebaseErr.Error()))
		}
	}(deviceTokens, request.Title, request.Body, request.Payload)

	return &response.SendNotificationResponse{
		NotificationID: notification.ID.String(),
		UserID:         user.ID.String(),
	}, nil
}

func (s *notificationSvc) sendNotificationToTopic(ctx context.Context, request *request.SendNotificationRequest) (*response.SendNotificationResponse, error) {
	var topic *ent.NotificationTopic
	var err error

	if request.TopicID != "" {
		topic, err = s.notificationTopicRepo.FindByID(ctx, uuid.MustParse(request.TopicID))
		if err != nil {
			s.logger.Error("Failed to find notification topic by ID", zap.String("topic_id", request.TopicID), zap.String("error", err.Error()))
			return nil, err
		}
	} else {
		topic, err = s.notificationTopicRepo.FindByName(ctx, request.TopicName)
		if err != nil {
			s.logger.Error("Failed to find notification topic by name", zap.String("topic_name", request.TopicName), zap.String("error", err.Error()))
			return nil, err
		}
		request.TopicID = topic.ID.String()
	}

	notification, err := repository.WithTxResult(ctx, s.cli, s.logger, func(tx *ent.Tx) (*ent.Notification, error) {
		return s.notificationRepo.Save(ctx, tx, request)
	})
	if err != nil {
		s.logger.Error("Failed to save notification", zap.String("topic_id", request.TopicID), zap.String("error", err.Error()))
		return nil, err
	}

	go func(topicName, title, body string, payload map[string]string) {
		if firebaseErr := s.fcmCli.SendToTopic(topicName, title, body, payload); firebaseErr != nil {
			s.logger.Error("Failed to send notification to topic", zap.String("topic_id", request.TopicID), zap.String("error", firebaseErr.Error()))
		}
	}(topic.Name, request.Title, request.Body, request.Payload)

	return &response.SendNotificationResponse{
		NotificationID: notification.ID.String(),
		TopicID:        topic.ID.String(),
	}, nil
}

func (s *notificationSvc) GetNotifications(ctx context.Context, request *request.ListNotificationRequest) (*response.ListResponse, error) {
	util.NormalizePaginationRequest(request.Paginate)

	items, totalItems, err := s.notificationRepo.FindAll(ctx, request)
	if err != nil {
		s.logger.Error("Failed to get notifications", zap.String("error", err.Error()))
		return nil, err
	}

	return &response.ListResponse{
		Items: s.notificationMapper.AsList(items),
		Meta:  util.AsPageMeta(request.Paginate, totalItems),
	}, nil
}

func (s *notificationSvc) MarkNotificationAsRead(ctx context.Context, notificationID string) (*response.EmptyResponse, error) {
	exists, err := s.notificationRepo.ExistsByID(ctx, uuid.MustParse(notificationID))
	if err != nil {
		s.logger.Error("Failed to check notification existence by ID", zap.String("notification_id", notificationID), zap.String("error", err.Error()))
		return nil, err
	}

	if !exists {
		s.logger.Info("Notification not found by ID", zap.String("notification_id", notificationID))
		return nil, errors.New("notification not found")
	}

	if err = repository.WithTx(ctx, s.cli, s.logger, func(tx *ent.Tx) error {
		return s.notificationRepo.MarkAsRead(ctx, tx, uuid.MustParse(notificationID))
	}); err != nil {
		s.logger.Error("Failed to mark notification as read", zap.String("notification_id", notificationID), zap.String("error", err.Error()))
		return nil, err
	}

	return &response.EmptyResponse{}, nil
}

func (s *notificationSvc) MarkAllNotificationsAsRead(ctx context.Context) (*response.EmptyResponse, error) {
	if err := repository.WithTx(ctx, s.cli, s.logger, func(tx *ent.Tx) error {
		return s.notificationRepo.MarkAllAsRead(ctx, tx)
	}); err != nil {
		s.logger.Error("Failed to mark all notifications as read", zap.Error(err))
		return nil, err
	}

	return &response.EmptyResponse{}, nil
}

func NewNotificationService(
	cli *ent.Client,
	logger *logger.Logger,
	fcmCli firebase.FCMClient,
	notificationRepo repository.NotificationRepository,
	notificationTopicRepo repository.NotificationTopicRepository,
	deviceTokenRepo repository.DeviceTokenRepository,
	userRepo repository.UserRepository,
	notificationMapper mapper.NotificationMapper,
) NotificationService {
	return &notificationSvc{
		cli:                   cli,
		logger:                logger,
		fcmCli:                fcmCli,
		notificationRepo:      notificationRepo,
		notificationTopicRepo: notificationTopicRepo,
		deviceTokenRepo:       deviceTokenRepo,
		userRepo:              userRepo,
		notificationMapper:    notificationMapper,
	}
}
