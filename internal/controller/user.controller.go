/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package controller

import (
	"go-firebase/internal/handler"
	"go-firebase/pkg/request"
	"go-firebase/pkg/response"
	"go-firebase/pkg/util"
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
// @Security     JWT
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

// GetUsers godoc
// @Summary      Get users list
// @Description  Retrieve a list of users with pagination and optional search
// @Tags         User API
// @Accept       json
// @Produce      json
// @Param        page       query     int                        false "Page number"
// @Param        page_size  query     int                        false "Number of items per page"
// @Param        sort       query     string                     false "Sort by field"
// @Param        order      query     string                     false "Order (asc or desc)"
// @Param        search     query     string                     false "Search term"
// @Param        preload    query     string                     false "Comma-separated list of related entities to preload"
// @Success      200        {object}  response.ApiResponse
// @Failure      400        {object}  response.ApiResponse
// @Failure      500        {object}  response.ApiResponse
// @Router       /users [get]
// @Security     JWT
func (c *UserCtrl) GetUsers(ctx *fiber.Ctx) error {
	req := &request.ListUserRequest{
		Paginate: util.AsPaginateRequest(ctx),
		Search:   ctx.Query("search"),
		Preload:  strings.Split(ctx.Query("preload"), ","),
	}

	res, err := c.userHandler.GetUsers(ctx, req)
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
// @Param        identifier  path      string                     true  "User Identifier (ID or Email)"
// @Param        preload     query     string                     false "Comma-separated list of related entities to preload"
// @Success      200         {object}  response.ApiResponse
// @Failure      400         {object}  response.ApiResponse
// @Failure      500         {object}  response.ApiResponse
// @Router       /users/{identifier} [get]
// @Security     JWT
func (c *UserCtrl) GetUser(ctx *fiber.Ctx) error {
	req := &request.GetUserRequest{
		Identifier: ctx.Params("identifier"),
		Preload:    strings.Split(ctx.Query("preload"), ","),
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
