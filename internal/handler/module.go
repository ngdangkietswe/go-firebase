/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package handler

import "go.uber.org/fx"

var Module = fx.Provide(
	NewAuthHandler,
	NewUserHandler,
	NewDeviceTokenHandler,
	NewNotificationHandler,
	NewNotificationTopicHandler,
)
