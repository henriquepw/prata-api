package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func Generate(secret string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(secret), 10)
	return string(hash), err
}

func Validate(hash, value string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
	return err == nil
}
