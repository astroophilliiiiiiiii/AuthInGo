package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

// tells ki isme kya kya methods honge -- kya kya krregi repo
type PermissionRepository interface {
	GetPermissionById(id int) (*models.Permission, error)
	GetPermissionByName(name string) (*models.Permission, error)
	GetAllPermissions() ([]*models.Permission, error)
	CreatePermission(name string, description string, resource string, action string) (*models.Permission, error)
	DeletePermissionByID(id int) error
	UpdatePermission(id int, name string, description string, resource string, action string) (*models.Permission, error)
}

// what it will get access to   -- repo jo jo krregi vo kaise krregi kiski help se
type PermissionRepositoryImpl struct {
	db *sql.DB // Mere paas database ka connection hai.
}

// constructor   -- give me db connection and ill give u PermissionRepository ( object )
func NewPermissionRepository(_db *sql.DB) PermissionRepository {
	return &PermissionRepositoryImpl{
		db: _db,
	}
}

func (r *PermissionRepositoryImpl) GetPermissionById(id int) (*models.Permission, error) {
	query := "SELECT id , name , description , resource , action ,  created_at , updated_at FROM Permissions WHERE id=?"
	row := r.db.QueryRow(query, id)

	Permission := &models.Permission{}

	err := row.Scan(&Permission.Id, &Permission.Name, &Permission.Description, &Permission.Resource, &Permission.Action, &Permission.Created_at, &Permission.Updated_at)
	if err != nil {
		return nil, err
	}

	return Permission, nil
}

func (r *PermissionRepositoryImpl) GetPermissionByName(name string) (*models.Permission, error) {
	query := "SELECT id , name , description , resource , action , created_at , updated_at FROM Permissions WHERE name=?"
	row := r.db.QueryRow(query, name)

	Permission := &models.Permission{}

	err := row.Scan(&Permission.Id, &Permission.Name, &Permission.Description, &Permission.Resource, &Permission.Action, &Permission.Created_at, &Permission.Updated_at)
	if err != nil {
		return nil, err
	}

	return Permission, nil
}

func (r *PermissionRepositoryImpl) GetAllPermissions() ([]*models.Permission, error) {
	query := "SELECT id , name , description , resource , action , created_at , updated_at FROM Permissions"
	rows, err := r.db.Query(query)
	if err != nil {
		fmt.Println("Error fetching the rows!!")
		return nil, err
	}

	Permissions := []*models.Permission{} // as there is array of users
	defer rows.Close()                    // Taaki database resources/free connection release ho jaye aur memory leak na ho. 🚀

	for rows.Next() { // rows ke andar ek internal pointer/cursor hota hai ✅
		// next krne se pointer mover to the next
		Permission := &models.Permission{}

		err := rows.Scan(
			&Permission.Id,
			&Permission.Name,
			&Permission.Description,
			&Permission.Resource, &Permission.Action,
			&Permission.Created_at,
			&Permission.Updated_at,
		)

		if err != nil {
			return nil, err // stops the flow -- nil returned !!
		}

		Permissions = append(Permissions, Permission)
	}

	fmt.Println("All users fetched successfullyy !! ")

	return Permissions, nil
}

func (r *PermissionRepositoryImpl) CreatePermission(name string, description string, resource string, action string) (*models.Permission, error) {
	query := "INSERT INTO Permissions ( name , description , resource , action ) VALUES ( ? , ? , ? , ? )"
	result, err := r.db.Exec(query, name, description, resource, action)
	if err != nil {
		fmt.Println("Error inserting the Permission!!")
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected:", err)
		return nil, err
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, Permission not created!")
		return nil, fmt.Errorf("Permission not created")
	}

	fmt.Println("Permission created successfully, rows affected:", rowsAffected)

	// Fetch the created Permission
	Permission, err := r.GetPermissionByName(name)
	if err != nil {
		return nil, err
	}

	return Permission, nil
}

func (r *PermissionRepositoryImpl) DeletePermissionByID(id int) error {
	query := "DELETE FROM Permissions WHERE id=?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Println("Error deleting the Permission!!")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, Permission not deleted!")
		return fmt.Errorf("Permission not deleted")
	}

	fmt.Println("Permission deleted successfully, rows affected:", rowsAffected)

	return nil
}

func (r *PermissionRepositoryImpl) UpdatePermission(id int, name string, description string, resource string, action string) (*models.Permission, error) {
	query := "UPDATE Permissions SET name=?, description=? , resource=? , action=? WHERE id=?"
	result, err := r.db.Exec(query, name, description, resource, action, id)
	if err != nil {
		fmt.Println("Error updating the Permission!!")
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected:", err)
		return nil, err
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, Permission not updated!")
		return nil, fmt.Errorf("Permission not updated")
	}

	fmt.Println("Permission updated successfully, rows affected:", rowsAffected)

	Permission, err := r.GetPermissionById(id)
	if err != nil {
		return nil, err
	}

	return Permission, nil
}
