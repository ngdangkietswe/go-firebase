/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package server

import (
	"context"
	"go-firebase/internal/middleware"

	"fmt"
	"go-firebase/internal/route"

	_ "go-firebase/docs" // swagger docs

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type HttpServerProps struct {
	fx.In
	Lifecycle   fx.Lifecycle
	Logger      *logger.Logger
	FiberApp    *fiber.App
	Route       *route.AppRoute
	Middlewares []middleware.Middleware `group:"middlewares"`
}

func NewFiberApp() *fiber.App {
	return fiber.New(fiber.Config{
		AppName: "Go Firebase App",
	})
}

func NewHttpServer(props HttpServerProps) {
	props.Lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			port := config.GetInt("APP_PORT", 3000)

			for _, m := range props.Middlewares {
				props.FiberApp.Use(m.AsMiddleware())
			}

			props.FiberApp.Get("/swagger/*", swagger.HandlerDefault)
			props.Route.Register(props.FiberApp)

			go func() {
				if err := props.FiberApp.Listen(fmt.Sprintf(":%d", port)); err != nil {
					panic(err)
				}
			}()

			props.Logger.Info("Server started", zap.Int("port", port))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := props.FiberApp.Shutdown(); err != nil {
				return err
			}
			props.Logger.Info("Server stopped")
			return nil
		},
	})
}

var Module = fx.Module("http-server",
	fx.Provide(
		NewFiberApp,
	),
	fx.Invoke(
		NewHttpServer,
	),
)
