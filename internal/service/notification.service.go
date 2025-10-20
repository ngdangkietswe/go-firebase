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
	"go-firebase/internal/model"
	"go-firebase/internal/request"
	"go-firebase/internal/response"

	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type notificationSvc struct {
	cli                *ent.Client
	logger             *logger.Logger
	fcmCli             firebase.FCMClient
	notificationRepo   repository.NotificationRepository
	deviceTokenRepo    repository.DeviceTokenRepository
	userRepo           repository.UserRepository
	notificationMapper mapper.NotificationMapper
}

func (s *notificationSvc) SendNotification(ctx context.Context, request *request.SendNotificationRequest) (*response.SendNotificationResponse, error) {
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

	go func(deviceTokens []*ent.DeviceToken) {
		tokens := lo.Map(deviceTokens, func(item *ent.DeviceToken, _ int) string {
			return item.Token
		})

		firebaseErr := s.fcmCli.SendToTokens(tokens, request.Title, request.Body, request.Payload)
		if firebaseErr != nil {
			s.logger.Error("Failed to send notification via Firebase", zap.String("user_id", request.UserID), zap.String("error", firebaseErr.Error()))
		}
	}(deviceTokens)

	return &response.SendNotificationResponse{
		UserID:         user.ID.String(),
		NotificationID: notification.ID.String(),
	}, nil
}

func (s *notificationSvc) GetNotifications(ctx context.Context, userID string) ([]*model.Notification, error) {
	exists, err := s.userRepo.ExistsByID(ctx, uuid.MustParse(userID))
	if err != nil {
		s.logger.Error("Failed to check user existence by ID", zap.String("user_id", userID), zap.String("error", err.Error()))
		return nil, err
	}

	if !exists {
		s.logger.Info("User not found by ID", zap.String("user_id", userID))
		return nil, errors.New("user not found")
	}

	notifications, err := s.notificationRepo.FindAllByUserID(ctx, uuid.MustParse(userID))
	if err != nil {
		s.logger.Error("Failed to find notifications by user ID", zap.String("user_id", userID), zap.String("error", err.Error()))
		return nil, err
	}

	return s.notificationMapper.AsList(notifications), nil
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

func (s *notificationSvc) MarkAllNotificationsAsRead(ctx context.Context, userID string) (*response.EmptyResponse, error) {
	exists, err := s.userRepo.ExistsByID(ctx, uuid.MustParse(userID))
	if err != nil {
		s.logger.Error("Failed to check user existence by ID", zap.String("user_id", userID), zap.String("error", err.Error()))
		return nil, err
	}

	if !exists {
		s.logger.Info("User not found by ID", zap.String("user_id", userID))
		return nil, errors.New("user not found")
	}

	if err = repository.WithTx(ctx, s.cli, s.logger, func(tx *ent.Tx) error {
		return s.notificationRepo.MarkAllAsReadByUserID(ctx, tx, uuid.MustParse(userID))
	}); err != nil {
		s.logger.Error("Failed to mark all notifications as read", zap.String("user_id", userID), zap.String("error", err.Error()))
		return nil, err
	}

	return &response.EmptyResponse{}, nil
}

func NewNotificationService(
	cli *ent.Client,
	logger *logger.Logger,
	fcmCli firebase.FCMClient,
	notificationRepo repository.NotificationRepository,
	deviceTokenRepo repository.DeviceTokenRepository,
	userRepo repository.UserRepository,
	notificationMapper mapper.NotificationMapper,
) NotificationService {
	return &notificationSvc{
		cli:                cli,
		logger:             logger,
		fcmCli:             fcmCli,
		notificationRepo:   notificationRepo,
		deviceTokenRepo:    deviceTokenRepo,
		userRepo:           userRepo,
		notificationMapper: notificationMapper,
	}
}
