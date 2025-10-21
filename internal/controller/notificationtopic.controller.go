/**
 * Author : ngdangkietswe
 * Since  : 10/21/2025
 */

package controller

import (
	"go-firebase/internal/handler"
	"go-firebase/pkg/request"
	"go-firebase/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type NotificationTopicCtrl struct {
	notificationTopicHandler *handler.NotificationTopicHandler
}

// CreateNotificationTopic godoc
// @Summary      Create a notification topic
// @Description  Create a new notification topic
// @Tags         Notification Topic API
// @Accept       json
// @Produce      json
// @Param        topic  body      request.CreateNotificationTopicRequest  true  "Notification Topic Info"
// @Success      200           {object}  response.ApiResponse
// @Failure      400           {object}  response.ApiResponse
// @Failure      500           {object}  response.ApiResponse
// @Router       /notification-topics [post]
func (c *NotificationTopicCtrl) CreateNotificationTopic(ctx *fiber.Ctx) error {
	var req *request.CreateNotificationTopicRequest

	if err := ctx.BodyParser(&req); err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusBadRequest, err)
	}

	res, err := c.notificationTopicHandler.CreateNotificationTopic(ctx, req)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusInternalServerError, err)
	}

	return response.ApiSuccessResponse(ctx, res)
}

// SubscribeToNotificationTopic godoc
// @Summary      Subscribe to a notification topic
// @Description  Subscribe a device to a notification topic
// @Tags         Notification Topic API
// @Accept       json
// @Produce      json
// @Param        subscription  body      request.SubscribeNotificationTopicRequest  true  "Subscription Info"
// @Success      200           {object}  response.ApiResponse
// @Failure      400           {object}  response.ApiResponse
// @Failure      500           {object}  response.ApiResponse
// @Router       /notification-topics/subscribe [post]
func (c *NotificationTopicCtrl) SubscribeToNotificationTopic(ctx *fiber.Ctx) error {
	var req *request.SubscribeNotificationTopicRequest

	if err := ctx.BodyParser(&req); err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusBadRequest, err)
	}

	res, err := c.notificationTopicHandler.SubscribeToNotificationTopic(ctx, req)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusInternalServerError, err)
	}

	return response.ApiSuccessResponse(ctx, res)
}

func NewNotificationTopicCtrl(
	notificationTopicHandler *handler.NotificationTopicHandler,
) *NotificationTopicCtrl {
	return &NotificationTopicCtrl{
		notificationTopicHandler: notificationTopicHandler,
	}
}
