/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package response

import (
	"github.com/gofiber/fiber/v2"
)

type EmptyResponse struct{}

type IdResponse struct {
	ID string `json:"id"`
}

type ApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ListResponse struct {
	Items interface{} `json:"items"`
	Meta  *PageMeta   `json:"meta"`
}

type PageMeta struct {
	TotalItems  int  `json:"total_items"`
	TotalPages  int  `json:"total_pages"`
	CurrentPage int  `json:"current_page"`
	PageSize    int  `json:"page_size"`
	HasPrevious bool `json:"has_previous"`
	HasNext     bool `json:"has_next"`
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
