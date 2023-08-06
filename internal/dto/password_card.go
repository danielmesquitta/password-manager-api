package dto

type CreatePasswordCardDTO struct {
	Name     string `json:"name"     validate:"required"`
	Url      string `json:"url"      validate:"required,url"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type UpdatePasswordCardDTO struct {
	Name     string `json:"name"`
	Url      string `json:"url"      validate:"url"`
	Username string `json:"username"`
	Password string `json:"password" validate:"min=8"`
}
