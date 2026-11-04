package models

type Character struct {
	ID     uint
	UserID uint
	User   User

	Name  string
	Level int

	RaceID uint
	Race   *Race

	Classes    []*Class    `gorm:"many2many:character_classes;"`
	Subclasses []*Subclass `gorm:"many2many:character_subclasses;"`

	CampaignID uint
	Campaign   *Campaign

	ProfilePictureFileID uint
	ProfilePicture       *File `gorm:"foreignkey:ProfilePictureFileID"`
}
