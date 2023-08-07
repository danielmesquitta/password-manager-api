package app

import (
	"context"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/fx"

	"github.com/danielmesquitta/password-manager-api/internal/config"
	"github.com/danielmesquitta/password-manager-api/internal/http/handler"
	"github.com/danielmesquitta/password-manager-api/internal/http/middleware"
)

func New(
	lc fx.Lifecycle,
	env *config.EnvVars,
	passwordCardHandler *handler.PasswordCardHandler,
) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: middleware.ErrorHandlerMiddleware,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: env.CorsOrigin,
	}))

	app.Use(recover.New())

	initializeRoutes(app, passwordCardHandler)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				err := app.Listen("0.0.0.0:" + env.Port)

				if err != nil {
					panic(err)
				}
			}()

			return nil
		},

		OnStop: func(context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}
