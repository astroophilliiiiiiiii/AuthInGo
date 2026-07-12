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
func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Registered user called in UserController")

	// call the service layer
	uc.UserService.CreateUserService(r) // call the createuser function

	// send the response
	w.Write([]byte("User creation done !! "))
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Login user called in UserController")
	err := uc.UserService.LogInUser() // user service ka hainaa login function

	if err != nil {
		http.Error(w, "Login failed", http.StatusInternalServerError)
	}
	w.Write([]byte("User login endpoint done"))
}
