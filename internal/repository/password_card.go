package repository

import "github.com/danielmesquitta/password-manager-api/internal/model"

type PasswordCardRepository interface {
	CreatePasswordCard(data model.PasswordCard) error

	UpdatePasswordCard(id string, data model.PasswordCard) error

	DeletePasswordCard(id string) error

	GetPasswordCard(id string, pcPtr *model.PasswordCard) error

	ListPasswordCards(search string, pcsPtr *[]model.PasswordCard) error
}
