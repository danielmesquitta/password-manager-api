package controller

import (
	"net/http"

	"github.com/danielmesquitta/password-manager-api/internal/dto"
	"github.com/danielmesquitta/password-manager-api/internal/model"
	"github.com/danielmesquitta/password-manager-api/internal/repository"
	"github.com/danielmesquitta/password-manager-api/internal/service"
	"github.com/danielmesquitta/password-manager-api/pkg/crypt"
	"github.com/danielmesquitta/password-manager-api/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type PasswordCardController struct {
	PasswordCardRepository repository.PasswordCardRepository
	Crypter                crypt.Crypter
}

func NewPasswordCardController(r repository.PasswordCardRepository, c crypt.Crypter) *PasswordCardController {
	return &PasswordCardController{
		PasswordCardRepository: r,
		Crypter:                c,
	}
}

// @BasePath /api/v1
// @Summary Create password card
// @Description Create a new password card
// @Tags PasswordCards
// @Accept json
// @Produce json
// @Param request body request.CreatePasswordCardRequest true "Request body"
// @Success 201
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /password-cards [post]
func (c *PasswordCardController) CreatePasswordCard(ctx *fiber.Ctx) error {
	dto := dto.CreatePasswordCardDTO{}

	if err := ctx.BodyParser(&dto); err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to parse JSON")
	}

	if err := service.CreatePasswordCardService(c.PasswordCardRepository, c.Crypter, dto); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	ctx.Status(http.StatusCreated)

	return nil
}

// @Summary Delete password card
// @Description Delete a password card
// @Tags PasswordCards
// @Accept json
// @Produce json
// @Param id path string true "Password Card identification"
// @Success 204
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /password-cards/{id} [delete]
func (c PasswordCardController) DeletePasswordCard(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := service.DeletePasswordCardService(c.PasswordCardRepository, id); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	ctx.Status(http.StatusNoContent)

	return nil
}

// @BasePath /api/v1
// @Summary List password cards
// @Description List all password cards
// @Tags PasswordCards
// @Accept json
// @Produce json
// @Param search query string false "Search by name"
// @Success 200 {object} response.ListPasswordCardsResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /password-cards [get]
func (c PasswordCardController) ListPasswordCards(ctx *fiber.Ctx) error {
	search := ctx.Query("search")

	passwordCards := []model.PasswordCard{}
	if err := service.ListPasswordCardsService(c.PasswordCardRepository, c.Crypter, search, &passwordCards); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(response.ListResponse[model.PasswordCard]{Data: passwordCards})
}

// @BasePath /api/v1
// @Summary Update password card
// @Description Update a password card
// @Tags PasswordCards
// @Accept json
// @Produce json
// @Param request body request.UpdatePasswordCardRequest true "Request body"
// @Success 201
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /password-cards/{id} [put]
func (c PasswordCardController) UpdatePasswordCard(ctx *fiber.Ctx) error {
	dto := dto.UpdatePasswordCardDTO{}

	if err := ctx.BodyParser(&dto); err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to parse JSON")
	}

	id := ctx.Params("id")

	if err := service.UpdatePasswordCardService(c.PasswordCardRepository, c.Crypter, id, dto); err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to update password card")
	}

	ctx.Status(http.StatusOK)

	return nil
}
