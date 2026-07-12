package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plainpassword string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(plainpassword), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error hashing the password ! ")
		return "", err
	}

	HashPassword := string(hash) // converting array of bytes back to string

	return HashPassword, nil
}

func CheckHashPassword(plainpass string, haspass string) bool {
	// to check if these both r equal or not -- galat toh ni hash krdiyaa
	err := bcrypt.CompareHashAndPassword([]byte(haspass), []byte(plainpass))

	return err == nil
}
