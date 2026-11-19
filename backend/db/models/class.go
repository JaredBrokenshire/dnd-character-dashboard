package models

type Ability string

const (
	Strength     Ability = "Strength"
	Dexterity    Ability = "Dexterity"
	Constitution Ability = "Constitution"
	Intelligence Ability = "Intelligence"
	Wisdom       Ability = "Wisdom"
	Charisma     Ability = "Charisma"
)

type Class struct {
	ID   uint `gorm:"primary_key;auto_increment"`
	Name string

	Tagline          string
	CardDescription  string
	PrimaryAbility   string
	HitPointDieValue int
	Save             Ability
	SecondarySave    Ability

	LogoFileID uint
	Logo       *File `gorm:"foreignkey:LogoFileID"`

	CardBackgroundFileID uint
	CardBackground       *File `gorm:"foreignkey:CardBackgroundFileID"`
}
