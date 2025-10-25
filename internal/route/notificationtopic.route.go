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
	notificationTopicRouter.Get("/", r.notificationTopicCtrl.GetNotificationTopics)
	notificationTopicRouter.Post("/", r.notificationTopicCtrl.CreateNotificationTopic)
	notificationTopicRouter.Post("/subscribe", r.notificationTopicCtrl.SubscribeToNotificationTopic)
	notificationTopicRouter.Post("/unsubscribe", r.notificationTopicCtrl.UnsubscribeFromNotificationTopic)
}

func NewNotificationTopicRoute(
	notificationTopicCtrl *controller.NotificationTopicCtrl,
) *NotificationTopicRoute {
	return &NotificationTopicRoute{
		notificationTopicCtrl: notificationTopicCtrl,
	}
}
