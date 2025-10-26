/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package middleware

import (
	"go-firebase/internal/firebase"
	"go-firebase/internal/helper"
	"go-firebase/pkg/constant"
	"go-firebase/pkg/response"
	"go-firebase/pkg/util"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	fAuthCli   firebase.FAuthClient
	authHelper helper.AuthHelper
}

var SkipEndpoint = []string{
	"/api/v1/auth/login",
	"/api/v1/auth/verify-token",
	"/api/v1/auth/refresh-token",
	"/swagger",
}

var SkipEndpointAndMethod = map[string]string{
	//fiber.MethodPost: "/api/v1/users",
}

func (m *AuthMiddleware) Skip(ctx *fiber.Ctx) bool {
	for _, endpoint := range SkipEndpoint {
		if strings.Contains(ctx.Path(), endpoint) {
			return true
		}
	}

	//for method, endpoint := range SkipEndpointAndMethod {
	//	if ctx.Method() == method && strings.Contains(ctx.Path(), endpoint) {
	//		return true
	//	}
	//}

	return false
}

func (m *AuthMiddleware) AsMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if m.Skip(ctx) {
			return ctx.Next()
		}

		idToken, err := util.GetIDToken(ctx)
		if err != nil {
			return response.ApiErrorResponse(
				ctx,
				fiber.StatusUnauthorized,
				err,
			)
		}

		claims, err := m.fAuthCli.VerifyIDToken(idToken)
		if err != nil {
			return response.ApiErrorResponse(
				ctx,
				fiber.StatusUnauthorized,
				err,
			)
		}

		if principal, err := m.authHelper.BuildPrincipal(claims); err != nil {
			return response.ApiErrorResponse(
				ctx,
				fiber.StatusUnauthorized,
				err,
			)
		} else {
			ctx.Locals(constant.CtxPrincipalKey, principal)
		}

		return ctx.Next()
	}
}

func NewAuthMiddleware(
	fAuthCli firebase.FAuthClient,
	authHelper helper.AuthHelper,
) Middleware {
	return &AuthMiddleware{
		fAuthCli:   fAuthCli,
		authHelper: authHelper,
	}
}
