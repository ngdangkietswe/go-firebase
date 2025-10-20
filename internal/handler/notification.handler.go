/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package handler

import (
	"context"
	"go-firebase/internal/model"
	"go-firebase/internal/request"
	"go-firebase/internal/response"
	"go-firebase/internal/service"
	apiutil "go-firebase/internal/util"
	"go-firebase/pkg/constant"
	"go-firebase/pkg/util"

	"github.com/gofiber/fiber/v2"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
)

type NotificationHandler struct {
	logger          *logger.Logger
	notificationSvc service.NotificationService
}

func (h *NotificationHandler) SendNotification(c *fiber.Ctx, request *request.SendNotificationRequest) (*response.SendNotificationResponse, error) {
	newCtx, cancel := context.WithTimeout(apiutil.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, request, h.notificationSvc.SendNotification)
}

func (h *NotificationHandler) GetNotifications(c *fiber.Ctx, userID string) ([]*model.Notification, error) {
	newCtx, cancel := context.WithTimeout(apiutil.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, userID, h.notificationSvc.GetNotifications)
}

func (h *NotificationHandler) MarkNotificationAsRead(c *fiber.Ctx, notificationID string) (*response.EmptyResponse, error) {
	newCtx, cancel := context.WithTimeout(apiutil.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, notificationID, h.notificationSvc.MarkNotificationAsRead)
}

func (h *NotificationHandler) MarkAllNotificationsAsRead(c *fiber.Ctx, userID string) (*response.EmptyResponse, error) {
	newCtx, cancel := context.WithTimeout(apiutil.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, userID, h.notificationSvc.MarkAllNotificationsAsRead)
}

func NewNotificationHandler(
	logger *logger.Logger,
	notificationSvc service.NotificationService,
) *NotificationHandler {
	return &NotificationHandler{
		logger:          logger,
		notificationSvc: notificationSvc,
	}
}
