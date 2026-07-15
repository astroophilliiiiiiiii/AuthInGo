package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/models"
	"AuthInGo/utils"
	"fmt"
)

// -- controllers -- services -- repository

type UserService interface {
	CreateUserService(payload dto.CreateUserRequestDto) (*models.User, error)
	LogInUser(payload *dto.LoginUserRequestDTO) (string, error)
	GetUserByIdService(id *int) (*models.User, error)
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
func (u *UserServiceImpl) CreateUserService(payload dto.CreateUserRequestDto) (*models.User, error) {
	fmt.Println("Getting user in UserService")

	hashpassword, err := utils.HashPassword(payload.Password) // its the hashed password

	if err != nil {
		fmt.Println("Error hashing")
		return nil, err
	}

	// call the repo layer -- of this Service  that was passed while creating this service
	// and in actual only the repo will be passed -- backend logic is on the interface

	// generating the JWT token
	// token , err  := GenJwtToken(&sample)

	user, err1 := u.userRepository.Create(&payload.Username, &payload.Email, &hashpassword)

	if err1 != nil {
		fmt.Println("Error creating the user ")
		return nil, err1
	}

	return user, nil
}

func (u *UserServiceImpl) GetUserByIdService(id *int) (*models.User, error) {
	fmt.Println("Getting user in UserService")

	user, err := u.userRepository.GetById(id)

	if err != nil {
		fmt.Println("User doesnt exits")
		return nil, err
	}

	return user, nil
}

func (u *UserServiceImpl) LogInUser(payload *dto.LoginUserRequestDTO) (string, error) {
	// email and password r given as parameters
	email := payload.Email
	password := payload.Password

	// Step-1 -- make a repo call to fetch the user by email
	user, err := u.userRepository.GetByEmail(&email) // needs * so pass &

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
