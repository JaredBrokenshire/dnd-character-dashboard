package factories

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/pkg/logging"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
)

func NewRace(db *gorm.DB, race *models.Race) {
	fillRaceDefaults(race)
	err := db.Create(race).Error
	if err != nil {
		logging.Logger.LogWarn().Msg("Error creating race in factory: " + err.Error())
	}
}

func fillRaceDefaults(race *models.Race) {
	if race.Name == "" {
		race.Name = random.String(16)
	}
}
