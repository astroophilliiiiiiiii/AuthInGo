package services

import (
	"AuthInGo/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//Implementing JWT (JSON Web Token) authentication in Golang allows you to
// securely manage user login sessions without storing state on the server.

//JWT → user identify karne ke liye only id/email

func GenJwtToken(i *models.User) (string, error) {
	// NewWithClaims creates a new Token with the specified signing method and claims.
	// MapClaims is a claims type that uses the mapstringany for JSON decoding.
	// This is the default claims type if you don't supply one
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       i.Id,
		"username": i.Username,
		"email":    i.Email,
		//"password": i.Password,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Tera secret key nikalne wala code yahan hoga...
	secretKey := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secretKey))

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
