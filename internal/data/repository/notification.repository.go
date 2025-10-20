/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package repository

import (
	"context"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/data/ent/notification"
	"go-firebase/internal/request"

	"github.com/google/uuid"
)

type notificationRepo struct {
	cli *ent.Client
}

func (r *notificationRepo) Save(ctx context.Context, tx *ent.Tx, request *request.SendNotificationRequest) (*ent.Notification, error) {
	builder := tx.Notification.Create().
		SetUserID(uuid.MustParse(request.UserID)).
		SetTitle(request.Title).
		SetBody(request.Body).
		SetData(request.Payload).
		SetIsRead(false)
	return builder.Save(ctx)
}

func (r *notificationRepo) FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]*ent.Notification, error) {
	return r.cli.Notification.Query().Where(notification.UserID(userID)).Order(ent.Desc(notification.FieldSentAt)).All(ctx)
}

func (r *notificationRepo) MarkAsRead(ctx context.Context, tx *ent.Tx, notificationID uuid.UUID) error {
	return tx.Notification.UpdateOneID(notificationID).SetIsRead(true).Exec(ctx)
}

func (r *notificationRepo) MarkAllAsReadByUserID(ctx context.Context, tx *ent.Tx, userID uuid.UUID) error {
	return tx.Notification.Update().Where(notification.UserID(userID)).SetIsRead(true).Exec(ctx)
}

func (r *notificationRepo) ExistsByID(ctx context.Context, notificationID uuid.UUID) (bool, error) {
	return r.cli.Notification.Query().Where(notification.ID(notificationID)).Exist(ctx)
}

func NewNotificationRepository(cli *ent.Client) NotificationRepository {
	return &notificationRepo{cli: cli}
}
