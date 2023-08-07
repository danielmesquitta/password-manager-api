package service

import (
	"errors"
	"time"

	"github.com/danielmesquitta/password-manager-api/internal/dto"
	"github.com/danielmesquitta/password-manager-api/internal/model"
	"github.com/danielmesquitta/password-manager-api/internal/pkg/crypt"
	"github.com/danielmesquitta/password-manager-api/internal/pkg/validator"
	"github.com/danielmesquitta/password-manager-api/internal/repository"
	"github.com/google/uuid"
)

func CreatePasswordCardService(
	r repository.PasswordCardRepository,
	c crypt.Crypter,
	v *validator.Validator,
	data dto.CreatePasswordCardDTO,
) error {
	if errs := v.Validate(data); errs != nil {
		return errors.New(v.FormatErrs(errs))
	}

	encryptedPassword, err := c.Encrypt(data.Password)
	if err != nil {
		return err
	}

	passwordCard := model.PasswordCard{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      data.Name,
		Url:       data.Url,
		Username:  data.Username,
		Password:  encryptedPassword,
	}

	err = r.CreatePasswordCard(passwordCard)
	if err != nil {
		return err
	}

	return nil
}
