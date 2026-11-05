package models

type Race struct {
	ID   uint
	Name string

	CardDescription string

	Traits []*Trait `gorm:"many2many:race_traits;"`

	CardBackgroundFileID uint
	CardBackground       *File `gorm:"foreignkey:CardBackgroundFileID"`
}
