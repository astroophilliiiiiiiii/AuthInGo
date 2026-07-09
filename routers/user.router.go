package routers

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	UserController *controllers.UserController
}

// constructor
// its not following the dependency injection behaviour with the controllers layer
// router is a concrete class directly depending on the controllers
// we already know that controllers dont have business logic -- they only have 1 responsibility
// taking the request passing to the service layer -- and reponse passing back
// they are just passer
// if in future they'll have more logic then u can intoduce dependency injection
func NewUserRouter(_usercontroller *controllers.UserController) Router { // this return type is in router.go

	// just by making the return type Router and returning the router
	// doesnt make UserRouter a router -- it needs to implement all the functions of the router to become that
	return &UserRouter{
		UserController: _usercontroller,
	}
}

// member function -- registering to the main chi router
func (ur *UserRouter) Register(chiR chi.Router) {

	// basically registering the routes to the one main chi ROUTERRR -- router.go file
	chiR.Post("/signup", ur.UserController.RegisterUser)

}
