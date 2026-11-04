package factories

import (
	"dnd-character-dashboard/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
)

func NewCharacter(db *gorm.DB, character *models.Character) {
	fillCharacterDefaults(character)
	err := db.Create(character).Error
	if err != nil {
		panic("error creating character in factory: " + err.Error())
	}
}

func fillCharacterDefaults(character *models.Character) {
	if character.UserID == 0 {
		character.UserID = 1
	}
	if character.Name == "" {
		character.Name = random.String(16)
	}
	if character.Level == 0 {
		character.Level = 1
	}
}
