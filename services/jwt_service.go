package services

import (
	"AuthInGo/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//Implementing JWT (JSON Web Token) authentication in Golang allows you to
// securely manage user login sessions without storing state on the server.

//JWT → user identify karne ke liye only id/email

// 📌📌 TODO -- create this in the env package
var secretKey = []byte("secret-key") // := sirf function ke andar allowed hota hai ✅
// Global/package variables → var

func GenJwtToken(i *models.User) (string, error) {
	// NewWithClaims creates a new Token with the specified signing method and claims.
	// MapClaims is a claims type that uses the mapstringany for JSON decoding.
	// This is the default claims type if you don't supply one
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": i.Username,
		"email":    i.Email,
		//"password": i.Password,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	fmt.Println("Successfully created JWT token ! ")

	return tokenString, nil
}

// JWT encrypted nahi hota, sirf signed hota hai ✅
// Matlab koi bhi JWT ka payload decode karke dekh sakta hai.
// toh hashed password bhi visible ho jayega ❌
// Hash ko reverse karna mushkil hai, but phir bhi unnecessary risk hai.
