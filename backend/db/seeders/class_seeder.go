package seeders

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/pkg/logging"
	"fmt"
)

func (s *Seeder) SetClasses() {
	classes := []models.Class{
		{
			ID:   1,
			Name: "Artificer",
		},
		{
			ID:   2,
			Name: "Barbarian",
		},
		{
			ID:   3,
			Name: "Bard",
		},
		{
			ID:   4,
			Name: "Blood Hunter",
		},
		{
			ID:   5,
			Name: "Cleric",
		},
		{
			ID:   6,
			Name: "Druid",
		},
		{
			ID:   7,
			Name: "Fighter",
		},
		{
			ID:   8,
			Name: "Monk",
		},
		{
			ID:   9,
			Name: "Paladin",
		},
		{
			ID:   10,
			Name: "Ranger",
		},
		{
			ID:   11,
			Name: "Rogue",
		},
		{
			ID:   12,
			Name: "Sorcerer",
		},
		{
			ID:   13,
			Name: "Warlock",
		},
		{
			ID:   14,
			Name: "Wizard",
		},
	}

	for _, class := range classes {
		err := s.DB.Where("id = ?", class.ID).FirstOrCreate(&class).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating class with id `%v` in seeder: %v", class.ID, err.Error()))
		}
	}
}
