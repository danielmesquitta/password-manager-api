package model

import (
	"time"

	_ "github.com/go-playground/validator/v10"
)

type PasswordCard struct {
	ID        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Name      string    `json:"name,omitempty"`
	Url       string    `json:"url,omitempty"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
}
