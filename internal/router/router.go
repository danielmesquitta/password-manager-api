package router

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
)

// @title Password Manager API
// @version 1.0
// @description Password Manager API documentation
// @contact.name Daniel Mesquita
// @contact.email danielmesquitta123@gmail.com
// @host localhost:3000
// @BasePath /api/v1
func Init() {
	app := NewApp()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
