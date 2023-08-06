package crypt

type MockCrypt struct{}

func NewMock() *MockCrypt {
	return &MockCrypt{}
}

func (c *MockCrypt) Encrypt(plaintext string) (string, error) {
	return plaintext, nil
}

func (c *MockCrypt) Decrypt(encryptedString string) (string, error) {
	return encryptedString, nil
}
