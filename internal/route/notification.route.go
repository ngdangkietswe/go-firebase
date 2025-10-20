/**
 * Author : ngdangkietswe
 * Since  : 10/19/2025
 */

package route

import (
	"go-firebase/internal/controller"

	"github.com/gofiber/fiber/v2"
)

type NotificationRoute struct {
	notificationCtrl *controller.NotificationCtrl
}

func (r *NotificationRoute) Register(router fiber.Router) {
	notificationRouter := router.Group("/notifications")
	notificationRouter.Post("/", r.notificationCtrl.SendNotification)
	notificationRouter.Get("/users/:id", r.notificationCtrl.GetNotifications)
	notificationRouter.Patch("/:id/read", r.notificationCtrl.MarkNotificationAsRead)
	notificationRouter.Patch("/users/:id/read", r.notificationCtrl.MarkAllNotificationsAsRead)
}

func NewNotificationRoute(
	notificationCtrl *controller.NotificationCtrl,
) *NotificationRoute {
	return &NotificationRoute{
		notificationCtrl: notificationCtrl,
	}
}
