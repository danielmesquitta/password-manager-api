package service

import (
	"testing"

	"github.com/danielmesquitta/password-manager-api/internal/config"
	"github.com/danielmesquitta/password-manager-api/internal/dto"
	"github.com/danielmesquitta/password-manager-api/internal/pkg/crypt"
	"github.com/danielmesquitta/password-manager-api/internal/pkg/validator"
	"github.com/danielmesquitta/password-manager-api/internal/repository"
)

func TestCreatePasswordCardService(t *testing.T) {
	val := validator.New()
	cryptMock := crypt.NewMock(&config.EnvVars{})
	inMemoryPasswordCardRepository := repository.NewInMemoryPasswordCardRepository()

	type args struct {
		data dto.CreatePasswordCardDTO
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should create password card",
			args: args{
				data: dto.CreatePasswordCardDTO{
					Name:     "John Doe",
					Url:      "https://www.google.com",
					Username: "john.doe",
					Password: "P@ssw0rd",
				},
			},
			wantErr: false,
		},
		{
			name: "should throw error if name is empty",
			args: args{
				data: dto.CreatePasswordCardDTO{
					Url:      "https://www.google.com",
					Username: "john.doe",
					Password: "P@ssw0rd",
				},
			},
			wantErr: true,
		},
		{
			name: "should throw error if url is empty",
			args: args{
				data: dto.CreatePasswordCardDTO{
					Name:     "John Doe",
					Username: "john.doe",
					Password: "P@ssw0rd",
				},
			},
			wantErr: true,
		},
		{
			name: "should throw error if username is empty",
			args: args{
				data: dto.CreatePasswordCardDTO{
					Name:     "John Doe",
					Url:      "https://www.google.com",
					Password: "P@ssw0rd",
				},
			},
			wantErr: true,
		},
		{
			name: "should throw error if password is empty",
			args: args{
				data: dto.CreatePasswordCardDTO{
					Name:     "John Doe",
					Url:      "https://www.google.com",
					Username: "john.doe",
				},
			},
			wantErr: true,
		},
		{
			name: "should throw error if password is less than 8 characters",
			args: args{
				data: dto.CreatePasswordCardDTO{
					Name:     "John Doe",
					Url:      "https://www.google.com",
					Username: "john.doe",
					Password: "short",
				},
			},
			wantErr: true,
		},
		{
			name: "should throw error if url is invalid",
			args: args{
				data: dto.CreatePasswordCardDTO{
					Name:     "John Doe",
					Url:      "invalid-url",
					Username: "john.doe",
					Password: "P@ssw0rd",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreatePasswordCardService(inMemoryPasswordCardRepository, cryptMock, val, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf(
					"CreatePasswordCardService() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}
