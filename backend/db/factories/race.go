package factories

import (
	"dnd-character-dashboard/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"math/rand"
)

func NewRace(db *gorm.DB, race *models.Race) {
	fillRaceDefaults(race)
	err := db.Create(race).Error
	if err != nil {
		panic("error creating race in factory: " + err.Error())
	}
}

func fillRaceDefaults(race *models.Race) {
	if race.Name == "" {
		race.Name = random.String(16)
	}
	if race.CardDescription == "" {
		race.CardDescription = random.String(64)
	}
	if race.Description == "" {
		race.Description = random.String(64)
	}
	if race.CreatureType == "" {
		race.CreatureType = random.String(16)
	}
	if race.Size == "" {
		race.Size = random.String(16)
	}
	if race.BaseSpeed == 0 {
		race.BaseSpeed = rand.Intn(100)
	}
}
