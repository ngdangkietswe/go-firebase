/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package mapper

import "go.uber.org/fx"

var Module = fx.Provide(
	NewUserMapper,
	NewNotificationMapper,
	NewDeviceTokenMapper,
)
