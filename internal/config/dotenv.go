package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file", err)
	}
}
