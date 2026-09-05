package middlewares

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// time rate (limit) , no. of requests
var limiter = rate.NewLimiter(rate.Every(1*time.Minute), 5) // 5 requests at max per second

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			//Allow() check karta hai: Token available? ✅ → true  Token nahi? ❌ → false Aur agar token mila, to 1 token consume bhi kar deta hai.
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
