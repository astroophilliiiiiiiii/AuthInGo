package controllers

import (
	"AuthInGo/services"
	"fmt"
	"net/http"
)

// no interface -- as it only does one work of passing the logic to the service layer --- passer !!

// it will connect to the service layer interface
type UserController struct {
	UserService services.UserService
}

// constructor in which the service layer will be passed
func NewUserController(__userService services.UserService) *UserController {
	return &UserController{
		UserService: __userService,
	}
}

// member function
func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Registered user called in UserController")
	// call the service layer
	uc.UserService.CreateUser() // call the createuser function

	// send the response
	w.Write([]byte("User Registration Endpoint"))
}
