package routers

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	// any router will be identified only if it implements the register method
	Register(r chi.Router)
}

// setting up a chi router
// Chi is a library in Go that helps you create API routes
// 1. this will be called when app started
func SetUpRouter(UserRouter Router) *chi.Mux {

	// setting up a new chi router -- in-built support
	chirouter := chi.NewRouter() // this is chi router inbuilt function for making a router
	// route added to this router
	chirouter.Get("/ping", controllers.PingHandler) // passing to controllers layer

	UserRouter.Register(chirouter)

	return chirouter
}
