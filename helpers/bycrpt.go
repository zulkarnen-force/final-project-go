package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPass(p string) string {
	salt := 8
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)
	return string(hash)
}

func ComparePassword(h, p string) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	if err != nil {
		return false
	}

	return true
}