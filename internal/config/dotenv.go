package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	Port       string
	CorsOrigin string
	HashSecret string
}

func LoadEnv() *EnvVars {
	err := godotenv.Load()
	if err != nil {
		panic("cannot load .env file")
	}

	corsOrigin := os.Getenv("CORS_ORIGIN")
	if corsOrigin == "" {
		corsOrigin = "*"
	}

	hashSecret := os.Getenv("HASH_SECRET")
	switch length := len(hashSecret); {
	case length >= 32:
		hashSecret = hashSecret[:32]
	case length >= 24:
		hashSecret = hashSecret[:24]
	case length >= 16:
		hashSecret = hashSecret[:16]
	default:
		panic("HASH_SECRET must be at least 16 characters long")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &EnvVars{
		Port:       port,
		CorsOrigin: corsOrigin,
		HashSecret: hashSecret,
	}
}
