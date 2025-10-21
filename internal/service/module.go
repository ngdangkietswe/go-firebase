/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package service

import "go.uber.org/fx"

var Module = fx.Provide(
	NewAuthService,
	NewUserService,
	NewDeviceTokenService,
	NewNotificationService,
	NewNotificationTopicService,
)
