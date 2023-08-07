package service

import (
	"testing"
	"time"

	"github.com/danielmesquitta/password-manager-api/internal/config"
	"github.com/danielmesquitta/password-manager-api/internal/model"
	"github.com/danielmesquitta/password-manager-api/internal/pkg/crypt"
	"github.com/danielmesquitta/password-manager-api/internal/repository"
)

func TestListPasswordCardsService(t *testing.T) {
	cryptMock := crypt.NewMock(&config.EnvVars{})

	type args struct {
		r      repository.PasswordCardRepository
		search string
		pcsPtr *[]model.PasswordCard
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantLength int
	}{
		{
			name: "should list password cards",
			args: args{
				r: &repository.InMemoryPasswordCardRepository{
					PasswordCards: []model.PasswordCard{{
						ID:        "1",
						Name:      "John Doe",
						Username:  "john.doe",
						Password:  "hashed-password",
						Url:       "https://www.google.com",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					}},
				},
				search: "",
				pcsPtr: &[]model.PasswordCard{},
			},
			wantErr:    false,
			wantLength: 1,
		},
		{
			name: "should search and find password card",
			args: args{
				r: &repository.InMemoryPasswordCardRepository{
					PasswordCards: []model.PasswordCard{{
						ID:        "1",
						Name:      "John Doe",
						Username:  "john.doe",
						Password:  "hashed-password",
						Url:       "https://www.google.com",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					}},
				},
				search: "john",
				pcsPtr: &[]model.PasswordCard{},
			},
			wantErr:    false,
			wantLength: 1,
		},
		{
			name: "should search and not find password card",
			args: args{
				r: &repository.InMemoryPasswordCardRepository{
					PasswordCards: []model.PasswordCard{{
						ID:        "1",
						Name:      "John Doe",
						Username:  "john.doe",
						Password:  "hashed-password",
						Url:       "https://www.google.com",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					}},
				},
				search: "search-with-no-results",
				pcsPtr: &[]model.PasswordCard{},
			},
			wantErr:    false,
			wantLength: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ListPasswordCardsService(tt.args.r, cryptMock, tt.args.search, tt.args.pcsPtr); (err != nil) != tt.wantErr {
				t.Errorf(
					"ListPasswordCardsService() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}

			if length := len(*tt.args.pcsPtr); length != tt.wantLength {
				t.Errorf(
					"ListPasswordCardsService() length = %v, wantLength %v",
					length,
					tt.wantLength,
				)
			}
		})
	}
}
