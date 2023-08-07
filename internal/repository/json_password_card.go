package repository

import (
	"errors"
	"regexp"
	"time"

	"github.com/danielmesquitta/password-manager-api/internal/config"
	"github.com/danielmesquitta/password-manager-api/internal/model"
	"github.com/danielmesquitta/password-manager-api/internal/pkg/jsonmanager"
)

type JsonPasswordCardRepository struct{}

func NewJsonPasswordCardRepository() *JsonPasswordCardRepository {
	return &JsonPasswordCardRepository{}
}

func (r JsonPasswordCardRepository) CreatePasswordCard(
	data model.PasswordCard,
) error {
	file, err := jsonmanager.Open(config.JsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	passwordCards := []model.PasswordCard{}
	if err := jsonmanager.Decode(file, &passwordCards); err != nil {
		return err
	}

	passwordCards = append(passwordCards, data)

	if err := jsonmanager.Encode(file, passwordCards); err != nil {
		return err
	}

	return nil
}

func (r JsonPasswordCardRepository) UpdatePasswordCard(
	id string,
	data model.PasswordCard,
) error {
	file, err := jsonmanager.Open(config.JsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	passwordCards := []model.PasswordCard{}
	if err := jsonmanager.Decode(file, &passwordCards); err != nil {
		return err
	}

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
				passwordCard.Password = data.Password
			}

			passwordCards[i] = passwordCard

			break
		}
	}

	if err := jsonmanager.Encode(file, passwordCards); err != nil {
		return err
	}

	return nil
}

func (r JsonPasswordCardRepository) DeletePasswordCard(id string) error {
	file, err := jsonmanager.Open(config.JsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	passwordCards := []model.PasswordCard{}
	if err := jsonmanager.Decode(file, &passwordCards); err != nil {
		return err
	}

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

	passwordCards = append(passwordCards[:index], passwordCards[index+1:]...)

	if err := jsonmanager.Encode(file, passwordCards); err != nil {
		return err
	}

	return nil
}

func (r JsonPasswordCardRepository) GetPasswordCard(
	id string,
	p *model.PasswordCard,
) error {
	file, err := jsonmanager.Open(config.JsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	passwordCards := []model.PasswordCard{}
	if err := jsonmanager.Decode(file, &passwordCards); err != nil {
		return err
	}

	for _, passwordCard := range passwordCards {
		if passwordCard.ID == id {
			*p = passwordCard
			return nil
		}
	}

	return errors.New("not found")
}

func (r JsonPasswordCardRepository) ListPasswordCards(
	pcsPtr *[]model.PasswordCard,
	f ListPasswordCardsFilters,
) error {
	file, err := jsonmanager.Open(config.JsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := jsonmanager.Decode(file, pcsPtr); err != nil {
		return err
	}

	if f.Search == "" {
		return nil
	}

	searchRegex, err := regexp.Compile("(?i)" + f.Search)
	if err != nil {
		return err
	}

	filteredPasswordCards := []model.PasswordCard{}

	for _, passwordCard := range *pcsPtr {
		if searchRegex.MatchString(passwordCard.Name) {
			filteredPasswordCards = append(filteredPasswordCards, passwordCard)
		}
	}

	*pcsPtr = filteredPasswordCards

	return err
}
