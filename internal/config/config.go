package config

import (
	"github.com/danielmesquitta/password-manager-api/pkg/crypt"
	"github.com/danielmesquitta/password-manager-api/pkg/validator"
)

func Init() {
	loadEnv()
	crypt.Init()
	validator.Init()
}
