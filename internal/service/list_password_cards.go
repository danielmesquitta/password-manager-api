package service

import (
	"regexp"

	"github.com/danielmesquitta/password-manager-api/internal/config"
	"github.com/danielmesquitta/password-manager-api/internal/model"
	"github.com/danielmesquitta/password-manager-api/pkg/crypt"
	"github.com/danielmesquitta/password-manager-api/pkg/jsonmanager"
)

func ListPasswordCardsService(search string) ([]model.PasswordCard, error) {

	file, err := jsonmanager.Open(config.JsonFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	passwordCards := []model.PasswordCard{}
	if err := jsonmanager.Decode(file, &passwordCards); err != nil {
		return nil, err
	}

	for i, passwordCard := range passwordCards {
		decryptedPassword, err := crypt.Decrypt(passwordCard.Password)
		if err != nil {
			return nil, err
		}

		passwordCards[i].Password = decryptedPassword
	}

	if search != "" {
		searchRegex, err := regexp.Compile("(?i)" + search)
		if err != nil {
			return nil, err
		}

		filteredPasswordCards := make([]model.PasswordCard, 0, len(passwordCards))

		for _, passwordCard := range passwordCards {
			if searchRegex.MatchString(passwordCard.Name) {
				filteredPasswordCards = append(filteredPasswordCards, passwordCard)
			}
		}

		passwordCards = filteredPasswordCards
	}

	return passwordCards, nil
}
