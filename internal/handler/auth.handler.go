/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
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

type AuthHandler struct {
	logger  *logger.Logger
	authSvc service.AuthService
}

func (h *AuthHandler) Login(c *fiber.Ctx, request *request.LoginRequest) (*response.LoginResponse, error) {
	newCtx, cancel := context.WithTimeout(apiutil.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, request, h.authSvc.Login)
}

func (h *AuthHandler) VerifyToken(c *fiber.Ctx, request *request.VerifyTokenRequest) (*response.EmptyResponse, error) {
	newCtx, cancel := context.WithTimeout(apiutil.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, request, h.authSvc.VerifyToken)
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
