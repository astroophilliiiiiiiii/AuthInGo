package config // as in the config folder

// .env values load and give -- { functions }

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// .env loaded
func Load() {
	// return error object
	err := godotenv.Load()

	if err != nil { // error occured in loading
		fmt.Printf("Error loading the env file ")
	}
}

// will return string based env value for the string key
func GetString(key string, fallback string) string {

	// os package by go --
	// populates the env variable -- that is either in the system (done by terminal ) || .env file
	//LookupEnv retrieves the value of the environment variable named by the key.
	// If the variable is present in the environment the value (which may be empty) is returned
	//and the boolean is true. Otherwise the returned value will be empty and the boolean will be false.
	value, ok := os.LookupEnv(key) // in go , a func can return 2 values

	if !ok {
		return fallback
	}

	return value
}

// if in .env file its the int not the string value
func GetInt(key string, fallback int) int {

	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	// manually do the conversion of value to int
	intValue, err := strconv.Atoi(value)

	if err != nil {
		fmt.Printf("Error converting %s into int: %v\n", key, err)
	}

	return intValue
}

func GetBool(key string, fallback bool) bool {

	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	// manually do the conversion of value to int
	boolValue, err := strconv.ParseBool(value)

	if err != nil {
		fmt.Printf("Error converting %s into int: %v\n", key, err)
	}

	return boolValue
}
