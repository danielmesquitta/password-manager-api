package service

import (
	"testing"

	"github.com/danielmesquitta/password-manager-api/internal/config"
	"github.com/danielmesquitta/password-manager-api/internal/dto"
	"github.com/danielmesquitta/password-manager-api/internal/model"
	"github.com/danielmesquitta/password-manager-api/internal/pkg/crypt"
	"github.com/danielmesquitta/password-manager-api/internal/pkg/validator"
	"github.com/danielmesquitta/password-manager-api/internal/repository"
)

func TestUpdatePasswordCardService(t *testing.T) {
	val := validator.New()
	cMock := crypt.NewMock(&config.EnvVars{})

	type args struct {
		r    repository.PasswordCardRepository
		id   string
		data dto.UpdatePasswordCardDTO
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should update password card",
			args: args{
				r: &repository.InMemoryPasswordCardRepository{
					PasswordCards: []model.PasswordCard{
						{
							ID:       "1",
							Name:     "John Doe",
							Url:      "https://www.google.com",
							Username: "john.doe",
							Password: "hashed-password",
						},
					},
				},
				id: "1",
				data: dto.UpdatePasswordCardDTO{
					Name:     "Jane Doe",
					Url:      "https://www.facebook.com",
					Username: "jane.doe",
					Password: "P@ssw0rd",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdatePasswordCardService(tt.args.r, cMock, val, tt.args.id, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf(
					"UpdatePasswordCardService() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}
