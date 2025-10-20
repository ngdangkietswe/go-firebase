/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package response

import (
	"github.com/gofiber/fiber/v2"
)

type EmptyResponse struct{}

type ApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ApiSuccessResponse(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(
		&ApiResponse{
			Status:  200,
			Message: "Success",
			Data:    data,
		},
	)
}

func ApiErrorResponse(ctx *fiber.Ctx, status int, err error) error {
	return ctx.Status(status).JSON(
		&ApiResponse{
			Status:  status,
			Message: err.Error(),
		},
	)
}
