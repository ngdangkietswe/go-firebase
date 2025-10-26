/**
 * Author : ngdangkietswe
 * Since  : 10/26/2025
 */

package util

import (
	"context"
	"errors"
	"go-firebase/pkg/constant"
	"go-firebase/pkg/model"
	"go-firebase/pkg/response"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func HasPermission(perm *model.Permission) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		principal := ctx.Locals(constant.CtxPrincipalKey).(*model.Principal)
		if principal == nil {
			return response.ApiErrorResponse(
				ctx,
				fiber.StatusForbidden,
				fiber.ErrForbidden,
			)
		}

		if principal.HasPermission(perm) {
			return ctx.Next()
		}

		return response.ApiErrorResponse(
			ctx,
			fiber.StatusForbidden,
			fiber.ErrForbidden,
		)
	}
}

func GetIDToken(ctx *fiber.Ctx) (string, error) {
	authHeader := ctx.Get(fiber.HeaderAuthorization)
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}

	if !strings.HasPrefix(authHeader, constant.AuthHeaderPrefixBearer) {
		return "", errors.New("invalid authorization header format")
	}

	idToken := strings.TrimPrefix(authHeader, constant.AuthHeaderPrefixBearer)
	return strings.TrimSpace(idToken), nil
}

func GetPrincipal(ctx context.Context) *model.Principal {
	if principal, ok := ctx.Value(constant.CtxPrincipalKey).(*model.Principal); ok {
		return principal
	}
	return nil
}
