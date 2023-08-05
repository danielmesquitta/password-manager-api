package response

import "github.com/danielmesquitta/password-manager-api/internal/model"

type ListPasswordCardsResponse struct {
	Data []model.PasswordCard `json:"data"`
}
