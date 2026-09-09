package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
	"strings"
)

type UserRoleRepository interface {
	GetUserRoles(userId int64) ([]*models.Role, error)               // all roles of the user
	AssignRoleToUser(userId int64, roleId int64) error               // assigning role to a user
	RemoveRoleFromUser(userId int64, roleId int64) error             // removing role from user
	GetUserPermissions(userId int64) ([]*models.Permission, error)   // permissions of a user
	HasPermission(userId int64, permissionName string) (bool, error) // has a specific permission or not
	HasRole(userId int64, roleName string) (bool, error)             // has a specific role or not
	HasAllRoles(userId int64, roleNames []string) (bool, error)
	HasAnyRole(userId int64, roleNames []string) (bool, error)
}

type UserRoleRepositoryImpl struct {
	db *sql.DB
}

func NewUserRoleRepository(_db *sql.DB) UserRoleRepository {
	return &UserRoleRepositoryImpl{
		db: _db,
	}
}

func (u *UserRoleRepositoryImpl) GetUserRoles(userId int64) ([]*models.Role, error) {
	query := `
        SELECT r.id, r.name, r.description, r.created_at, r.updated_at
        FROM user_roles ur
        INNER JOIN roles r ON ur.role_id = r.id
        WHERE ur.user_id = ?`
	rows, err := u.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*models.Role
	for rows.Next() {
		role := &models.Role{}
		if err := rows.Scan(&role.Id, &role.Name, &role.Description, &role.Created_at, &role.Updated_at); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (u *UserRoleRepositoryImpl) AssignRoleToUser(userId int64, roleId int64) error {
	query := "INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)"
	_, err := u.db.Exec(query, userId, roleId)
	return err
}

func (u *UserRoleRepositoryImpl) RemoveRoleFromUser(userId int64, roleId int64) error {
	query := "DELETE FROM user_roles WHERE user_id = ? AND role_id = ?"
	_, err := u.db.Exec(query, userId, roleId)
	return err
}

func (u *UserRoleRepositoryImpl) GetUserPermissions(userId int64) ([]*models.Permission, error) {
	query := `
        SELECT p.id, p.name, p.description, p.resource, p.action, p.created_at, p.updated_at
        FROM user_roles ur
        INNER JOIN role_permissions rp ON ur.role_id = rp.role_id
        INNER JOIN permissions p ON rp.permission_id = p.id
        WHERE ur.user_id = ?`
	rows, err := u.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permission
	for rows.Next() {
		permission := &models.Permission{}
		if err := rows.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.Created_at, &permission.Updated_at); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (u *UserRoleRepositoryImpl) HasPermission(userId int64, permissionName string) (bool, error) {
	query := `
        SELECT COUNT(*) > 0
        FROM user_roles ur
        INNER JOIN role_permissions rp ON ur.role_id = rp.role_id
        INNER JOIN permissions p ON rp.permission_id = p.id
        WHERE ur.user_id = ? AND p.name = ?`
	var exists bool
	err := u.db.QueryRow(query, userId, permissionName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *UserRoleRepositoryImpl) HasRole(userId int64, roleName string) (bool, error) {
	query := `
        SELECT COUNT(*) > 0
        FROM user_roles ur
        INNER JOIN roles r ON ur.role_id = r.id
        WHERE ur.user_id = ? AND r.name = ?`
	var exists bool
	err := u.db.QueryRow(query, userId, roleName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *UserRoleRepositoryImpl) HasAllRoles(userId int64, roleNames []string) (bool, error) {
	if len(roleNames) == 0 {
		return true, nil // If no roles are specified, return true
	}

	// Create dynamic placeholders for the IN clause (e.g., "?,?,?")
	placeholders := strings.Repeat("?,", len(roleNames))
	placeholders = strings.TrimSuffix(placeholders, ",")

	query := fmt.Sprintf(`
        SELECT COUNT(*)
        FROM user_roles ur
        INNER JOIN roles r ON ur.role_id = r.id
        WHERE ur.user_id = ? AND r.name IN (%s)`, placeholders)

	// Create args slice with userId first, then all roleNames
	args := make([]interface{}, 0, 1+len(roleNames))
	args = append(args, userId)
	for _, roleName := range roleNames {
		args = append(args, roleName)
	}

	var count int
	err := u.db.QueryRow(query, args...).Scan(&count)
	if err != nil {
		return false, err
	}

	// If the count of matched roles equals the number of required roles, they have all of them.
	return count == len(roleNames), nil
}

func (u *UserRoleRepositoryImpl) HasAnyRole(userId int64, roleNames []string) (bool, error) {
	if len(roleNames) == 0 {
		return true, nil // If no roles are specified, return true
	}

	placeholders := strings.Repeat("?,", len(roleNames))
	placeholders = strings.TrimSuffix(placeholders, ",")

	query := fmt.Sprintf(`
		SELECT COUNT(*) > 0 
		FROM user_roles ur 
		INNER JOIN roles r ON ur.role_id = r.id 
		WHERE ur.user_id = ? AND r.name IN (%s)`, placeholders)

	args := make([]interface{}, 0, 1+len(roleNames))
	args = append(args, userId)
	for _, roleName := range roleNames {
		args = append(args, roleName)
	}

	var hasAnyRole bool
	err := u.db.QueryRow(query, args...).Scan(&hasAnyRole)
	if err != nil {
		return false, err
	}

	return hasAnyRole, nil
}
