package controllers

import (
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsEmailValid(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
			return "", false
	}
	return addr.Address, true
}