package main

import (
	"github.com/danielmesquitta/password-manager-api/internal/config"
	"github.com/danielmesquitta/password-manager-api/internal/router"
	"github.com/danielmesquitta/password-manager-api/pkg/crypt"
	"github.com/danielmesquitta/password-manager-api/pkg/validator"
)

func main() {
	config.Init()
	crypt.Init()
	validator.Init()
	router.Init()
}
