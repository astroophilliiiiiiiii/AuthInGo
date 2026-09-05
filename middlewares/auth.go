package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// jwt token will come like ---  Authorisation : Bearer token

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Authorisation header is required ", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Authorisation header must start with the Bearer", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			http.Error(w, "Token is required", http.StatusUnauthorized)
			return
		}

		//---------------------------------to check the validity of the token now -------------------------------------------
		claims := jwt.MapClaims{} // declared dictionary that has valid funtion too -- that cn check expiry

		_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			//return []byte(env.GetString("JWT_SECRET", "TOKEN")), nil

			// 1. os.Getenv se secret nikal
			secretKey := os.Getenv("JWT_SECRET")

			// 2. Agar .env me nahi mili, toh default "TOKEN" set kar de
			if secretKey == "" {
				secretKey = "TOKEN"
			}

			// 3. Phir bytes me return kar de
			return []byte(secretKey), nil
		})

		if err != nil {
			fmt.Println("PARSING ERROR:", err)
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// jb v koin number aata hai hum claims mein daalte parse krvaake toh vo float mein hi ata hai
		userID, okID := claims["id"].(float64)
		email, okEmail := claims["email"].(string)

		if !okID || !okEmail {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		fmt.Println("Authenticated user ID :", userID, "Email:", email)

		ctx := context.WithValue(r.Context(), "userID", int64(userID))
		ctx = context.WithValue(ctx, "email", email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
