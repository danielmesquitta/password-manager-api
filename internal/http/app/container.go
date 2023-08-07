package app

import (
	"github.com/danielmesquitta/password-manager-api/internal/config"
	"github.com/danielmesquitta/password-manager-api/internal/http/handler"
	"github.com/danielmesquitta/password-manager-api/internal/pkg/crypt"
	"github.com/danielmesquitta/password-manager-api/internal/pkg/validator"
	"github.com/danielmesquitta/password-manager-api/internal/repository"
	"github.com/gofiber/fiber/v2"

	"go.uber.org/fx"
)

func Init() {
	depsProvider := fx.Provide(
		config.LoadEnv,
		validator.New,
		handler.NewPasswordCardHandler,
		fx.Annotate(
			crypt.New,
			fx.As(new(crypt.Crypter)),
		),
		fx.Annotate(
			repository.NewJsonPasswordCardRepository,
			fx.As(new(repository.PasswordCardRepository)),
		),
		New,
	)

	container := fx.New(
		depsProvider,
		fx.Invoke(func(*fiber.App) {}),
	)

	container.Run()
}
