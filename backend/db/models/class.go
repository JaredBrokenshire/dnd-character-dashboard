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
	ID   uint
	Name string

	LogoFileID uint
	Logo       *File `gorm:"foreignkey:LogoFileID"`

	CardBackgroundFileID uint
	CardBackground       *File `gorm:"foreignkey:CardBackgroundFileID"`

	Tagline          string
	CardDescription  string
	PrimaryAbility   string
	HitPointDieValue int
	Save             Ability
	SecondarySave    Ability

	Subclasses []Subclass
}
