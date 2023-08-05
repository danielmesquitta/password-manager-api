package service

import (
	"time"

	"github.com/danielmesquitta/password-manager-api/internal/config"
	"github.com/danielmesquitta/password-manager-api/internal/model"
	"github.com/danielmesquitta/password-manager-api/pkg/crypt"
	"github.com/danielmesquitta/password-manager-api/pkg/jsonmanager"
	"github.com/google/uuid"
)

func CreatePasswordCardService(data model.PasswordCard) error {
	encryptedPassword, err := crypt.Encrypt(data.Password)
	if err != nil {
		return err
	}

	passwordCard := model.PasswordCard{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      data.Name,
		Url:       data.Url,
		Username:  data.Username,
		Password:  encryptedPassword,
	}

	file, err := jsonmanager.Open(config.JsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	passwordCards := []model.PasswordCard{}
	if err := jsonmanager.Decode(file, &passwordCards); err != nil {
		return err
	}

	passwordCards = append(passwordCards, passwordCard)

	if err := jsonmanager.Encode(file, passwordCards); err != nil {
		return err
	}

	return nil
}
