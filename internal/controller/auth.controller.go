/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package controller

import (
	"errors"
	"go-firebase/internal/handler"
	"go-firebase/pkg/request"
	"go-firebase/pkg/response"
	"go-firebase/pkg/util"

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

// RefreshToken godoc
// @Summary      Refresh Firebase ID Token
// @Description  Refreshes the Firebase ID token using the provided refresh token.
// @Tags         Auth API
// @Accept       json
// @Produce      json
// @Param        refresh  body      request.RefreshTokenRequest  true  "Refresh Token Info"
// @Success      200      {object}  response.ApiResponse
// @Failure      400      {object}  response.ApiResponse
// @Failure      401      {object}  response.ApiResponse
// @Router       /auth/refresh-token [post]
func (c *AuthCtrl) RefreshToken(ctx *fiber.Ctx) error {
	var req request.RefreshTokenRequest

	if err := ctx.BodyParser(&req); err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusBadRequest, err)
	}

	res, err := c.authHandler.RefreshToken(ctx, &req)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusUnauthorized, err)
	}

	return response.ApiSuccessResponse(ctx, res)
}

// RevokeToken godoc
// @Summary		Revoke Firebase ID Token
// @Description	Revokes the Firebase ID token for the specified user.
// @Tags		Auth API
// @Accept		json
// @Produce		json
// @Success		200           {object}  response.ApiResponse
// @Failure		400           {object}  response.ApiResponse
// @Failure		500           {object}  response.ApiResponse
// @Router		/auth/revoke-token [post]
func (c *AuthCtrl) RevokeToken(ctx *fiber.Ctx) error {
	req := &request.RevokeTokenRequest{
		FirebaseUID: util.GetPrincipalByFiberCtx(ctx).FirebaseUID,
	}

	res, err := c.authHandler.RevokeToken(ctx, req)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusInternalServerError, err)
	}

	return response.ApiSuccessResponse(ctx, res)
}

// CurrentUser godoc
// @Summary      Get Current User
// @Description  Retrieves information about the currently authenticated user.
// @Tags         Auth API
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.ApiResponse
// @Failure      500  {object}  response.ApiResponse
// @Router       /auth/me [get]
func (c *AuthCtrl) CurrentUser(ctx *fiber.Ctx) error {
	res, err := c.authHandler.CurrentUser(ctx)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusInternalServerError, err)
	}

	return response.ApiSuccessResponse(ctx, res)
}

// ForgotPassword godoc
// @Summary      	Forgot Password
// @Description  	Sends a password reset email to the user.
// @Tags         	Auth API
// @Accept       	json
// @Produce      	json
// @Param        	forgot_password  body      request.SendPasswordResetMailRequest  true  "Forgot Password Info"
// @Success      	200              {object}  response.ApiResponse
// @Failure     	400              {object}  response.ApiResponse
// @Failure      	500              {object}  response.ApiResponse
// @Router       	/auth/forgot-password [post]
func (c *AuthCtrl) ForgotPassword(ctx *fiber.Ctx) error {
	var req request.SendPasswordResetMailRequest

	if err := ctx.BodyParser(&req); err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusBadRequest, err)
	}

	_, err := c.authHandler.ForgotPassword(ctx, &req)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusInternalServerError, err)
	}

	return response.ApiSuccessEmptyResponse(ctx)
}

// ResetPassword godoc
// @Summary 		Reset Password
// @Description 	Resets the user's password using the provided oobCode and new password.
// @Tags        	Auth API
// @Accept      	json
// @Produce     	json
// @Param       	reset_password  body      request.ResetPasswordRequest  true  "Reset Password Info"
// @Success     	200             {object}  response.ApiResponse
// @Failure     	400             {object}  response.ApiResponse
// @Failure     	500             {object}  response.ApiResponse
// @Router      	/auth/reset-password [post]
func (c *AuthCtrl) ResetPassword(ctx *fiber.Ctx) error {
	var req request.ResetPasswordRequest

	if err := ctx.BodyParser(&req); err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusBadRequest, err)
	}

	_, err := c.authHandler.ResetPassword(ctx, &req)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusInternalServerError, err)
	}

	return response.ApiSuccessEmptyResponse(ctx)
}

// AdminChangePassword godoc
// @Summary				Admin Change Password
// @Description			Allows an admin to change a user's password using the user's Firebase UID and the new password.
// @Tags				Auth API
// @Accept				json
// @Produce				json
// @Param				change_password  body      request.ChangePasswordRequest  true  "Change Password Info"
// @Success				200              {object}  response.ApiResponse
// @Failure				400              {object}  response.ApiResponse
// @Failure				500              {object}  response.ApiResponse
// @Router				/auth/admin/change-password [post]
func (c *AuthCtrl) AdminChangePassword(ctx *fiber.Ctx) error {
	var req request.ChangePasswordRequest

	if err := ctx.BodyParser(&req); err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusBadRequest, err)
	}

	_, err := c.authHandler.AdminChangePassword(ctx, &req)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusInternalServerError, err)
	}

	return response.ApiSuccessEmptyResponse(ctx)
}

func NewAuthController(
	authHandler *handler.AuthHandler,
) *AuthCtrl {
	return &AuthCtrl{
		authHandler: authHandler,
	}
}
