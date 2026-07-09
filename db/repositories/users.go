package db

import (
	"fmt"
)

// a struct that creates a user entry in the database

type UserRespository interface {
	Create() error
}

// actual -- that will talk to the database
// implements UserRepository interface
type UserRespositoryImpl struct {

	// every repo should have a db property
	// it needs to make sure that they r actually making the db calls
	// db *sql.DB // gives sql db instance -- using this we'll be able to make db queries
}

// constructor
func NewUserRepository() UserRespository {
	return &UserRespositoryImpl{
		//db: db,
	}
}

// member functions of Impl kaa hainaa  -- implements the interface
func (u *UserRespositoryImpl) Create() error {
	fmt.Println("Creating user in UserRepository ")
	return nil
}
