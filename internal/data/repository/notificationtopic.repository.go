/**
 * Author : ngdangkietswe
 * Since  : 10/21/2025
 */

package repository

import (
	"context"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/data/ent/notificationtopic"
	"go-firebase/pkg/request"
	"go-firebase/pkg/util"

	"github.com/google/uuid"
)

type notificationTopicRepo struct {
	cli *ent.Client
}

func (r *notificationTopicRepo) Save(ctx context.Context, tx *ent.Tx, topicName string) (*ent.NotificationTopic, error) {
	builder := tx.NotificationTopic.Create().SetName(topicName)
	return builder.Save(ctx)
}

func (r *notificationTopicRepo) FindAll(ctx context.Context, request *request.ListNotificationTopicRequest) ([]*ent.NotificationTopic, int, error) {
	query := r.cli.NotificationTopic.Query()

	if request.Search != "" {
		query = query.Where(notificationtopic.NameContainsFold(request.Search))
	}

	totalItems, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	items, err := query.Clone().Order(util.ToSortOrder(request.Paginate)).
		Offset(request.Paginate.Page * request.Paginate.PageSize).
		Limit(request.Paginate.PageSize).
		All(ctx)

	return items, totalItems, err
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
