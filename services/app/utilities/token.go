package utilities

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateSecureToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func MustGenerateSecureToken(length int) string {
	token, err := GenerateSecureToken(length)
	if err != nil {
		panic(token)
	}
	return token
}
