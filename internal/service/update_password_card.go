package service

import (
	"errors"
	"time"

	"github.com/danielmesquitta/password-manager-api/internal/http/dto"
	"github.com/danielmesquitta/password-manager-api/internal/model"
	"github.com/danielmesquitta/password-manager-api/internal/pkg/crypt"
	"github.com/danielmesquitta/password-manager-api/internal/pkg/validator"
	"github.com/danielmesquitta/password-manager-api/internal/repository"
)

func UpdatePasswordCardService(
	r repository.PasswordCardRepository,
	c crypt.Crypter,
	v *validator.Validator,
	id string,
	data dto.UpdatePasswordCardDTO,
) error {
	if errs := v.Validate(data); errs != nil {
		return errors.New(v.FormatErrs(errs))
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
