/**
 * Author : ngdangkietswe
 * Since  : 10/21/2025
 */

package repository

import (
	"context"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/data/ent/usernotificationtopic"

	"github.com/google/uuid"
)

type userNotificationTopicRepo struct {
	cli *ent.Client
}

func (r *userNotificationTopicRepo) Save(ctx context.Context, tx *ent.Tx, userID uuid.UUID, topicID uuid.UUID) (*ent.UserNotificationTopic, error) {
	builder := tx.UserNotificationTopic.Create().SetUserID(userID).SetNotificationTopicID(topicID)
	return builder.Save(ctx)
}

func (r *userNotificationTopicRepo) SaveAll(ctx context.Context, tx *ent.Tx, userID uuid.UUID, topicIDs []uuid.UUID) error {
	builders := make([]*ent.UserNotificationTopicCreate, 0, len(topicIDs))
	for _, topicID := range topicIDs {
		builder := tx.UserNotificationTopic.Create().SetUserID(userID).SetNotificationTopicID(topicID)
		builders = append(builders, builder)
	}
	return tx.UserNotificationTopic.CreateBulk(builders...).Exec(ctx)
}

func (r *userNotificationTopicRepo) ExistsByUserIDAndTopicID(ctx context.Context, userID uuid.UUID, topicID uuid.UUID) (bool, error) {
	return r.cli.UserNotificationTopic.Query().
		Where(
			usernotificationtopic.UserID(userID),
			usernotificationtopic.NotificationTopicID(topicID),
		).
		Exist(ctx)
}

func (r *userNotificationTopicRepo) DeleteByUserIDAndTopicID(ctx context.Context, tx *ent.Tx, userID uuid.UUID, topicID uuid.UUID) error {
	_, err := tx.UserNotificationTopic.Delete().
		Where(
			usernotificationtopic.UserID(userID),
			usernotificationtopic.NotificationTopicID(topicID),
		).
		Exec(ctx)
	return err
}

func (r *userNotificationTopicRepo) DeleteByUserIDAndTopicIDIn(ctx context.Context, tx *ent.Tx, userID uuid.UUID, topicIDs []uuid.UUID) error {
	_, err := tx.UserNotificationTopic.Delete().
		Where(
			usernotificationtopic.UserID(userID),
			usernotificationtopic.NotificationTopicIDIn(topicIDs...),
		).
		Exec(ctx)
	return err
}

func NewUserNotificationTopicRepository(cli *ent.Client) UserNotificationTopicRepository {
	return &userNotificationTopicRepo{
		cli: cli,
	}
}
