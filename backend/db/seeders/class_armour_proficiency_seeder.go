package seeders

import (
	"dnd/db/models"
	"dnd/db/models/types"
	"dnd/pkg/logging"
	"fmt"
)

func (s *Seeder) SetClassArmourProficiencies() {
	classArmourProficiencies := []models.ClassArmourProficiency{
		{
			ID:      1,
			ClassID: 1, // Artificer
			Armour:  types.ArmourTypeLight,
		},
		{
			ID:      2,
			ClassID: 1, // Artificer
			Armour:  types.ArmourTypeMedium,
		},
		{
			ID:      3,
			ClassID: 1, // Artificer
			Armour:  types.ArmourTypeShields,
		},
		{
			ID:      4,
			ClassID: 2, // Barbarian
			Armour:  types.ArmourTypeLight,
		},
		{
			ID:      5,
			ClassID: 2, // Barbarian
			Armour:  types.ArmourTypeMedium,
		},
		{
			ID:      6,
			ClassID: 2, // Barbarian
			Armour:  types.ArmourTypeShields,
		},
		{
			ID:      7,
			ClassID: 3, // Bard
			Armour:  types.ArmourTypeLight,
		},
		{
			ID:      8,
			ClassID: 4, // Blood Hunter
			Armour:  types.ArmourTypeLight,
		},
		{
			ID:      9,
			ClassID: 4, // Blood Hunter
			Armour:  types.ArmourTypeMedium,
		},
		{
			ID:      10,
			ClassID: 4, // Blood Hunter
			Armour:  types.ArmourTypeShields,
		},
		{
			ID:      11,
			ClassID: 5, // Cleric
			Armour:  types.ArmourTypeLight,
		},
		{
			ID:      12,
			ClassID: 5, // Cleric
			Armour:  types.ArmourTypeMedium,
		},
		{
			ID:      13,
			ClassID: 5, // Cleric
			Armour:  types.ArmourTypeShields,
		},
		{
			ID:      14,
			ClassID: 6, // Druid
			Armour:  types.ArmourTypeLight,
		},
		{
			ID:      15,
			ClassID: 6, // Druid
			Armour:  types.ArmourTypeShields,
		},
		{
			ID:      16,
			ClassID: 7, // Fighter
			Armour:  types.ArmourTypeLight,
		},
		{
			ID:      17,
			ClassID: 7, // Fighter
			Armour:  types.ArmourTypeMedium,
		},
		{
			ID:      18,
			ClassID: 7, // Fighter
			Armour:  types.ArmourTypeHeavy,
		},
		{
			ID:      19,
			ClassID: 7, // Fighter
			Armour:  types.ArmourTypeShields,
		},
		{
			ID:      20,
			ClassID: 9, // Paladin
			Armour:  types.ArmourTypeLight,
		},
		{
			ID:      21,
			ClassID: 9, // Paladin
			Armour:  types.ArmourTypeMedium,
		},
		{
			ID:      22,
			ClassID: 9, // Paladin
			Armour:  types.ArmourTypeHeavy,
		},
		{
			ID:      23,
			ClassID: 9, // Paladin
			Armour:  types.ArmourTypeShields,
		},
		{
			ID:      24,
			ClassID: 10, // Ranger
			Armour:  types.ArmourTypeLight,
		},
		{
			ID:      25,
			ClassID: 10, // Ranger
			Armour:  types.ArmourTypeMedium,
		},
		{
			ID:      26,
			ClassID: 10, // Ranger
			Armour:  types.ArmourTypeShields,
		},
		{
			ID:      27,
			ClassID: 11, // Rogue
			Armour:  types.ArmourTypeLight,
		},
		{
			ID:      28,
			ClassID: 13, // Warlock
			Armour:  types.ArmourTypeLight,
		},
	}

	for _, classArmourProficiency := range classArmourProficiencies {
		err := s.DB.Where("id = ?", classArmourProficiency.ID).FirstOrCreate(&classArmourProficiency).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating class armour proficiency with id `%v` in seeder: %v", classArmourProficiency.ID, s.DB.Error))
		}
	}
}
