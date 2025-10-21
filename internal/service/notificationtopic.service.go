/**
 * Author : ngdangkietswe
 * Since  : 10/21/2025
 */

package service

import (
	"context"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/data/repository"
	"go-firebase/internal/firebase"
	"go-firebase/internal/helper"
	"go-firebase/pkg/constant"
	"go-firebase/pkg/request"
	"go-firebase/pkg/response"
	"go-firebase/pkg/util"

	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type notificationTopicSvc struct {
	cli                       *ent.Client
	logger                    *logger.Logger
	fcmCli                    firebase.FCMClient
	notificationTopicRepo     repository.NotificationTopicRepository
	userNotificationTopicRepo repository.UserNotificationTopicRepository
	deviceTokenRepo           repository.DeviceTokenRepository
	notificationTopicHelper   helper.NotificationTopicHelper
}

func (s *notificationTopicSvc) CreateNotificationTopic(ctx context.Context, request *request.CreateNotificationTopicRequest) (*response.IdResponse, error) {
	topic, err := repository.WithTxResult(ctx, s.cli, s.logger, func(tx *ent.Tx) (*ent.NotificationTopic, error) {
		return s.notificationTopicRepo.Save(ctx, tx, request.TopicName)
	})

	if err != nil {
		s.logger.Error("Failed to create notification topic", zap.Error(err))
		return nil, err
	}

	return &response.IdResponse{ID: topic.ID.String()}, nil
}

func (s *notificationTopicSvc) SubscribeNotificationTopic(ctx context.Context, request *request.SubscribeNotificationTopicRequest) (*response.EmptyResponse, error) {
	userID := uuid.MustParse(ctx.Value(constant.CtxSysUIDKey).(string))
	topicIDs := lo.Map(request.TopicIDs, func(id string, _ int) uuid.UUID {
		return uuid.MustParse(id)
	})

	if err := repository.WithTx(ctx, s.cli, s.logger, func(tx *ent.Tx) error {
		return s.userNotificationTopicRepo.SaveAll(ctx, tx, userID, topicIDs)
	}); err != nil {
		s.logger.Error("Failed to subscribe notification topics", zap.String("user_id", userID.String()), zap.Error(err))
		return nil, err
	}

	go func(userID uuid.UUID, topicIDs []uuid.UUID) {
		defer util.RecoverPanic()
		if err := s.notificationTopicHelper.FirebaseSubscribeToTopic(context.Background(), userID, topicIDs); err != nil {
			s.logger.Error("Failed to subscribe to notification topics in Firebase", zap.String("user_id", userID.String()), zap.Error(err))
		}
	}(userID, topicIDs)

	return &response.EmptyResponse{}, nil
}

func (s *notificationTopicSvc) UnsubscribeNotificationTopic(ctx context.Context, request *request.SubscribeNotificationTopicRequest) (*response.EmptyResponse, error) {
	userID := uuid.MustParse(ctx.Value(constant.CtxSysUIDKey).(string))
	topicIDs := lo.Map(request.TopicIDs, func(id string, _ int) uuid.UUID {
		return uuid.MustParse(id)
	})

	if err := repository.WithTx(ctx, s.cli, s.logger, func(tx *ent.Tx) error {
		return s.userNotificationTopicRepo.DeleteByUserIDAndTopicIDIn(ctx, tx, userID, topicIDs)
	}); err != nil {
		s.logger.Error("Failed to unsubscribe notification topics", zap.String("user_id", userID.String()), zap.Error(err))
		return nil, err
	}

	go func(userID uuid.UUID, topicIDs []uuid.UUID) {
		defer util.RecoverPanic()
		if err := s.notificationTopicHelper.FirebaseUnsubscribeFromTopic(context.Background(), userID, topicIDs); err != nil {
			s.logger.Error("Failed to unsubscribe from notification topics in Firebase", zap.String("user_id", userID.String()), zap.Error(err))
		}
	}(userID, topicIDs)

	return &response.EmptyResponse{}, nil
}

func NewNotificationTopicService(
	cli *ent.Client,
	logger *logger.Logger,
	fcmCli firebase.FCMClient,
	notificationTopicRepo repository.NotificationTopicRepository,
	userNotificationTopicRepo repository.UserNotificationTopicRepository,
	deviceTokenRepo repository.DeviceTokenRepository,
	notificationTopicHelper helper.NotificationTopicHelper,
) NotificationTopicService {
	return &notificationTopicSvc{
		cli:                       cli,
		logger:                    logger,
		fcmCli:                    fcmCli,
		notificationTopicRepo:     notificationTopicRepo,
		userNotificationTopicRepo: userNotificationTopicRepo,
		deviceTokenRepo:           deviceTokenRepo,
		notificationTopicHelper:   notificationTopicHelper,
	}
}
