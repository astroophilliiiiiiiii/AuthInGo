package middlewares

import (
	"fmt"
	"net/http"
)

func RequestLogger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // wrapping a func into middleware

		fmt.Println("middleware :- Received request:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // Call the next handler in the chain

	})
}
