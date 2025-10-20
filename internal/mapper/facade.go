/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package mapper

import (
	"go-firebase/internal/data/ent"
	"go-firebase/internal/model"
)

type (
	UserMapper interface {
		AsMono(user *ent.User) *model.User
		AsList(users []*ent.User) []*model.User
	}

	NotificationMapper interface {
		AsMono(notification *ent.Notification) *model.Notification
		AsList(notifications []*ent.Notification) []*model.Notification
	}

	DeviceTokenMapper interface {
		AsMono(deviceToken *ent.DeviceToken) *model.Device
		AsList(deviceTokens []*ent.DeviceToken) []*model.Device
	}
)
