package controllers

// handler functions

import (
	"net/http"
)

// Request = user se aaya hua data (URL, headers, body)
// bada struct hai, copy avoid karne ke liye pointer use hota hai (*Request)

// ResponseWriter = response bhejne ka tool
// ye interface hai, internally already reference jaisa behave karta hai
// isliye *ResponseWriter nahi likhte

// ResponseWriter se hum likhte hain → no pointer ✅
// Request ko hum receive karte hain → pointer

func PingHandler(w http.ResponseWriter, r *http.Request) {

	// w = response bhejne wala pen/tool ✏️   Write() = us pen se user ko data likh ke bhejna
	// HTTP response body me data bytes form me send hota hai.
	w.Write([]byte("ping"))
}
