package models

type Class struct {
	ID   uint `gorm:"primary_key;auto_increment"`
	Name string

	Tagline          string
	CardDescription  string
	PrimaryAbility   string
	HitPointDieValue int

	SaveProficiencies          []*ClassSaveProficiency
	WeaponProficiencies        []*ClassWeaponProficiency

	LogoFileID uint
	Logo       *File `gorm:"foreignkey:LogoFileID"`

	CardBackgroundFileID uint
	CardBackground       *File `gorm:"foreignkey:CardBackgroundFileID"`
}
