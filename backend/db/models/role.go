package models

// RoleName is the string identifier for a user role.
type RoleName string

const (
	RoleAdmin RoleName = "admin"
	RoleUser  RoleName = "user"
)

// Role represents a user role in the system.
type Role struct {
	ID   uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name RoleName `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`

	Users []User `gorm:"many2many:user_roles;" json:"-"`
}
