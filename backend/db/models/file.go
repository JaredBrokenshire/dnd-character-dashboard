package models

import "time"

type File struct {
	ID       uint `gorm:"primary_key;auto_increment"`
	Name     string
	Location string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
