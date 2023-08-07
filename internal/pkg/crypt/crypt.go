package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"

	"github.com/danielmesquitta/password-manager-api/internal/config"
)

type Crypter interface {
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
}

type Crypt struct {
	secret string
}

func New(env *config.EnvVars) *Crypt {
	return &Crypt{secret: env.HashSecret}
}

func (c *Crypt) Decrypt(encryptedString string) (string, error) {
	key, err := hex.DecodeString(c.secret)
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.URLEncoding.DecodeString(encryptedString)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		err = errors.New("ciphertext too short")
		return "", err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	decrypted := string(ciphertext)

	return decrypted, nil
}

func (c *Crypt) Encrypt(plaintext string) (string, error) {
	// convert key to bytes
	key, err := hex.DecodeString(c.secret)
	if err != nil {
		return "", err
	}

	plaintextInBytes := []byte(plaintext)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintextInBytes))
	iv := ciphertext[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintextInBytes)

	// convert to base64
	encryptedString := base64.URLEncoding.EncodeToString(ciphertext)

	return encryptedString, nil
}
