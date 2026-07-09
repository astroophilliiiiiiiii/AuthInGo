package db

// db related configurationss

import (
	env "AuthInGo/config/env" // jo file use krni h aage usko mention
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func SetupDb() (*sql.DB, error) {

	// Capture connection properties. -- documentation of the go.dev -- sql connection
	cfg := mysql.NewConfig()
	cfg.User = env.GetString("DB_USER", "root")
	cfg.Passwd = env.GetString("DB_PASSWORD", "root")
	cfg.Net = env.GetString("DB_NET", "tcp")
	cfg.Addr = env.GetString("DB_ADDR", "127.0.0.1:3306")
	cfg.DBName = env.GetString("DBname", "airbnb_dev")

	fmt.Println("Connecting to the database", cfg.DBName, cfg.FormatDSN())

	// Get a database handle.
	// FormatDSN formats the given Config into a DSN string which can be passed to the driver.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("Error connecting to the database", err)
		return nil, err // no db object to be returned
	}

	fmt.Println("Trying to connect to the database ")
	pingErr := db.Ping() // will ping the db to check if everything's fine or not
	if pingErr != nil {
		fmt.Println("Error connecting to the database", pingErr)
		return nil, pingErr // no db object to be returned
	}

	fmt.Println("Connected to the database:- ", cfg.DBName)

	return db, nil // no error to return

}
