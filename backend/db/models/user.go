package models

import (
	"time"
)

// User represents an application user.
type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"username"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"` // bcrypt hash — never serialised
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Roles []Role `gorm:"many2many:user_roles;" json:"roles,omitempty"`
}

// HasRole returns true if the user has the given role.
func (u *User) HasRole(name RoleName) bool {
	for _, r := range u.Roles {
		if r.Name == name {
			return true
		}
	}
	return false
}
