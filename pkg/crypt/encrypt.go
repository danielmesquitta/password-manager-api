package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

func (c *Crypt) Encrypt(plaintext string) (string, error) {
	// convert key to bytes
	key, err := hex.DecodeString(secret)
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
