/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
)

type corsMiddleware struct {
	logger *logger.Logger
}

func (m *corsMiddleware) Skip(ctx *fiber.Ctx) bool {
	//TODO implement me
	panic("implement me")
}

func (m *corsMiddleware) AsMiddleware() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	})
}

func NewCORSMiddleware(
	logger *logger.Logger,
) Middleware {
	return &corsMiddleware{
		logger: logger,
	}
}
