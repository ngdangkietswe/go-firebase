/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package route

import (
	"go-firebase/internal/controller"

	"github.com/gofiber/fiber/v2"
)

type AuthRoute struct {
	authCtrl *controller.AuthCtrl
}

func (r *AuthRoute) Register(router fiber.Router) {
	authRouter := router.Group("/auth")
	authRouter.Post("/login", r.authCtrl.Login)
	authRouter.Get("/verify-token", r.authCtrl.VerifyToken)
}

func NewAuthRoute(
	authCtrl *controller.AuthCtrl,
) *AuthRoute {
	return &AuthRoute{
		authCtrl: authCtrl,
	}
}
