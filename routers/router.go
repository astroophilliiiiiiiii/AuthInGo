package routers

// Dependency Injection + Separation of Responsibility

import (
	"AuthInGo/controllers"
	// "AuthInGo/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router interface {
	// "Jo bhi struct Register() function implement karega, vo Router maana jayega."
	Register(r chi.Router)
}

// setting up a chi router
// Chi is a library in Go that helps you create API routes
// 1. this will be called when app started
// Mujhe koi bhi router dedo jiske paas Register method ho
func SetUpRouter(UserRouter Router) *chi.Mux {

	// setting up a new chi router -- in-built support
	chirouter := chi.NewRouter() // this is chi router inbuilt function for making a router
	// route added to this router

	//chirouter.Use(middlewares.RequestLogger) // that will ne registered for all set of below requests
	chirouter.Use(middleware.Logger) // built-in middleware

	chirouter.Get("/ping", controllers.PingHandler) // passing to controllers layer

	//"UserRouter, ye mera empty chi router le aur iske andar apne user wale routes daal de."
	//Ab chirouter UserRouter ko pass kar rahe ho taaki vo uspe apne routes attach kar de.
	UserRouter.Register(chirouter) // calling the function of the UserRouter

	return chirouter
}

// Example maan lo UserRouter ke andar:

// func (u *UserRouter) Register(r chi.Router) {
//     r.Post("/users", u.UserController.CreateUser)
//     r.Get("/users", u.UserController.GetUser)
// }

// Ab jab call hua:

// UserRouter.Register(chirouter)

// actually ye hua:

// chirouter.Post("/users", ...)
// chirouter.Get("/users", ...)

// same object modify ho gaya.

// Pehle:

// chirouter    /ping

// Register ke baad:

// chirouter  /ping  /users POST   /users GET

//UserRouter     -> user routes register karega
// HotelRouter    -> hotel routes register karega
// PaymentRouter  -> payment routes register

// Ye mera chi router hai, tum log apne-apne routes attach kar do
