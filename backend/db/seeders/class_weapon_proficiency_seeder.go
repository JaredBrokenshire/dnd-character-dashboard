package seeders

import (
	"dnd/db/models"
	"dnd/db/models/types"
	"dnd/pkg/logging"
	"fmt"
)

func (s *Seeder) SetClassWeaponProficiencies() {
	classWeaponProficiencies := []models.ClassWeaponProficiency{
		{
			ID:         1,
			ClassID:    1, // Artificer
			WeaponType: types.WeaponTypeSimple,
		},
		{
			ID:         2,
			ClassID:    2, // Barbarian
			WeaponType: types.WeaponTypeSimple,
		},
		{
			ID:         3,
			ClassID:    2, // Barbarian
			WeaponType: types.WeaponTypeMartial,
		},
		{
			ID:         4,
			ClassID:    3, // Bard
			WeaponType: types.WeaponTypeSimple,
		},
		{
			ID:         5,
			ClassID:    4, // Blood Hunter
			WeaponType: types.WeaponTypeSimple,
		},
		{
			ID:         6,
			ClassID:    4, // Blood Hunter
			WeaponType: types.WeaponTypeMartial,
		},
		{
			ID:         7,
			ClassID:    5, // Cleric
			WeaponType: types.WeaponTypeSimple,
		},
		{
			ID:         8,
			ClassID:    6, // Druid
			WeaponType: types.WeaponTypeSimple,
		},
		{
			ID:         9,
			ClassID:    7, // Fighter
			WeaponType: types.WeaponTypeSimple,
		},
		{
			ID:         10,
			ClassID:    7, // Fighter
			WeaponType: types.WeaponTypeMartial,
		},
		{
			ID:         11,
			ClassID:    8, // Monk
			WeaponType: types.WeaponTypeSimple,
		},
		{
			ID:         12,
			ClassID:    8, // Monk
			WeaponType: types.WeaponTypeMartial,
			Condition:  "Light",
		},
		{
			ID:         13,
			ClassID:    9, // Paladin
			WeaponType: types.WeaponTypeSimple,
		},
		{
			ID:         14,
			ClassID:    9, // Paladin
			WeaponType: types.WeaponTypeMartial,
		},
		{
			ID:         15,
			ClassID:    10, // Ranger
			WeaponType: types.WeaponTypeSimple,
		},
		{
			ID:         16,
			ClassID:    10, // Ranger
			WeaponType: types.WeaponTypeMartial,
		},
		{
			ID:         17,
			ClassID:    11, // Rogue
			WeaponType: types.WeaponTypeSimple,
		},
		{
			ID:         18,
			ClassID:    11, // Rogue
			WeaponType: types.WeaponTypeMartial,
			Condition:  "Finesse or Light",
		},
		{
			ID:         19,
			ClassID:    12, // Sorcerer
			WeaponType: types.WeaponTypeSimple,
		},
		{
			ID:         20,
			ClassID:    13, // Warlock
			WeaponType: types.WeaponTypeSimple,
		},
		{
			ID:         21,
			ClassID:    14, // Wizard
			WeaponType: types.WeaponTypeSimple,
		},
	}

	for _, cwp := range classWeaponProficiencies {
		err := s.DB.Where("id = ?", cwp.ID).FirstOrCreate(&cwp).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating class weapon proficiency with id `%v` in seeder: %v", cwp.ID, err.Error()))
		}
	}
}
