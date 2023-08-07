package app

import (
	"context"

	"github.com/danielmesquitta/password-manager-api/internal/config"
	"github.com/danielmesquitta/password-manager-api/internal/handler"
	"github.com/danielmesquitta/password-manager-api/internal/middleware"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/fx"
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
			go log.Fatal(app.Listen("0.0.0.0:" + env.Port))

			return nil
		},

		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}
