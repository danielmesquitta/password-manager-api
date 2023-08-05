package controller

import (
	"net/http"

	"github.com/danielmesquitta/password-manager-api/internal/controller/request"
	"github.com/danielmesquitta/password-manager-api/internal/controller/response"
	"github.com/danielmesquitta/password-manager-api/internal/model"
	"github.com/danielmesquitta/password-manager-api/internal/service"
	"github.com/danielmesquitta/password-manager-api/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type PasswordCardController struct{}

func NewPasswordCardController() PasswordCardController {
	return PasswordCardController{}
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
func (_ PasswordCardController) CreatePasswordCard(c *fiber.Ctx) error {
	dto := request.CreatePasswordCardRequest{}

	if err := c.BodyParser(&dto); err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to parse JSON")
	}

	if errs := validator.Validator.Validate(dto); errs != nil {
		return fiber.NewError(http.StatusBadRequest, validator.Validator.FormatErrs(errs))
	}

	data := model.PasswordCard{
		Name:     dto.Name,
		Url:      dto.Url,
		Username: dto.Username,
		Password: dto.Password,
	}

	if err := service.CreatePasswordCardService(data); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	c.Status(http.StatusCreated)

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
func (_ PasswordCardController) DeletePasswordCard(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := service.DeletePasswordCardService(id); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	c.Status(http.StatusNoContent)

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
func (_ PasswordCardController) ListPasswordCards(c *fiber.Ctx) error {
	search := c.Query("search")

	passwordCards, err := service.ListPasswordCardsService(search)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())

	}

	return c.Status(http.StatusOK).JSON(response.ListPasswordCardsResponse{Data: passwordCards})
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
func (_ PasswordCardController) UpdatePasswordCard(c *fiber.Ctx) error {
	dto := request.UpdatePasswordCardRequest{}

	if err := c.BodyParser(&dto); err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to parse JSON")
	}

	if errs := validator.Validator.Validate(dto); errs != nil {
		return fiber.NewError(http.StatusBadRequest, validator.Validator.FormatErrs(errs))
	}

	id := c.Params("id")

	data := model.PasswordCard{
		Name:     dto.Name,
		Url:      dto.Url,
		Username: dto.Username,
		Password: dto.Password,
	}

	if err := service.UpdatePasswordCardService(id, data); err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to update password card")
	}

	c.Status(http.StatusOK)

	return nil
}
