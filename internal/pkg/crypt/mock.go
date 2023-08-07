package crypt

import "github.com/danielmesquitta/password-manager-api/internal/config"

type MockCrypt struct{}

func NewMock(_ *config.EnvVars) *MockCrypt {
	return &MockCrypt{}
}

func (c *MockCrypt) Encrypt(plaintext string) (string, error) {
	return plaintext, nil
}

func (c *MockCrypt) Decrypt(encryptedString string) (string, error) {
	return encryptedString, nil
}
