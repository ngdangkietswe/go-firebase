/**
 * Author : ngdangkietswe
 * Since  : 10/21/2025
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

type NotificationTopicHandler struct {
	logger               *logger.Logger
	notificationTopicSvc service.NotificationTopicService
}

func (h *NotificationTopicHandler) GetNotificationTopics(c *fiber.Ctx, request *request.ListNotificationTopicRequest) (*response.ListResponse, error) {
	newCtx, cancel := context.WithTimeout(util.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, request, h.notificationTopicSvc.GetNotificationTopics)
}

func (h *NotificationTopicHandler) CreateNotificationTopic(c *fiber.Ctx, request *request.CreateNotificationTopicRequest) (*response.IdResponse, error) {
	newCtx, cancel := context.WithTimeout(util.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, request, h.notificationTopicSvc.CreateNotificationTopic)
}

func (h *NotificationTopicHandler) SubscribeToNotificationTopic(c *fiber.Ctx, request *request.SubscribeNotificationTopicRequest) (*response.EmptyResponse, error) {
	newCtx, cancel := context.WithTimeout(util.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, request, h.notificationTopicSvc.SubscribeNotificationTopic)
}

func (h *NotificationTopicHandler) UnsubscribeFromNotificationTopic(c *fiber.Ctx, request *request.SubscribeNotificationTopicRequest) (*response.EmptyResponse, error) {
	newCtx, cancel := context.WithTimeout(util.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, request, h.notificationTopicSvc.UnsubscribeNotificationTopic)
}

func NewNotificationTopicHandler(
	logger *logger.Logger,
	notificationTopicSvc service.NotificationTopicService,
) *NotificationTopicHandler {
	return &NotificationTopicHandler{
		logger:               logger,
		notificationTopicSvc: notificationTopicSvc,
	}
}
