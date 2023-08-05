package service

import (
	"time"

	"github.com/danielmesquitta/password-manager-api/internal/config"
	"github.com/danielmesquitta/password-manager-api/internal/model"
	"github.com/danielmesquitta/password-manager-api/pkg/crypt"
	"github.com/danielmesquitta/password-manager-api/pkg/jsonmanager"
)

func UpdatePasswordCardService(id string, data model.PasswordCard) error {
	file, err := jsonmanager.Open(config.JsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	passwordCards := []model.PasswordCard{}
	if err := jsonmanager.Decode(file, &passwordCards); err != nil {
		return err
	}

	// Find the password card with the given id and update it
	for i, passwordCard := range passwordCards {
		if passwordCard.ID == id {
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
				encryptedPassword, err := crypt.Encrypt(data.Password)
				if err != nil {
					return err
				}

				passwordCard.Password = encryptedPassword
			}

			passwordCards[i] = passwordCard
		}
	}

	if err := jsonmanager.Encode(file, passwordCards); err != nil {
		return err
	}

	return nil
}
