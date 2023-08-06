package router

import (
	"github.com/danielmesquitta/password-manager-api/internal/controller"
	docs "github.com/danielmesquitta/password-manager-api/internal/docs"
	"github.com/danielmesquitta/password-manager-api/internal/repository"
	"github.com/danielmesquitta/password-manager-api/pkg/crypt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func initializeRoutes(app *fiber.App) {
	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	jsonPasswordCardRepository := repository.NewJsonPasswordCardRepository()
	crypt := crypt.New()

	passwordCardController := controller.NewPasswordCardController(jsonPasswordCardRepository, crypt)

	v1 := app.Group(basePath)
	{
		v1.Post("/password-cards", passwordCardController.CreatePasswordCard)
		v1.Get("/password-cards", passwordCardController.ListPasswordCards)
		v1.Delete("/password-cards/:id", passwordCardController.DeletePasswordCard)
		v1.Put("/password-cards/:id", passwordCardController.UpdatePasswordCard)
	}

	app.Get("/docs/*", swagger.New())
}
