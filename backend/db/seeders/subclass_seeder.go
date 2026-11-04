package seeders

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/pkg/logging"
	"fmt"
)

func (s *Seeder) SetSubclasses() {
	subclasses := []models.Subclass{
		{
			ID:      1,
			Name:    "Order of the Mutant",
			ClassID: 4,
		},
		{
			ID:      2,
			Name:    "Echo Knight",
			ClassID: 7,
		},
		{
			ID:      3,
			Name:    "Order of the Lycan",
			ClassID: 4,
		},
		{
			ID:      4,
			Name:    "Path of the Totem Warrior",
			ClassID: 2,
		},
		{
			ID:      5,
			Name:    "College of Glamour",
			ClassID: 3,
		},
		{
			ID:      6,
			Name:    "Oath of the Ancients",
			ClassID: 9,
		},
		{
			ID:      7,
			Name:    "Battle Smith",
			ClassID: 1,
		},
		{
			ID:      8,
			Name:    "Shadow Magic",
			ClassID: 12,
		},
		{
			ID:      9,
			Name:    "Twilight Domain",
			ClassID: 5,
		},
		{
			ID:      10,
			Name:    "Way of the Drunken Master",
			ClassID: 8,
		},
		{
			ID:      11,
			Name:    "Gloom Stalker",
			ClassID: 10,
		},
		{
			ID:      12,
			Name:    "Path of the Berserker",
			ClassID: 2,
		},
	}

	for _, subclass := range subclasses {
		err := s.DB.Where("id = ?", subclass.ID).FirstOrCreate(&subclass).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating subclass with id `%v` in seeder: %v", subclass.ID, s.DB.Error))
		}
	}
}
