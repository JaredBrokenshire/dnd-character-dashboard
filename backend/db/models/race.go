package models

type Race struct {
	ID   uint
	Name string

	CardDescription string
	Description     string

	CreatureType string
	Size         string
	BaseSpeed    int

	Variants []*RaceVariant
	Traits   []*Trait `gorm:"many2many:race_traits;"`

	CardBackgroundFileID uint
	CardBackground       *File `gorm:"foreignkey:CardBackgroundFileID"`

	DetailsImageFileID uint
	DetailsImage       *File `gorm:"foreignkey:DetailsImageFileID"`
}
