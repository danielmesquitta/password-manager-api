package service

import (
	"github.com/danielmesquitta/password-manager-api/internal/model"
	"github.com/danielmesquitta/password-manager-api/internal/pkg/crypt"
	"github.com/danielmesquitta/password-manager-api/internal/repository"
)

func ListPasswordCardsService(
	r repository.PasswordCardRepository,
	c crypt.Crypter,
	search string,
	pcsPtr *[]model.PasswordCard,
) error {
	err := r.ListPasswordCards(pcsPtr, repository.ListPasswordCardsFilters{
		Search: search,
	})

	if err != nil {
		return err
	}

	for i, passwordCard := range *pcsPtr {
		decryptedPassword, err := c.Decrypt(passwordCard.Password)
		if err != nil {
			return err
		}

		(*pcsPtr)[i].Password = decryptedPassword
	}

	return nil
}
