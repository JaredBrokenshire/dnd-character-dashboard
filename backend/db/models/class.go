package models

type Class struct {
	ID   uint
	Name string

	LogoFileID uint
	Logo       *File `gorm:"foreignkey:LogoFileID"`

	CardBackgroundFileID uint
	CardBackground       *File `gorm:"foreignkey:CardBackgroundFileID"`

	Subclasses []Subclass
}
