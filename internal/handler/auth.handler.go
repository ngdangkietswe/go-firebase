/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package handler

import (
	"context"
	"go-firebase/internal/service"
	"go-firebase/pkg/constant"
	"go-firebase/pkg/model"
	"go-firebase/pkg/request"
	"go-firebase/pkg/response"
	"go-firebase/pkg/util"

	"github.com/gofiber/fiber/v2"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
)

type AuthHandler struct {
	logger  *logger.Logger
	authSvc service.AuthService
}

func (h *AuthHandler) Login(c *fiber.Ctx, request *request.LoginRequest) (*response.LoginResponse, error) {
	newCtx, cancel := context.WithTimeout(util.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, request, h.authSvc.Login)
}

func (h *AuthHandler) VerifyToken(c *fiber.Ctx, request *request.VerifyTokenRequest) (*response.EmptyResponse, error) {
	newCtx, cancel := context.WithTimeout(util.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, request, h.authSvc.VerifyToken)
}

func (h *AuthHandler) RefreshToken(c *fiber.Ctx, request *request.RefreshTokenRequest) (*response.RefreshTokenResponse, error) {
	newCtx, cancel := context.WithTimeout(util.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, request, h.authSvc.RefreshToken)
}

func (h *AuthHandler) CurrentUser(c *fiber.Ctx) (*model.User, error) {
	newCtx, cancel := context.WithTimeout(util.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFuncNoReq(newCtx, h.authSvc.CurrentUser)
}

func NewAuthHandler(
	logger *logger.Logger,
	authSvc service.AuthService,
) *AuthHandler {
	return &AuthHandler{
		logger:  logger,
		authSvc: authSvc,
	}
}
