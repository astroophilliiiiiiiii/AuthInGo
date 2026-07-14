package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/services"
	"AuthInGo/utils"
	"encoding/json"
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
	fmt.Println("CreateUser called in UserController")

	// call the service layer
	uc.UserService.CreateUserService(r) // call the createuser function

	// send the response -- JSON format
	s := "User creation done !! "
	jsonStr, err := json.Marshal(s)
	if err != nil {
		fmt.Println("Error in converting to json reponse ")
		return
	}

	w.Write(jsonStr)
}
func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login user called in UserController")

	var payload dto.LoginUserRequestDTO // dto type ke object mein payload fill

	if jsonerr := utils.ReadJsonBody(r, &payload); jsonerr != nil { // shortcut to handle the error
		w.Write([]byte("Something went wrong while logging in"))
		return
	}

	if validationerr := utils.Validator.Struct(payload); validationerr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid input data", validationerr)
		return
	}

	token, err := uc.UserService.LogInUser(&payload) // user service -- contains logic of hwt and all --validating

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Login failed", err)
		return
	}

	// send the response -- JSON format
	utils.WriteJsonSuccessResponse(w, http.StatusOK, token, "User logged in sucessfully ! ")
}
