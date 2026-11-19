package models

type User struct {
	ID       uint
	Username string
	Password string

	Roles []*Role `gorm:"many2many:user_roles"`
}
