/**
 * Author : ngdangkietswe
 * Since  : 10/26/2025
 */

package handler

import (
	"context"
	"go-firebase/internal/service"
	"go-firebase/pkg/constant"
	"go-firebase/pkg/request"
	"go-firebase/pkg/response"
	"go-firebase/pkg/util"

	"github.com/gofiber/fiber/v2"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
)

type FileStorageHandler struct {
	logger         *logger.Logger
	fileStorageSvc service.FileStorageService
}

func (h *FileStorageHandler) UploadFile(c *fiber.Ctx, request *request.GetPresignURLRequest) (*response.PresignURLResponse, error) {
	newCtx, cancel := context.WithTimeout(util.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, request, h.fileStorageSvc.GenerateUploadURL)
}

func (h *FileStorageHandler) DownloadFile(c *fiber.Ctx, request *request.GetPresignURLRequest) (*response.PresignURLResponse, error) {
	newCtx, cancel := context.WithTimeout(util.FiberCtxToContext(c), constant.CtxTimeOut)
	defer cancel()
	return util.SafeFunc(newCtx, request, h.fileStorageSvc.GenerateDownloadURL)
}

func NewFileStorageHandler(
	logger *logger.Logger,
	fileStorageSvc service.FileStorageService,
) *FileStorageHandler {
	return &FileStorageHandler{
		logger:         logger,
		fileStorageSvc: fileStorageSvc,
	}
}
