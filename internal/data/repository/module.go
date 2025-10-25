/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package repository

import (
	"go-firebase/internal/data/datasource"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	datasource.NewEntClient,
	NewUserRepository,
	NewDeviceTokenRepository,
	NewNotificationRepository,
	NewNotificationTopicRepository,
	NewUserNotificationTopicRepository,
	NewRoleRepository,
	NewPermissionRepository,
)
