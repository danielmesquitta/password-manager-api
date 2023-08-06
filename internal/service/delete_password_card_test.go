package service

import (
	"testing"
	"time"

	"github.com/danielmesquitta/password-manager-api/internal/model"
	"github.com/danielmesquitta/password-manager-api/internal/repository"
)

func TestDeletePasswordCardService(t *testing.T) {
	type args struct {
		r  repository.PasswordCardRepository
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should delete password card",
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
				id: "1",
			},
			wantErr: false,
		},
		{
			name: "should throw error if not exists",
			args: args{
				r:  repository.NewInMemoryPasswordCardRepository(),
				id: "1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeletePasswordCardService(
				tt.args.r,
				tt.args.id,
			); (err != nil) != tt.wantErr {
				t.Errorf(
					"DeletePasswordCardService() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}
