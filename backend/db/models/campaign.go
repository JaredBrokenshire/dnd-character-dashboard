package models

import "time"

type Campaign struct {
	ID     uint
	UserID uint
	User   User

	Name string

	Characters []*Character

	CreatedAt time.Time
	UpdatedAt time.Time
}
