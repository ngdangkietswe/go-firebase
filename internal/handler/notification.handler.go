/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package handler

import (
	"context"
	"go-firebase/internal/service"
	"go-firebase/pkg/constant"
	"go-firebase/pkg/request"
	"go-firebase/pkg/response"
	"go-firebase/pkg/util"

	"github.com/gofiber/fiber/v2"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
)

type NotificationHandler struct {
	logger          *logger.Logger
	notificationSvc service.NotificationService
}

func (h *NotificationHandler) SendNotification(c *fiber.Ctx, request *request.SendNotificationRequest) (*response.SendNotificationResponse, error) {
	newCtx, cancel := context.WithTimeout(util.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, request, h.notificationSvc.SendNotification)
}

func (h *NotificationHandler) GetNotifications(c *fiber.Ctx, request *request.ListNotificationRequest) (*response.ListResponse, error) {
	newCtx, cancel := context.WithTimeout(util.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, request, h.notificationSvc.GetNotifications)
}

func (h *NotificationHandler) MarkNotificationAsRead(c *fiber.Ctx, notificationID string) (*response.EmptyResponse, error) {
	newCtx, cancel := context.WithTimeout(util.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, notificationID, h.notificationSvc.MarkNotificationAsRead)
}

func (h *NotificationHandler) MarkAllNotificationsAsRead(c *fiber.Ctx) (*response.EmptyResponse, error) {
	newCtx, cancel := context.WithTimeout(util.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFuncNoReq(newCtx, h.notificationSvc.MarkAllNotificationsAsRead)
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
