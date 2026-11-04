package models

type Subclass struct {
	ID   uint
	Name string

	ClassID uint
	Class   Class
}
