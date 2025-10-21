/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package route

import (
	"github.com/gofiber/fiber/v2"
)

type AppRoute struct {
	authRoute              *AuthRoute
	userRoute              *UserRoute
	deviceTokenRoute       *DeviceTokenRoute
	notificationRoute      *NotificationRoute
	notificationTopicRoute *NotificationTopicRoute
}

func (r *AppRoute) Register(app *fiber.App) {
	api := app.Group("/api/v1")

	r.authRoute.Register(api)
	r.userRoute.Register(api)
	r.deviceTokenRoute.Register(api)
	r.notificationRoute.Register(api)
	r.notificationTopicRoute.Register(api)
}

func NewAppRoute(
	authRoute *AuthRoute,
	userRoute *UserRoute,
	deviceTokenRoute *DeviceTokenRoute,
	notificationRoute *NotificationRoute,
	notificationTopicRoute *NotificationTopicRoute,
) *AppRoute {
	return &AppRoute{
		authRoute:              authRoute,
		userRoute:              userRoute,
		deviceTokenRoute:       deviceTokenRoute,
		notificationRoute:      notificationRoute,
		notificationTopicRoute: notificationTopicRoute,
	}
}
