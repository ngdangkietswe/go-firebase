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
	fiber.MethodPost: "/api/v1/users",
}

func (m *AuthMiddleware) Skip(ctx *fiber.Ctx) bool {
	for _, endpoint := range SkipEndpoint {
		if strings.Contains(ctx.Path(), endpoint) {
			return true
		}
	}

	for method, endpoint := range SkipEndpointAndMethod {
		if ctx.Method() == method && strings.Contains(ctx.Path(), endpoint) {
			return true
		}
	}

	return false
}

func (m *AuthMiddleware) AsMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if m.Skip(ctx) {
			return ctx.Next()
		}

		authHeader := ctx.Get(fiber.HeaderAuthorization)
		if authHeader == "" {
			return response.ApiErrorResponse(
				ctx,
				fiber.StatusUnauthorized,
				fiber.ErrUnauthorized,
			)
		}

		if !strings.HasPrefix(authHeader, constant.AuthHeaderPrefixBearer) {
			return response.ApiErrorResponse(
				ctx,
				fiber.StatusUnauthorized,
				fiber.ErrUnauthorized,
			)
		}

		idToken := strings.TrimPrefix(authHeader, constant.AuthHeaderPrefixBearer)
		claims, err := m.fAuthCli.VerifyIDToken(idToken)
		if err != nil {
			return response.ApiErrorResponse(
				ctx,
				fiber.StatusUnauthorized,
				err,
			)
		}

		ctx.Locals(constant.CtxFirebaseUIDKey, claims["firebase_uid"])
		ctx.Locals(constant.CtxSysUIDKey, claims["system_uid"])

		principal, err := m.authHelper.BuildPrincipal(claims)
		if err != nil {
			return response.ApiErrorResponse(
				ctx,
				fiber.StatusUnauthorized,
				err,
			)
		}

		ctx.Locals(constant.CtxPrincipalKey, principal)

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
