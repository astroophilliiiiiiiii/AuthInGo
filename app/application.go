package app

import (
	config "AuthInGo/config/env"
	"AuthInGo/controllers"
	db "AuthInGo/db/repositories"
	"AuthInGo/routers"
	"AuthInGo/services"
	"fmt"
	"net/http" // built-in http package for the server
	"time"
)

// .env -->  Config struct -->  Application --> Server

// Application ki saari settings ek jagah store karne ke liye. ✅
type Config struct {
	Addr string // address -- PORT
}

// constructor for it
func NewConfig() *Config {
	port := config.GetString("PORT", ":8080") // loading port from .env

	return &Config{
		Addr: port,
	}
}

// app -- instance of server
type Application struct {
	Config Config
	Store  db.Storage // ye interface btaayaa actual constructor m pass krrenge
}

// constructor
func NewApp(cfg *Config) *Application {
	return &Application{
		Config: *cfg,             // ye bahr bnraa h  -- isko v NewConfig() kr sktee thee !!
		Store:  *db.NewStorage(), // ye andr hi bnaadiyaa
	}
}

// member function   -- here app.run()
func (app *Application) Run() error {

	// all connection will happen here
	ur := db.NewUserRepository()
	us := services.NewUserService(ur)
	uc := controllers.NewUserController(us)
	uRouter := routers.NewUserRouter(uc)

	// server object created -- reference to it returned
	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      routers.SetUpRouter(uRouter), // setup chi router and put it here
		ReadTimeout:  10 * time.Second,             // req recieve krne kaa time // starts when client connects
		WriteTimeout: 10 * time.Second,             // response send krne ka time -- server res likhna start krrega tb timie chalegaa
		IdleTimeout:  60 * time.Second,             // connection client stops
	}

	fmt.Println("Starting server on ", app.Config.Addr)

	return server.ListenAndServe() // server started
}
