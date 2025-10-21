/**
 * Author : ngdangkietswe
 * Since  : 10/21/2025
 */

package route

import (
	"go-firebase/internal/controller"

	"github.com/gofiber/fiber/v2"
)

type NotificationTopicRoute struct {
	notificationTopicCtrl *controller.NotificationTopicCtrl
}

func (r *NotificationTopicRoute) Register(router fiber.Router) {
	notificationTopicRouter := router.Group("/notification-topics")
	notificationTopicRouter.Post("/", r.notificationTopicCtrl.CreateNotificationTopic)
	notificationTopicRouter.Post("/subscribe", r.notificationTopicCtrl.SubscribeToNotificationTopic)
}

func NewNotificationTopicRoute(
	notificationTopicCtrl *controller.NotificationTopicCtrl,
) *NotificationTopicRoute {
	return &NotificationTopicRoute{
		notificationTopicCtrl: notificationTopicCtrl,
	}
}
