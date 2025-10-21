/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package repository

import (
	"context"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/data/ent/notification"
	"go-firebase/internal/data/ent/user"
	"go-firebase/pkg/constant"
	"go-firebase/pkg/request"
	"go-firebase/pkg/util"

	"github.com/google/uuid"
)

type notificationRepo struct {
	cli *ent.Client
}

func (r *notificationRepo) Save(ctx context.Context, tx *ent.Tx, request *request.SendNotificationRequest) (*ent.Notification, error) {
	builder := tx.Notification.Create().
		SetTitle(request.Title).
		SetBody(request.Body).
		SetData(request.Payload).
		SetIsRead(false)

	if request.UserID != "" {
		builder.SetUserID(uuid.MustParse(request.UserID))
	} else {
		builder.SetNotificationTopicID(uuid.MustParse(request.TopicID))
	}

	return builder.Save(ctx)
}

func (r *notificationRepo) FindAll(ctx context.Context, request *request.ListNotificationRequest) ([]*ent.Notification, int, error) {
	firebaseUID := ctx.Value(constant.CtxFirebaseUIDKey).(string)

	query := r.cli.Notification.Query().Where(notification.HasUserWith(user.FirebaseUID(firebaseUID)))

	if request.IsRead != nil {
		query = query.Where(notification.IsRead(*request.IsRead))
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

func (r *notificationRepo) MarkAsRead(ctx context.Context, tx *ent.Tx, notificationID uuid.UUID) error {
	return tx.Notification.UpdateOneID(notificationID).SetIsRead(true).Exec(ctx)
}

func (r *notificationRepo) MarkAllAsRead(ctx context.Context, tx *ent.Tx) error {
	firebaseUID := ctx.Value(constant.CtxFirebaseUIDKey).(string)
	return tx.Notification.Update().
		Where(notification.HasUserWith(user.FirebaseUID(firebaseUID))).
		SetIsRead(true).
		Exec(ctx)
}

func (r *notificationRepo) ExistsByID(ctx context.Context, notificationID uuid.UUID) (bool, error) {
	firebaseUID := ctx.Value(constant.CtxFirebaseUIDKey).(string)
	return r.cli.Notification.Query().
		Where(
			notification.ID(notificationID),
			notification.HasUserWith(user.FirebaseUID(firebaseUID)),
		).
		Exist(ctx)
}

func NewNotificationRepository(cli *ent.Client) NotificationRepository {
	return &notificationRepo{cli: cli}
}
