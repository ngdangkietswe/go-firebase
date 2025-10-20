/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package mapper

import (
	"go-firebase/internal/data/ent"
	"go-firebase/internal/model"

	"github.com/samber/lo"

	"github.com/ngdangkietswe/swe-go-common-shared/util"
)

type notificationMapper struct {
}

func (m *notificationMapper) AsMono(notification *ent.Notification) *model.Notification {
	builder := &model.Notification{
		NotificationID: notification.ID.String(),
		UserID:         notification.UserID.String(),
		IsRead:         notification.IsRead,
		SentAt:         util.Format(lo.ToPtr(notification.SentAt), util.LayoutISOWithTime),
	}

	if notification.Title != "" {
		builder.Title = notification.Title
	}

	if notification.Body != "" {
		builder.Body = notification.Body
	}

	return builder
}

func (m *notificationMapper) AsList(notifications []*ent.Notification) []*model.Notification {
	if len(notifications) == 0 {
		return []*model.Notification{}
	}

	builders := make([]*model.Notification, 0, len(notifications))

	lo.ForEach(notifications, func(notification *ent.Notification, _ int) {
		builders = append(builders, m.AsMono(notification))
	})

	return builders
}

func NewNotificationMapper() NotificationMapper {
	return &notificationMapper{}
}
