/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type loggerMiddleware struct {
}

func (m *loggerMiddleware) Skip(ctx *fiber.Ctx) bool {
	//TODO implement me
	panic("implement me")
}

func (m *loggerMiddleware) AsMiddleware() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "[${time}] ${status} | ${latency} | ${ip} | ${method} ${path} ${queryParams} | ${bytesSent}B ${error}\n",
		TimeFormat: time.RFC3339,
		TimeZone:   "Asia/Ho_Chi_Minh",
		Output:     os.Stdout,
	})
}

func NewLoggerMiddleware() Middleware {
	return &loggerMiddleware{}
}
