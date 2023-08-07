package app

import (
	docs "github.com/danielmesquitta/password-manager-api/internal/http/docs"
	"github.com/danielmesquitta/password-manager-api/internal/http/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func initializeRoutes(
	app *fiber.App,
	passwordCardHandler *handler.PasswordCardHandler,
) {
	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	apiV1 := app.Group(basePath)
	{
		apiV1.Post(
			"/password-cards",
			passwordCardHandler.CreatePasswordCard,
		)
		apiV1.Get(
			"/password-cards",
			passwordCardHandler.ListPasswordCards,
		)
		apiV1.Delete(
			"/password-cards/:id",
			passwordCardHandler.DeletePasswordCard,
		)
		apiV1.Put(
			"/password-cards/:id",
			passwordCardHandler.UpdatePasswordCard,
		)
	}

	app.Get("/docs/*", swagger.New())
}
