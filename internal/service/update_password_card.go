package service

import (
	"errors"
	"time"

	"github.com/danielmesquitta/password-manager-api/internal/dto"
	"github.com/danielmesquitta/password-manager-api/internal/model"
	"github.com/danielmesquitta/password-manager-api/internal/repository"
	"github.com/danielmesquitta/password-manager-api/pkg/crypt"
	"github.com/danielmesquitta/password-manager-api/pkg/validator"
)

func UpdatePasswordCardService(r repository.PasswordCardRepository, c crypt.Crypter, id string, data dto.UpdatePasswordCardDTO) error {
	validator := validator.NewValidator()
	if errs := validator.Validate(data); errs != nil {
		return errors.New(validator.FormatErrs(errs))
	}

	passwordCard := model.PasswordCard{}
	if err := r.GetPasswordCard(id, &passwordCard); err != nil {
		return err
	}

	passwordCard.UpdatedAt = time.Now()

	if data.Name != "" {
		passwordCard.Name = data.Name
	}

	if data.Url != "" {
		passwordCard.Url = data.Url
	}

	if data.Username != "" {
		passwordCard.Username = data.Username
	}

	if data.Password != "" {
		encryptedPassword, err := c.Encrypt(data.Password)
		if err != nil {
			return err
		}

		passwordCard.Password = encryptedPassword
	}

	if err := r.UpdatePasswordCard(id, passwordCard); err != nil {
		return err
	}

	return nil
}
