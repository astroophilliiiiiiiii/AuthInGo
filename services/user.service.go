package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// -- controllers -- services -- repository

type UserService interface {
	CreateUserService(r *http.Request) error
	LogInUser() (string, error)
}

type UserServiceImpl struct {
	// depending on the repo interface
	userRepository db.UserRespository // declared -- passed in the constructor
}

// constructor - isme pass hori actual repo layerr
func NewUserService(_userRespository db.UserRespository) UserService {
	return &UserServiceImpl{ // abhi isko UserServiceImpl ko UserService kaa type same krna padega
		userRepository: _userRespository,
	}
}

type Inputbody struct {
	Username string // capital needed -- as when decoded -- in diff package these variables will not be accessible
	Email    string
	Password string
}

// member function
func (u *UserServiceImpl) CreateUserService(r *http.Request) error {
	fmt.Println("Getting user in UserService")

	// fetching the data from the req.body and encrypting the password and forwarding it
	var sample Inputbody
	err := json.NewDecoder(r.Body).Decode(&sample) // decode accepts pointer
	// r.Body = raw JSON data
	// NewDecoder := made the json readable
	// Decode reads the next JSON-encoded value from its input & stores the in sample

	if err != nil {
		fmt.Println("Error reading the data !! ")
	}

	sample.Password, err = utils.HashPassword(sample.Password) // its the hashed password

	if err != nil {
		fmt.Println("Error hashing")
		return err
	}

	// call the repo layer -- of this Service  that was passed while creating this service
	// and in actual only the repo will be passed -- backend logic is on the interface

	// generating the JWT token
	// token , err  := GenJwtToken(&sample)

	u.userRepository.Create(&sample.Username, &sample.Email, &sample.Password)

	return nil
}

func (u *UserServiceImpl) LogInUser() (string, error) {
	// email and password r given as parameters
	email := "kriti@gmail.com"
	password := "testpassword"

	// Step-1 -- make a repo call to fetch the user by email
	user, err := u.userRepository.GetByEmail(email)

	// Step-2 -- if error exists
	if err != nil {
		fmt.Println("User doesnt exists !! ")
		return "", err
	}

	// Step-3 -- check the password using checkhash
	response := utils.CheckHashPassword(password, user.Password)

	if response == false {
		fmt.Println("Password doesnt matches !! ")
		return "", err
	}

	// Step-4 -- print JWT token
	token, err := GenJwtToken(user)

	if err != nil {
		fmt.Println("Error fetching the token!! ")
		return "", err
	}

	fmt.Println("Token is:-- ", token)

	return token, nil
}
