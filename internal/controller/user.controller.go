/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package controller

import (
	"errors"
	"go-firebase/internal/handler"
	"go-firebase/internal/request"
	"go-firebase/internal/response"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type UserCtrl struct {
	userHandler *handler.UserHandler
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Create a new user with the provided information
// @Tags         User API
// @Accept       json
// @Produce      json
// @Param        user  body      request.CreateUserRequest  true  "User Info"
// @Success      200   {object}  response.ApiResponse
// @Failure      400   {object}  response.ApiResponse
// @Failure      500   {object}  response.ApiResponse
// @Router       /users [post]
func (c *UserCtrl) CreateUser(ctx *fiber.Ctx) error {
	var req *request.CreateUserRequest

	if err := ctx.BodyParser(&req); err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusBadRequest, err)
	}

	res, err := c.userHandler.CreateUser(ctx, req)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusInternalServerError, err)
	}

	return response.ApiSuccessResponse(ctx, res)
}

// GetUser godoc
// @Summary      Get user
// @Description  Retrieve user information
// @Tags         User API
// @Accept       json
// @Produce      json
// @Param        identifier  query     string                     true  "User Email or ID"
// @Param        preload     query     string                     false "Comma-separated list of related entities to preload"
// @Success      200         {object}  response.ApiResponse
// @Failure      400         {object}  response.ApiResponse
// @Failure      500         {object}  response.ApiResponse
// @Router       /users [get]
func (c *UserCtrl) GetUser(ctx *fiber.Ctx) error {
	req := &request.GetUserRequest{}

	identifier := ctx.Query("identifier")
	if identifier == "" {
		return response.ApiErrorResponse(ctx, fiber.StatusBadRequest, errors.New("identifier is required"))
	}

	req.Identifier = identifier

	preload := ctx.Query("preload")
	if preload != "" {
		req.Preload = strings.Split(preload, ",")
	}

	res, err := c.userHandler.GetUser(ctx, req)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusInternalServerError, err)
	}

	return response.ApiSuccessResponse(ctx, res)
}

func NewUserCtrl(
	userHandler *handler.UserHandler,
) *UserCtrl {
	return &UserCtrl{
		userHandler: userHandler,
	}
}
