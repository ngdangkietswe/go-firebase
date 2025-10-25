/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package util

import (
	"context"
	"go-firebase/pkg/constant"
	"go-firebase/pkg/request"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var FiberKeys = []string{
	constant.CtxFirebaseUIDKey,
	constant.CtxSysUIDKey,
	constant.CtxPrincipalKey,
}

func FiberCtxToContext(ctx *fiber.Ctx) context.Context {
	newCtx := context.Background()
	for _, key := range FiberKeys {
		if val := ctx.Locals(key); val != nil {
			newCtx = context.WithValue(newCtx, key, val)
		}
	}
	return newCtx
}

func AsPaginateRequest(ctx *fiber.Ctx) *request.PaginateRequest {
	return &request.PaginateRequest{
		Page:     ctx.QueryInt("page", 0),
		PageSize: ctx.QueryInt("page_size", 10),
		Sort:     ctx.Query("sort", ""),
		Order:    ctx.Query("order", ""),
	}
}

func ToBool(val string) bool {
	if strings.TrimSpace(val) == "true" || strings.TrimSpace(val) == "1" {
		return true
	}
	return false
}
