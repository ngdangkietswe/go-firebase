/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package route

import "go.uber.org/fx"

var Module = fx.Provide(
	NewAuthRoute,
	NewUserRoute,
	NewDeviceTokenRoute,
	NewNotificationRoute,
	NewAppRoute,
)
