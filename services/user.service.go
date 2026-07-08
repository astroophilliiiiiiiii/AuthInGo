package services

import (
	db "AuthInGo/db/repositories"
	"fmt"
)

// -- controllers -- services -- repository

type UserService interface {
	CreateUser() error
}

type UserServiceImpl struct {
	// depending on the repo interface
	userRepository db.UserRespository // declared -- passed in the constructor
}

// constructor
func NewUserService(_userRespository db.UserRespository) UserService {
	return &UserServiceImpl{ // abhi isko UserServiceImpl ko UserService kaa type same krna padega
		userRepository: _userRespository,
	}
}

// member function
func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("Creating user in UserService")

	// call the repo layer -- of this Service  that was passed while creating this service
	// and in actual only the repo will be passed -- backend logic is on the interface
	u.userRepository.Create()

	return nil
}
