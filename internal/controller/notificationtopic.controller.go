/**
 * Author : ngdangkietswe
 * Since  : 10/21/2025
 */

package controller

import (
	"go-firebase/internal/handler"
	"go-firebase/pkg/request"
	"go-firebase/pkg/response"
	"go-firebase/pkg/util"

	"github.com/gofiber/fiber/v2"
)

type NotificationTopicCtrl struct {
	notificationTopicHandler *handler.NotificationTopicHandler
}

// GetNotificationTopics godoc
// @Summary      Get notification topics
// @Description  Retrieve a list of notification topics
// @Tags         Notification Topic API
// @Accept       json
// @Produce      json
// @Param        page       query     int     false  "Page number"
// @Param        page_size  query     int     false  "Number of items per page"
// @Param        sort       query     string  false  "Sort by field"
// @Param        order      query     string  false  "Order (asc or desc)"
// @Param        search     query     string  false  "Search term"
// @Success      200        {object}  response.ApiResponse
// @Failure      400        {object}  response.ApiResponse
// @Failure      500        {object}  response.ApiResponse
// @Router       /notification-topics [get]
func (c *NotificationTopicCtrl) GetNotificationTopics(ctx *fiber.Ctx) error {
	req := &request.ListNotificationTopicRequest{
		Paginate: util.AsPaginateRequest(ctx),
		Search:   ctx.Query("search"),
	}

	res, err := c.notificationTopicHandler.GetNotificationTopics(ctx, req)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusInternalServerError, err)
	}

	return response.ApiSuccessResponse(ctx, res)
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

// UnsubscribeFromNotificationTopic godoc
// @Summary      Unsubscribe from a notification topic
// @Description  Unsubscribe a device from a notification topic
// @Tags         Notification Topic API
// @Accept       json
// @Produce      json
// @Param        subscription  body      request.SubscribeNotificationTopicRequest  true  "Subscription Info"
// @Success      200           {object}  response.ApiResponse
// @Failure      400           {object}  response.ApiResponse
// @Failure      500           {object}  response.ApiResponse
// @Router       /notification-topics/unsubscribe [post]
func (c *NotificationTopicCtrl) UnsubscribeFromNotificationTopic(ctx *fiber.Ctx) error {
	var req *request.SubscribeNotificationTopicRequest

	if err := ctx.BodyParser(&req); err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusBadRequest, err)
	}

	res, err := c.notificationTopicHandler.UnsubscribeFromNotificationTopic(ctx, req)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusInternalServerError, err)
	}

	return response.ApiSuccessResponse(ctx, res)
}

func NewNotificationTopicController(
	notificationTopicHandler *handler.NotificationTopicHandler,
) *NotificationTopicCtrl {
	return &NotificationTopicCtrl{
		notificationTopicHandler: notificationTopicHandler,
	}
}
