/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package controller

import (
	"go-firebase/internal/handler"
	"go-firebase/internal/request"
	"go-firebase/internal/response"

	"github.com/gofiber/fiber/v2"
)

type DeviceTokenCtrl struct {
	deviceTokenHandler *handler.DeviceTokenHandler
}

// RegisterDeviceToken godoc
// @Summary      Register a device token
// @Description  Register a device token for push notifications
// @Tags         Device Token API
// @Accept       json
// @Produce      json
// @Param        device  body      request.RegisterDeviceRequest  true  "Device Info"
// @Success      200     {object}  response.ApiResponse
// @Failure      400     {object}  response.ApiResponse
// @Failure      500     {object}  response.ApiResponse
// @Router       /device-tokens [post]
func (c *DeviceTokenCtrl) RegisterDeviceToken(ctx *fiber.Ctx) error {
	var req *request.RegisterDeviceRequest

	if err := ctx.BodyParser(&req); err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusBadRequest, err)
	}

	res, err := c.deviceTokenHandler.RegisterDeviceToken(ctx, req)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusInternalServerError, err)
	}

	return response.ApiSuccessResponse(ctx, res)
}

func NewDeviceTokenCtrl(
	deviceTokenHandler *handler.DeviceTokenHandler,
) *DeviceTokenCtrl {
	return &DeviceTokenCtrl{
		deviceTokenHandler: deviceTokenHandler,
	}
}
