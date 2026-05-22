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
	authRouter.Post("refresh-token", r.authCtrl.RefreshToken)
	authRouter.Post("/revoke-token", r.authCtrl.RevokeToken)
	authRouter.Get("/verify-token", r.authCtrl.VerifyToken)
	authRouter.Get("/me", r.authCtrl.CurrentUser)
	authRouter.Post("/forgot-password", r.authCtrl.ForgotPassword)
	authRouter.Post("/reset-password", r.authCtrl.ResetPassword)
	authRouter.Post("/admin/change-password", r.authCtrl.AdminChangePassword)
}

func NewAuthRoute(
	authCtrl *controller.AuthCtrl,
) *AuthRoute {
	return &AuthRoute{
		authCtrl: authCtrl,
	}
}
