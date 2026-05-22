/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package route

import (
	"go-firebase/internal/controller"
	"go-firebase/internal/permission"
	"go-firebase/pkg/util"

	"github.com/gofiber/fiber/v2"
)

type UserRoute struct {
	userCtrl *controller.UserCtrl
}

func (r *UserRoute) Register(router fiber.Router) {
	userRouter := router.Group("/users")
	userRouter.Post("/", util.HasPermission(permission.CreateUserPerm()), r.userCtrl.CreateUser)
	userRouter.Get("/:identifier", util.HasPermission(permission.ReadUserPerm()), r.userCtrl.GetUser)
	userRouter.Get("/", util.HasPermission(permission.ReadUserPerm()), r.userCtrl.GetUsers)
	userRouter.Put("/:id/en-disable", util.HasPermission(permission.UpdateUserPerm()), r.userCtrl.EnDisableUser)
	userRouter.Delete("/:id", util.HasPermission(permission.DeleteUserPerm()), r.userCtrl.DeleteUser)
}

func NewUserRoute(userCtrl *controller.UserCtrl) *UserRoute {
	return &UserRoute{
		userCtrl: userCtrl,
	}
}
