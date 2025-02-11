package hash

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// MustGenerate hashes the given string.
// If the string is greater than 72 bytes, it truncates and hashes only the first 72 bytes.
// It panics if any error occurs.
func MustGenerate(s string) string {
	toHash := []byte(s)

	h, err := generate(toHash, 4)
	if err != nil {
		if !errors.Is(err, bcrypt.ErrPasswordTooLong) {
			panic(err)
		}

		h, _ = generate(toHash[:72], 4)
	}

	return h
}

func Generate(secret string) (string, error) {
	return generate([]byte(secret), 4)
}

func generate(s []byte, cost int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(s, cost)
	return string(hash), err
}

func Validate(hash, value string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
	return err == nil
}
