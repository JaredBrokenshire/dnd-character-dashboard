package models

import "time"

type File struct {
	ID       uint
	Name     string
	Location string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
