package models

// models created -- to be used in repo layer as a template !!
type Role struct {
	Id          int64
	Name        string
	Description string
	Created_at  string
	Updated_at  string
}

type Permission struct {
	Id          int
	Name        string
	Description string
	Resource    string
	Action      string
	Created_at  string
	Updated_at  string
}

type RolePermission struct {
	Id           int64
	RoleId       int64
	PermissionID int
	Created_at   string
	Updated_at   string
}
