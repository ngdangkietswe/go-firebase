/**
 * Author : ngdangkietswe
 * Since  : 10/26/2025
 */

package route

import (
	"go-firebase/internal/controller"

	"github.com/gofiber/fiber/v2"
)

type FileStorageRoute struct {
	fileStorageCtrl *controller.FileStorageCtrl
}

func (r *FileStorageRoute) Register(router fiber.Router) {
	fileStorageRouter := router.Group("/file-storage")
	fileStorageRouter.Post("/upload", r.fileStorageCtrl.UploadFile)
	fileStorageRouter.Post("/download", r.fileStorageCtrl.DownloadFile)
}

func NewFileStorageRoute(
	fileStorageCtrl *controller.FileStorageCtrl,
) *FileStorageRoute {
	return &FileStorageRoute{
		fileStorageCtrl: fileStorageCtrl,
	}
}
