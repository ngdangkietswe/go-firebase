/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package middleware

import (
	"github.com/gofiber/fiber/v2"
)

type Middleware interface {
	Skip(ctx *fiber.Ctx) bool
	AsMiddleware() fiber.Handler
}
