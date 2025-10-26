/**
 * Author : ngdangkietswe
 * Since  : 10/26/2025
 */

package controller

import (
	"go-firebase/internal/handler"
	"go-firebase/pkg/request"
	"go-firebase/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type FileStorageCtrl struct {
	fileStorageHandler *handler.FileStorageHandler
}

// UploadFile godoc
// @Summary      Generate a presigned upload URL
// @Description  Generate a presigned URL for uploading a file to Firebase Storage
// @Tags         File Storage API
// @Accept       json
// @Produce      json
// @Param        bucket_name   query     string  true   "Bucket Name"
// @Param        object_name   query     string  true   "Object Name"
// @Param        content_type  query     string  false  "Content Type"  default(image/png)
// @Param        expire_time   query     int64   false  "Expire Time in seconds"  default(0)
// @Success      200           {object}  response.ApiResponse
// @Failure      400           {object}  response.ApiResponse
// @Failure      500           {object}  response.ApiResponse
// @Router       /file-storage/upload [post]
// @Security     JWT
func (c *FileStorageCtrl) UploadFile(ctx *fiber.Ctx) error {
	var req *request.GetPresignURLRequest

	if err := ctx.BodyParser(&req); err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusBadRequest, err)
	}

	res, err := c.fileStorageHandler.UploadFile(ctx, req)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusInternalServerError, err)
	}

	return response.ApiSuccessResponse(ctx, res)
}

// DownloadFile godoc
// @Summary      Generate a presigned download URL
// @Description  Generate a presigned URL for downloading a file from Firebase Storage
// @Tags         File Storage API
// @Accept       json
// @Produce      json
// @Param        bucket_name   query     string  true   "Bucket Name"
// @Param        object_name   query     string  true   "Object Name"
// @Param        expire_time   query     int64   false  "Expire Time in seconds"  default(0)
// @Success      200           {object}  response.ApiResponse
// @Failure      400           {object}  response.ApiResponse
// @Failure      500           {object}  response.ApiResponse
// @Router       /file-storage/download [post]
// @Security     JWT
func (c *FileStorageCtrl) DownloadFile(ctx *fiber.Ctx) error {
	var req *request.GetPresignURLRequest

	if err := ctx.BodyParser(&req); err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusBadRequest, err)
	}

	res, err := c.fileStorageHandler.DownloadFile(ctx, req)
	if err != nil {
		return response.ApiErrorResponse(ctx, fiber.StatusInternalServerError, err)
	}

	return response.ApiSuccessResponse(ctx, res)
}

func NewFileStorageController(
	fileStorageHandler *handler.FileStorageHandler,
) *FileStorageCtrl {
	return &FileStorageCtrl{
		fileStorageHandler: fileStorageHandler,
	}
}
