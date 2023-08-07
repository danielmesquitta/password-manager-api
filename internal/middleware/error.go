package middleware

import (
	"errors"

	"github.com/danielmesquitta/password-manager-api/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func ErrorHandlerMiddleware(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	log.Error(err)

	return ctx.Status(code).JSON(response.ErrorResponse{
		Message:    err.Error(),
		StatusCode: code,
	})
}
