package models

type Class struct {
	ID   uint
	Name string

	LogoFileID uint
	Logo       *File `gorm:"foreignkey:LogoFileID"`

	Subclasses []Subclass
}
