package main

import (
	"github.com/danielmesquitta/password-manager-api/internal/app"
	"github.com/danielmesquitta/password-manager-api/internal/config"
	"github.com/danielmesquitta/password-manager-api/internal/handler"
	"github.com/danielmesquitta/password-manager-api/internal/pkg/crypt"
	"github.com/danielmesquitta/password-manager-api/internal/pkg/validator"
	"github.com/danielmesquitta/password-manager-api/internal/repository"

	"go.uber.org/fx"
)

func main() {
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
	)

	depsInjector := fx.New(
		depsProvider,
		fx.Invoke(app.New),
	)

	depsInjector.Run()
}
