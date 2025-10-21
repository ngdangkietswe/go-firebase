/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package helper

import "go.uber.org/fx"

var Module = fx.Provide(
	NewUserHelper,
	NewNotificationTopicHelper,
)
