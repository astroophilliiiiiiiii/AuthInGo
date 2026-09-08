package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

// tells ki isme kya kya methods honge -- kya kya krregi repo
type RoleRepository interface {
	GetRoleById(id int) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)
	CreateRole(name string, description string) (*models.Role, error)
	DeleteRoleByID(id int) error
	UpdateRole(id int, name string, description string) (*models.Role, error)
}

// what it will get access to   -- repo jo jo krregi vo kaise krregi kiski help se
type RoleRepositoryImpl struct {
	db *sql.DB // Mere paas database ka connection hai.
}

// constructor   -- give me db connection and ill give u RoleRepository ( object )
func NewRoleRepository(_db *sql.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db: _db,
	}
}

func (r *RoleRepositoryImpl) GetRoleById(id int) (*models.Role, error) {
	query := "SELECT id , name , description , created_at , updated_at FROM roles WHERE id=?"
	row := r.db.QueryRow(query, id)

	role := &models.Role{}

	err := row.Scan(&role.Id, &role.Name, &role.Description, &role.Created_at, &role.Updated_at)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r *RoleRepositoryImpl) GetRoleByName(name string) (*models.Role, error) {
	query := "SELECT id , name , description , created_at , updated_at FROM roles WHERE name=?"
	row := r.db.QueryRow(query, name)

	role := &models.Role{}

	err := row.Scan(&role.Id, &role.Name, &role.Description, &role.Created_at, &role.Updated_at)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r *RoleRepositoryImpl) GetAllRoles() ([]*models.Role, error) {
	query := "SELECT id , name , description , created_at , updated_at FROM roles"
	rows, err := r.db.Query(query)
	if err != nil {
		fmt.Println("Error fetching the rows!!")
		return nil, err
	}

	roles := []*models.Role{} // as there is array of users
	defer rows.Close()        // Taaki database resources/free connection release ho jaye aur memory leak na ho. 🚀

	for rows.Next() { // rows ke andar ek internal pointer/cursor hota hai ✅
		// next krne se pointer mover to the next
		role := &models.Role{}

		err := rows.Scan(
			&role.Id,
			&role.Name,
			&role.Description,
			&role.Created_at,
			&role.Updated_at,
		)

		if err != nil {
			return nil, err // stops the flow -- nil returned !!
		}

		roles = append(roles, role)
	}

	fmt.Println("All users fetched successfullyy !! ")

	return roles, nil
}

func (r *RoleRepositoryImpl) CreateRole(name string, description string) (*models.Role, error) {
	query := "INSERT INTO roles ( name , description ) VALUES ( ? , ? )"
	result, err := r.db.Exec(query, name, description)
	if err != nil {
		fmt.Println("Error inserting the role!!")
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected:", err)
		return nil, err
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, role not created!")
		return nil, fmt.Errorf("role not created")
	}

	fmt.Println("Role created successfully, rows affected:", rowsAffected)

	// Fetch the created role
	role, err := r.GetRoleByName(name)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r *RoleRepositoryImpl) DeleteRoleByID(id int) error {
	query := "DELETE FROM roles WHERE id=?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Println("Error deleting the role!!")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, role not deleted!")
		return fmt.Errorf("role not deleted")
	}

	fmt.Println("Role deleted successfully, rows affected:", rowsAffected)

	return nil
}

func (r *RoleRepositoryImpl) UpdateRole(id int, name string, description string) (*models.Role, error) {
	query := "UPDATE roles SET name=?, description=? WHERE id=?"
	result, err := r.db.Exec(query, name, description, id)
	if err != nil {
		fmt.Println("Error updating the role!!")
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected:", err)
		return nil, err
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, role not updated!")
		return nil, fmt.Errorf("role not updated")
	}

	fmt.Println("Role updated successfully, rows affected:", rowsAffected)

	role, err := r.GetRoleById(id)
	if err != nil {
		return nil, err
	}

	return role, nil
}
