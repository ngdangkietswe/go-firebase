/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package controller

import (
	"errors"
	"go-firebase/internal/handler"
	"go-firebase/internal/request"
	"go-firebase/internal/response"

	"github.com/gofiber/fiber/v2"
)

type AuthCtrl struct {
	authHandler *handler.AuthHandler
}

// Login godoc
// @Summary      User Login
// @Description  Authenticates a user and returns Firebase ID token.
// @Tags         Auth API
// @Accept       json
// @Produce      json
// @Param        login  body      request.LoginRequest  true  "Login Info"
// @Success      200    {object}  response.ApiResponse
// @Failure      400    {object}  response.ApiResponse
// @Failure      401    {object}  response.ApiResponse
// @Router       /auth/login [post]
func (c *AuthCtrl) Login(ctx *fiber.Ctx) error {
	var req request.LoginRequest

	if err := ctx.BodyParser(&req); err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusBadRequest, err)
	}

	res, err := c.authHandler.Login(ctx, &req)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusUnauthorized, err)
	}

	return response.ApiSuccessResponse(ctx, res)
}

// VerifyToken godoc
// @Summary      Verify Firebase ID Token
// @Description  Verifies the provided Firebase ID token.
// @Tags         Auth API
// @Accept       json
// @Produce      json
// @Param        id_token   query     string  true  "Firebase ID Token"
// @Success      200        {object}  response.ApiResponse
// @Failure      400        {object}  response.ApiResponse
// @Failure      401        {object}  response.ApiResponse
// @Router       /auth/verify-token [get]
func (c *AuthCtrl) VerifyToken(ctx *fiber.Ctx) error {
	var req request.VerifyTokenRequest

	idToken := ctx.Query("id_token")
	if idToken == "" {
		return response.ApiErrorResponse(ctx, fiber.StatusBadRequest, errors.New("id_token is required"))
	}

	req.IDToken = idToken

	res, err := c.authHandler.VerifyToken(ctx, &req)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusUnauthorized, err)
	}

	return response.ApiSuccessResponse(ctx, res)
}

func NewAuthCtrl(
	authHandler *handler.AuthHandler,
) *AuthCtrl {
	return &AuthCtrl{
		authHandler: authHandler,
	}
}
