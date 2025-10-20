/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package route

import (
	"go-firebase/internal/controller"

	"github.com/gofiber/fiber/v2"
)

type DeviceTokenRoute struct {
	deviceTokenCtrl *controller.DeviceTokenCtrl
}

func (r *DeviceTokenRoute) Register(router fiber.Router) {
	deviceTokenRouter := router.Group("/device-tokens")
	deviceTokenRouter.Post("/", r.deviceTokenCtrl.RegisterDeviceToken)
}

func NewDeviceTokenRoute(
	deviceTokenCtrl *controller.DeviceTokenCtrl,
) *DeviceTokenRoute {
	return &DeviceTokenRoute{
		deviceTokenCtrl: deviceTokenCtrl,
	}
}
