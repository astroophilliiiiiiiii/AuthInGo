package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

// a struct that creates a user entry in the database

type UserRespository interface {
	GetById() (*models.User, error)
	Create(username *string, email *string, password *string) error
	GetAll() ([]*models.User, error) // 📌⌛ should return array of objects
	DeleteById(id int64) error       // 📌⌛ should take an id parameter -- delete the row
	GetByEmail(email string) (*models.User, error)
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

func (u *UserRespositoryImpl) Create(un *string, em *string, hashpass *string) error {

	// Step-1 Prepare the query
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	// Step-2 Execute the query
	result, err := u.db.Exec(query, *un, *em, *hashpass) // Exec -- doesnt return any row -- affecting the db

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

func (u *UserRespositoryImpl) GetAll() ([]*models.User, error) {
	// Step-1 Prepare the query
	query := "SELECT id, username, email, password, created_at, updated_at FROM users"

	// Step-2 Execute the query
	rows, err := u.db.Query(query)

	if err != nil {
		fmt.Println("Error fetching the rows!!")
		return nil, nil
	}

	// Step-3 Process the result
	users := []*models.User{} // as there is array of users

	for rows.Next() { // rows ke andar ek internal pointer/cursor hota hai ✅
		// next krne se pointer mover to the next
		user := &models.User{}

		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, nil // stops the flow -- nil returned !!
		}

		users = append(users, user)
	}

	fmt.Println("All users fetched successfullyy !! ")

	return users, nil
}

// like CREATE a row -----
func (u *UserRespositoryImpl) DeleteById(id int64) error {

	// Step-1 Prepare the query  -- to delete 1 row
	query := "DELETE FROM users WHERE id=?"

	// Step-2 Execute the query
	result, err := u.db.Exec(query, id)

	if err != nil {
		fmt.Println("Error deleting the row")
		return err
	}

	// Step-3 Check deleted or not  -- Process the result
	rowsAffected, rowErr := result.RowsAffected()

	if rowErr != nil {
		fmt.Println("Error getting in rows affected:", rowErr)
		return rowErr
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected , user not deleted! ")
		return rowErr
	}

	fmt.Println("User deleted successfully , rows affected: ", rowsAffected)

	return nil
}

// fetching by the email
func (u *UserRespositoryImpl) GetByEmail(email string) (*models.User, error) {
	fmt.Println("Getting user in UserRepository ")

	// Step-1 Prepare the query  -- to fetch 1 single row
	query := "SELECT id , username , email , password , created_at , updated_at FROM users WHERE email=?"

	// ? is created to avoid the sql injection by the hackerss !!

	// Step-2 Execute the query
	row := u.db.QueryRow(query, email)

	// Step-3 Process the result
	user := &models.User{}

	//Step 4: Scan result -- Scan database columns ko struct fields me copy karta h
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
