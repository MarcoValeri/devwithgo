package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(getPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(getPassword), 14)
	if err != nil {
		fmt.Println("Error to generate hash:", err)
	}
	return string(bytes), err
}

func PasswordHashChecker(getPassword, getHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(getHash), []byte(getPassword))
	if err != nil {
		fmt.Println("Error checking hash password:", err)
	}
	return err == nil
}
