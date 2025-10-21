/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
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

type UserHandler struct {
	logger  *logger.Logger
	userSvc service.UserService
}

func (h *UserHandler) CreateUser(c *fiber.Ctx, request *request.CreateUserRequest) (*response.CreateUserResponse, error) {
	newCtx, cancel := context.WithTimeout(util.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, request, h.userSvc.CreateUser)
}

func (h *UserHandler) GetUser(c *fiber.Ctx, request *request.GetUserRequest) (*model.User, error) {
	newCtx, cancel := context.WithTimeout(util.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, request, h.userSvc.GetUser)
}

func NewUserHandler(
	logger *logger.Logger,
	userSvc service.UserService,
) *UserHandler {
	return &UserHandler{
		logger:  logger,
		userSvc: userSvc,
	}
}
