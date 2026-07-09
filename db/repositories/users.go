package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

// a struct that creates a user entry in the database

type UserRespository interface {
	GetById() (*models.User, error)
	Create() error
}

// actual -- that will talk to the database
// implements UserRepository interface
type UserRespositoryImpl struct {

	// every repo should have a db property
	// it needs to make sure that they r actually making the db calls
	db *sql.DB // gives sql db instance -- using this we'll be able to make db queries
}

// constructor
func NewUserRepository(_db *sql.DB) UserRespository {
	return &UserRespositoryImpl{
		db: _db,
	}
}

// member functions of Impl kaa hainaa  -- implements the interface
func (u *UserRespositoryImpl) GetById() (*models.User, error) {
	fmt.Println("Getting user in UserRepository ")

	// Step-1 Prepare the query  -- to fetch 1 single row
	query := "SELECT id , username , email , password , created_at , updated_at FROM users WHERE id=?"

	// ? is created to avoid the sql injection by the hackerss !!

	// Step-2 Execute the query
	row := u.db.QueryRow(query, 1) // db is the object that is connected to db -- on that query is made !!

	// Step-3 Process the result
	user := &models.User{} // Empty user object -- eventually that we need !

	//Step 4: Scan result -- Scan database columns ko struct fields me copy karta h
	// - here we have given the desitanation in the bracket ( row will be scanned and put in this empty object )
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	//Step 5: Error handling
	if err != nil {
		if err == sql.ErrNoRows { // row doesnt existss !!
			fmt.Println("No user found with given ID")
			return nil, err
		} else { // any other issue exitss
			fmt.Println("Error scanning user: ", err)
			return nil, err
		}
	}

	fmt.Println("User fetched successfully: ", user)

	return user, nil
}

func (u *UserRespositoryImpl) Create() error {

	// Step-1 Prepare the query
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	// Step-2 Execute the query
	result, err := u.db.Exec(query, "TestUser", "test@gmail.com", "123") // Exec -- doesnt return any row

	if err != nil {
		fmt.Println("Error returning the row!!")
		return err
	}

	// in sql -- when we insert something -- it says naa rows affected !!
	rowsAffected, rowErr := result.RowsAffected()

	if rowErr != nil {
		fmt.Println("Error getting in rows affected:", rowErr)
		return rowErr
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected , user not created! ")
		return rowErr
	}

	fmt.Println("User created successfully , rows affected: ", rowsAffected)

	return nil
}
