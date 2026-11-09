package factories

import (
	"dnd-character-dashboard/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
)

func NewRaceVariant(db *gorm.DB, raceVariant *models.RaceVariant) {
	fillRaceVariantDefaults(raceVariant)
	err := db.Create(raceVariant).Error
	if err != nil {
		panic("error creating race variant in factory: " + err.Error())
	}
}

func fillRaceVariantDefaults(raceVariant *models.RaceVariant) {
	if raceVariant.Name == "" {
		raceVariant.Name = random.String(16)
	}
	if raceVariant.Description == "" {
		raceVariant.Description = random.String(64)
	}
}
