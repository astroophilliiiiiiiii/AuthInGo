package routers

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

type RoleRouter struct {
	RoleController *controllers.RoleController
}

// constructor
func NewRoleRouter(_rolecontroller *controllers.RoleController) Router { // this return type is in router.go
	return &RoleRouter{
		RoleController: _rolecontroller,
	}
}

// member function -- registering to the main chi router
func (rr *RoleRouter) Register(r chi.Router) {
	r.Get("/roles/{id}", rr.RoleController.GetRoleById)
	r.Get("/roles", rr.RoleController.GetAllRoles)
	//	r.Post("/roles", rr.RoleController.CreateRole )
	//r.With(middlewares.UpdateRoleRequestValidator).Put("/roles/{id}", rr.RoleController.UpdateRole)
	r.Delete("/roles/{id}", rr.RoleController.DeleteRole)

	// Role permissions operations
	r.Get("/roles/{id}/permissions", rr.RoleController.GetRolePermissions)
	//r.With(middlewares.AssignPermissionRequestValidator).Post("/roles/{id}/permissions", rr.RoleController.AssignPermissionToRole)
	//r.With(middlewares.RemovePermissionRequestValidator).Delete("/roles/{id}/permissions", rr.RoleController.RemovePermissionFromRole)
	r.Get("/role-permissions", rr.RoleController.GetAllRolePermissions)
	//r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAllRoles("admin")).Post("/roles/{userId}/assign/{roleId}", rr.RoleController.AssignRoleToUser)
}
