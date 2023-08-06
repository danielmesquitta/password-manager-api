package config

import (
	"flag"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slices"
)

var GoEnv string

func loadEnv() {
	listenEnv := flag.String("env", "development", "Environment")
	flag.Parse()

	GoEnv = *listenEnv

	validEnvs := []string{"development", "production", "test"}

	if !slices.Contains(validEnvs, GoEnv) {
		log.Fatalf("Invalid environment, must be one of %v", validEnvs)
	}

	envFile := ".env." + GoEnv

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("error loading .env file", err)
	}
}
