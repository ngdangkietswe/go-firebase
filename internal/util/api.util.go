/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package util

import (
	"context"
	"go-firebase/pkg/constant"

	"github.com/gofiber/fiber/v2"
)

var FiberKeys = []string{
	constant.CtxFirebaseUIDKey,
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
