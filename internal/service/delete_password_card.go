package service

import (
	"errors"

	"github.com/danielmesquitta/password-manager-api/internal/config"
	"github.com/danielmesquitta/password-manager-api/internal/model"
	"github.com/danielmesquitta/password-manager-api/pkg/jsonmanager"
)

func DeletePasswordCardService(id string) error {
	file, err := jsonmanager.Open(config.JsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	passwordCards := []model.PasswordCard{}
	if err := jsonmanager.Decode(file, &passwordCards); err != nil {
		return err
	}

	// Find the password card with the given id index
	index := -1
	for i, item := range passwordCards {
		if item.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("not found")
	}

	// Remove the password card from the slice
	passwordCards = append(passwordCards[:index], passwordCards[index+1:]...)

	if err := jsonmanager.Encode(file, passwordCards); err != nil {
		return err
	}

	return nil
}
