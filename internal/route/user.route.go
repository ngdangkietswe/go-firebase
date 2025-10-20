/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package route

import (
	"go-firebase/internal/controller"

	"github.com/gofiber/fiber/v2"
)

type UserRoute struct {
	userCtrl *controller.UserCtrl
}

func (r *UserRoute) Register(router fiber.Router) {
	userRouter := router.Group("/users")
	userRouter.Post("/", r.userCtrl.CreateUser)
	userRouter.Get("/", r.userCtrl.GetUser)
}

func NewUserRoute(userCtrl *controller.UserCtrl) *UserRoute {
	return &UserRoute{
		userCtrl: userCtrl,
	}
}
