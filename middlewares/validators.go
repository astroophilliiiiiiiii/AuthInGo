package middlewares

import (
	"AuthInGo/dto"
	"AuthInGo/utils"
	"context"
	"net/http"
)

func UserLoginRequestValidator(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // wrapping a func into middleware

		var payload dto.LoginUserRequestDTO // as isko payload m bhrna haina toh valid type dto toh dena padega na

		// read and decode json body into the payload
		err := utils.ReadJsonBody(r, &payload)
		if err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}

		//Validate the payload using Validator instance
		verr := utils.Validator.Struct(payload)

		if verr != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation failed", verr)
			return
		}

		// passing the context to the controllers so that its possible to read the payload again from request
		req_context := r.Context() // original context coming for the current request

		ctx := context.WithValue(req_context, "payload", &payload) // create a new context with the payload

		next.ServeHTTP(w, r.WithContext(ctx)) // naya request bnaake bhejdia with this new context
	})
}

func UserCreateRequestValidator(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // wrapping a func into middleware

		var payload dto.CreateUserRequestDto // as isko payload m bhrna haina toh valid type dto toh dena padega na

		// read and decode json body into the payload
		err := utils.ReadJsonBody(r, &payload)
		if err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}

		//Validate the payload using Validator instance
		verr := utils.Validator.Struct(payload)

		if verr != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation failed", verr)
			return
		}

		req_context := r.Context()

		ctx := context.WithValue(req_context, "payload", &payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
