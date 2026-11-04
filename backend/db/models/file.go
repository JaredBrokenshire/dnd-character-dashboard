package models

import "time"

type File struct {
	ID       uint   `json:"-" gorm:"primary_key"`
	Name     string `json:"name"`
	Location string `json:"location"`

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
