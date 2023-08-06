package repository

import (
	"errors"
	"regexp"

	"github.com/danielmesquitta/password-manager-api/internal/model"
)

type InMemoryPasswordCardRepository struct {
	PasswordCards []model.PasswordCard
}

func NewInMemoryPasswordCardRepository() *InMemoryPasswordCardRepository {
	return &InMemoryPasswordCardRepository{
		PasswordCards: []model.PasswordCard{},
	}
}

func (r *InMemoryPasswordCardRepository) CreatePasswordCard(data model.PasswordCard) error {
	r.PasswordCards = append(r.PasswordCards, data)

	return nil
}

func (r *InMemoryPasswordCardRepository) UpdatePasswordCard(id string, data model.PasswordCard) error {
	index := -1
	for i, passwordCard := range r.PasswordCards {
		if passwordCard.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("Password card not found")
	}

	r.PasswordCards[index] = data

	return nil
}

func (r *InMemoryPasswordCardRepository) DeletePasswordCard(id string) error {
	index := -1
	for i, item := range r.PasswordCards {
		if item.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("not found")
	}

	r.PasswordCards = append(r.PasswordCards[:index], r.PasswordCards[index+1:]...)

	return nil
}

func (r *InMemoryPasswordCardRepository) GetPasswordCard(id string, pcPtr *model.PasswordCard) error {
	for _, passwordCard := range r.PasswordCards {
		if passwordCard.ID == id {
			*pcPtr = passwordCard

			return nil
		}
	}

	return nil
}

func (r *InMemoryPasswordCardRepository) ListPasswordCards(search string, pcsPtr *[]model.PasswordCard) error {
	if search == "" {
		*pcsPtr = r.PasswordCards

		return nil
	}

	searchRegex, err := regexp.Compile("(?i)" + search)
	if err != nil {
		return err
	}

	filteredPasswordCards := make([]model.PasswordCard, 0, len(r.PasswordCards))

	for _, passwordCard := range r.PasswordCards {
		if searchRegex.MatchString(passwordCard.Name) {
			filteredPasswordCards = append(filteredPasswordCards, passwordCard)
		}
	}

	*pcsPtr = filteredPasswordCards

	return nil
}
