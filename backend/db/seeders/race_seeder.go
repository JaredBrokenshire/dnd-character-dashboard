package seeders

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/pkg/logging"
	"fmt"
)

func (s *Seeder) SetRaces() {
	races := []models.Race{
		{
			ID:   1,
			Name: "Aasimar",
		},
		{
			ID:   2,
			Name: "Dragonborn",
		},
		{
			ID:   3,
			Name: "Dwarf",
		},
		{
			ID:   4,
			Name: "Elf",
		},
		{
			ID:   5,
			Name: "Gnome",
		},
		{
			ID:   6,
			Name: "Goliath",
		},
		{
			ID:   7,
			Name: "Halfling",
		},
		{
			ID:   8,
			Name: "Human",
		},
		{
			ID:   9,
			Name: "Orc",
		},
		{
			ID:   10,
			Name: "Tiefling",
		},
		{
			ID:   11,
			Name: "Half-Elf",
		},
		{
			ID:   12,
			Name: "Half-Orc",
		},
		{
			ID:   13,
			Name: "Air Genasi",
		},
		{
			ID:   14,
			Name: "Earth Genasi",
		},
		{
			ID:   15,
			Name: "Fire Genasi",
		},
		{
			ID:   16,
			Name: "Water Genasi",
		},
		{
			ID:   17,
			Name: "Warforged",
		},
		{
			ID:   18,
			Name: "Kobold",
		},
		{
			ID:   19,
			Name: "Dhampir",
		},
	}

	for _, race := range races {
		err := s.DB.Where("id = ?", race.ID).FirstOrCreate(&race).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating race with id `%v` in seeder: %v", race.ID, s.DB.Error))
		}

	}
}
