/**
 * Author : ngdangkietswe
 * Since  : 10/19/2025
 */

package controller

import (
	"go-firebase/internal/handler"
	"go-firebase/internal/request"
	"go-firebase/internal/response"

	"github.com/gofiber/fiber/v2"
)

type NotificationCtrl struct {
	notificationHandler *handler.NotificationHandler
}

// SendNotification godoc
// @Summary      Send a notification
// @Description  Send a push notification to a device
// @Tags         Notification API
// @Accept       json
// @Produce      json
// @Param        notification  body      request.SendNotificationRequest  true  "Notification Info"
// @Success      200           {object}  response.ApiResponse
// @Failure      400           {object}  response.ApiResponse
// @Failure      500           {object}  response.ApiResponse
// @Router       /notifications [post]
func (c *NotificationCtrl) SendNotification(ctx *fiber.Ctx) error {
	var req *request.SendNotificationRequest

	if err := ctx.BodyParser(&req); err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusBadRequest, err)
	}

	res, err := c.notificationHandler.SendNotification(ctx, req)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusInternalServerError, err)
	}

	return response.ApiSuccessResponse(ctx, res)
}

// GetNotifications godoc
// @Summary      Get notifications for a user
// @Description  Retrieve all notifications for a specific user
// @Tags         Notification API
// @Accept       json
// @Produce      json
// @Param        user_id  path      string  true  "User ID"
// @Success      200      {object}  response.ApiResponse
// @Failure      400      {object}  response.ApiResponse
// @Failure      500      {object}  response.ApiResponse
// @Router       /notifications/users/{id} [get]
func (c *NotificationCtrl) GetNotifications(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")

	res, err := c.notificationHandler.GetNotifications(ctx, userID)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusInternalServerError, err)
	}

	return response.ApiSuccessResponse(ctx, res)
}

// MarkNotificationAsRead godoc
// @Summary      Mark a notification as read
// @Description  Mark a specific notification as read
// @Tags         Notification API
// @Accept       json
// @Produce      json
// @Param        id  			  path      string  true  "Notification ID"
// @Success      200              {object}  response.ApiResponse
// @Failure      400              {object}  response.ApiResponse
// @Failure      500              {object}  response.ApiResponse
// @Router       /notifications/{id}/read [patch]
func (c *NotificationCtrl) MarkNotificationAsRead(ctx *fiber.Ctx) error {
	notificationID := ctx.Params("id")

	res, err := c.notificationHandler.MarkNotificationAsRead(ctx, notificationID)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusInternalServerError, err)
	}

	return response.ApiSuccessResponse(ctx, res)
}

// MarkAllNotificationsAsRead godoc
// @Summary      Mark all notifications as read for a user
// @Description  Mark all notifications as read for a specific user
// @Tags         Notification API
// @Accept       json
// @Produce      json
// @Param        id  	  path      string  true  "User ID"
// @Success      200      {object}  response.ApiResponse
// @Failure      400      {object}  response.ApiResponse
// @Failure      500      {object}  response.ApiResponse
// @Router       /notifications/users/{id}/read [patch]
func (c *NotificationCtrl) MarkAllNotificationsAsRead(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")

	res, err := c.notificationHandler.MarkAllNotificationsAsRead(ctx, userID)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusInternalServerError, err)
	}

	return response.ApiSuccessResponse(ctx, res)
}

func NewNotificationCtrl(
	notificationHandler *handler.NotificationHandler,
) *NotificationCtrl {
	return &NotificationCtrl{
		notificationHandler: notificationHandler,
	}
}
