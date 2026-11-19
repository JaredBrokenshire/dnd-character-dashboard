package models

type RoleName string

var (
	RoleAdmin RoleName = "admin"
	RoleUser  RoleName = "user"
)

type Role struct {
	ID   uint
	Name RoleName
}
