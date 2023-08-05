package crypt

import (
	"os"
)

var secret = ""

func Init() {
	secret = os.Getenv("SECRET")

	switch length := len(secret); {
	case length >= 32:
		secret = secret[:32]
	case length >= 24:
		secret = secret[:24]
	case length >= 16:
		secret = secret[:16]
	default:
		panic("SECRET must be at least 16 characters long")
	}
}
