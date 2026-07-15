package utils

// JSON MARSHALLING
// consumedby the controllers layer to send the responses
import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate //Validator naam ka ek variable bana diya.
// Yaani ye validator object ka pointer rakhegaa

func init() { //init() Go ka special function hai. ✅ Isko tum kabhi call nahi karti. Go khud call karta hai.
	Validator = NewValidator()
}

func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled()) //Yaha actual validator object ban raha hai.
	//Ye validator ki setting hai. Matlab validator ko bol rahe ho: "Required fields ko properly check karna."
	// u have to do the validation with the validators inside the struct
}

func WriteJsonSuccessResponse(w http.ResponseWriter, status int, data any, message string) error {
	response := map[string]any{
		"message": message,
		"data":    data,
		"status":  status,
	}

	return WriteJsonResponse(w, status, response)
}

func WriteJsonErrorResponse(w http.ResponseWriter, status int, message string, err error) error {
	response := map[string]any{
		"error":   err.Error(),
		"status":  status,
		"message": message,
	}

	return WriteJsonResponse(w, status, response)
}

// w -- with help of it we can write a JSON response to the frontend
// data -- which is needed to converted into the JSON
func WriteJsonResponse(w http.ResponseWriter, status int, data any) error {
	// 👉 Browser/frontend ko bata diya ki JSON aa raha hai.
	w.Header().Set("Content-Type", "application/json") // set the current type to application/json

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data) // data ko JSON me convert karke direct frontend ko bhej diya.
	//Jo data aa raha hai, usse JSON ki tarah parse karna ✅
	//client ko bataya jaata hai ki usse kaise interpret (samajhna) h
}

// result → jis struct me data bharna hai
// error → agar JSON galat hua toh error return karega
func ReadJsonBody(r *http.Request, result any) error {
	decoder := json.NewDecoder(r.Body) // Request body ko JSON reader bana diya.

	decoder.DisallowUnknownFields() // Agar user extra field bhej de jo struct me nahi hai, toh error de do

	return decoder.Decode(result) // JSON ko result struct me bhar do.
}

// Return type error kyu? --  Kyuki Decode() khud error return karta hai.
// Agar:
// if JSON is invalid❌
// Type mismatch is there❌
// Unknown field is there (DisallowUnknownFields) ❌
// then error will come
