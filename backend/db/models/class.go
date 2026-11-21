package models

type Class struct {
	ID   uint `gorm:"primary_key;auto_increment"`
	Name string

	Tagline          string
	CardDescription  string
	PrimaryAbility   string
	HitPointDieValue int

	Description string

	SaveProficiencies          []*ClassSaveProficiency
	SkillProficiencies         []*ClassSkillProficiency
	SkillProficiencyMaxOptions uint
	WeaponProficiencies        []*ClassWeaponProficiency
	ArmourProficiencies        []*ClassArmourProficiency

	Subclasses []*Subclass

	// Images
	LogoFileID uint
	Logo       *File `gorm:"foreignkey:LogoFileID"`

	CardBackgroundFileID uint
	CardBackground       *File `gorm:"foreignkey:CardBackgroundFileID"`

	DetailsImageID uint
	DetailsImage   *File `gorm:"foreignkey:DetailsImageID"`
}
