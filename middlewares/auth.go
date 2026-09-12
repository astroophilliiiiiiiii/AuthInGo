package middlewares

import (
	dbConfig "AuthInGo/config/db"
	repo "AuthInGo/db/repositories"
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
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

// token se user ID nikaalna, DB se uske roles verify karna,
// aur check fail hone par request ko wahin terminate karna.
func RequireAllRoles(roles ...string) func(http.Handler) http.Handler {
	// fxn return jo aage http.Handler hi return krraa hai

	// function that can create a middleware for checking the above set of roles

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			userIdStr := r.Context().Value("userID").(string)
			userId, err := strconv.ParseInt(userIdStr, 10, 64)
			if err != nil {
				http.Error(w, "Invalid user ID", http.StatusUnauthorized)
				return
			}

			dbConn, dbErr := dbConfig.SetupDb()
			if dbErr != nil {
				http.Error(w, "Database connection error: "+dbErr.Error(), http.StatusInternalServerError)
				return
			}

			urr := repo.NewUserRoleRepository(dbConn)

			hasAllRoles, hasAllRolesErr := urr.HasAllRoles(userId, roles)
			fmt.Println("userid", userId, "roles", roles, "hasAllRoles", hasAllRoles)
			if hasAllRolesErr != nil {
				http.Error(w, "Error checking user roles: "+hasAllRolesErr.Error(), http.StatusInternalServerError)
				return
			}

			if !hasAllRoles {
				http.Error(w, "Forbidden: You do not have the required roles", http.StatusForbidden)
				return
			}

			fmt.Println("User has all required roles:", roles)

			next.ServeHTTP(w, r)
		})
	}

}
