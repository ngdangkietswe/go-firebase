/**
 * Author : ngdangkietswe
 * Since  : 10/21/2025
 */

package repository

import (
	"context"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/data/ent/notificationtopic"

	"github.com/google/uuid"
)

type notificationTopicRepo struct {
	cli *ent.Client
}

func (r *notificationTopicRepo) Save(ctx context.Context, tx *ent.Tx, topicName string) (*ent.NotificationTopic, error) {
	builder := tx.NotificationTopic.Create().SetName(topicName)
	return builder.Save(ctx)
}

func (r *notificationTopicRepo) FindByID(ctx context.Context, topicID uuid.UUID) (*ent.NotificationTopic, error) {
	return r.cli.NotificationTopic.Query().Where(notificationtopic.ID(topicID)).Only(ctx)
}

func (r *notificationTopicRepo) FindByName(ctx context.Context, topicName string) (*ent.NotificationTopic, error) {
	return r.cli.NotificationTopic.Query().Where(notificationtopic.Name(topicName)).Only(ctx)
}

func (r *notificationTopicRepo) FindAllByIDIn(ctx context.Context, topicIDs []uuid.UUID) ([]*ent.NotificationTopic, error) {
	return r.cli.NotificationTopic.Query().Where(notificationtopic.IDIn(topicIDs...)).All(ctx)
}

func NewNotificationTopicRepository(cli *ent.Client) NotificationTopicRepository {
	return &notificationTopicRepo{
		cli: cli,
	}
}
