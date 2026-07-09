package main

import (
	"AuthInGo/app"
	dbConfig "AuthInGo/config/db"
	config "AuthInGo/config/env"
)

func main() {

	// 1. loaded the env variables
	config.Load()

	// 2.config object made -- its configurations !!
	cfg := app.NewConfig()

	// new app instance made
	app := app.NewApp(cfg)

	dbConfig.SetupDb()
	// run the server
	app.Run()
}
