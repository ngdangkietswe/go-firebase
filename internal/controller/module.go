/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package controller

import "go.uber.org/fx"

var Module = fx.Provide(
	NewAuthCtrl,
	NewUserCtrl,
	NewDeviceTokenCtrl,
	NewNotificationCtrl,
)
