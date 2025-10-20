/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package handler

import (
	"context"
	"go-firebase/internal/request"
	"go-firebase/internal/response"
	"go-firebase/internal/service"
	apiutil "go-firebase/internal/util"
	"go-firebase/pkg/constant"
	"go-firebase/pkg/util"

	"github.com/gofiber/fiber/v2"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
)

type DeviceTokenHandler struct {
	logger         *logger.Logger
	deviceTokenSvc service.DeviceTokenService
}

func (h *DeviceTokenHandler) RegisterDeviceToken(c *fiber.Ctx, request *request.RegisterDeviceRequest) (*response.RegisterDeviceResponse, error) {
	newCtx, cancel := context.WithTimeout(apiutil.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, request, h.deviceTokenSvc.RegisterDeviceToken)
}

func NewDeviceTokenHandler(
	logger *logger.Logger,
	deviceTokenSvc service.DeviceTokenService,
) *DeviceTokenHandler {
	return &DeviceTokenHandler{
		logger:         logger,
		deviceTokenSvc: deviceTokenSvc,
	}
}
