package service

import (
	"github.com/danielmesquitta/password-manager-api/internal/repository"
)

func DeletePasswordCardService(
	r repository.PasswordCardRepository,
	id string,
) error {
	err := r.DeletePasswordCard(id)
	if err != nil {
		return err
	}

	return nil
}
