/**
 * Author : ngdangkietswe
 * Since  : 10/21/2025
 */

package helper

import (
	"context"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/data/repository"
	"go-firebase/internal/firebase"

	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type notificationTopicHelper struct {
	logger                *logger.Logger
	fcmCli                firebase.FCMClient
	notificationTopicRepo repository.NotificationTopicRepository
	deviceTokenRepo       repository.DeviceTokenRepository
}

func (h *notificationTopicHelper) FirebaseSubscribeToTopic(ctx context.Context, userID uuid.UUID, topicIDs []uuid.UUID) error {
	tokens, err := h.getTokensByUserID(ctx, userID)
	if err != nil {
		return err
	}

	topics, err := h.getTopicsByIDs(ctx, topicIDs)
	if err != nil {
		return err
	}

	for _, topicName := range topics {
		if firebaseErr := h.fcmCli.SubscribeToTopic(tokens, topicName); firebaseErr != nil {
			h.logger.Error("Failed to subscribe to topic", zap.String("topic", topicName), zap.Error(firebaseErr))
		}
	}

	return nil
}

func (h *notificationTopicHelper) FirebaseUnsubscribeFromTopic(ctx context.Context, userID uuid.UUID, topicIDs []uuid.UUID) error {
	tokens, err := h.getTokensByUserID(ctx, userID)
	if err != nil {
		return err
	}

	topics, err := h.getTopicsByIDs(ctx, topicIDs)
	if err != nil {
		return err
	}

	for _, topicName := range topics {
		if firebaseErr := h.fcmCli.UnsubscribeFromTopic(tokens, topicName); firebaseErr != nil {
			h.logger.Error("Failed to subscribe to topic", zap.String("topic", topicName), zap.Error(firebaseErr))
		}
	}

	return nil
}

func (h *notificationTopicHelper) getTokensByUserID(ctx context.Context, userID uuid.UUID) ([]string, error) {
	deviceTokens, err := h.deviceTokenRepo.FindAllByUserID(ctx, userID)
	if err != nil {
		h.logger.Error("Failed to get device tokens by user ID", zap.String("user_id", userID.String()), zap.Error(err))
		return nil, err
	}
	return lo.Map(deviceTokens, func(item *ent.DeviceToken, _ int) string {
		return item.Token
	}), nil
}

func (h *notificationTopicHelper) getTopicsByIDs(ctx context.Context, topicIDs []uuid.UUID) ([]string, error) {
	topics, err := h.notificationTopicRepo.FindAllByIDIn(ctx, topicIDs)
	if err != nil {
		h.logger.Error("Failed to get topics by IDs", zap.Any("topic_ids", topicIDs), zap.Error(err))
		return nil, err
	}
	return lo.Map(topics, func(item *ent.NotificationTopic, _ int) string {
		return item.Name
	}), nil
}

func NewNotificationTopicHelper(
	logger *logger.Logger,
	fcmCli firebase.FCMClient,
	notificationTopicRepo repository.NotificationTopicRepository,
	deviceTokenRepo repository.DeviceTokenRepository,
) NotificationTopicHelper {
	return &notificationTopicHelper{
		logger:                logger,
		fcmCli:                fcmCli,
		notificationTopicRepo: notificationTopicRepo,
		deviceTokenRepo:       deviceTokenRepo,
	}
}
