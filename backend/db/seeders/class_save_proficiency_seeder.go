package seeders

import (
	"dnd/db/models"
	"dnd/db/models/types"
	"dnd/pkg/logging"
	"fmt"
)

func (s *Seeder) SetClassSaveProficiencies() {
	classSaveProficiencies := []models.ClassSaveProficiency{
		{
			ID:      1,
			ClassID: 1, // Artificer
			Ability: types.AbilityConstitution,
		},
		{
			ID:      2,
			ClassID: 1, // Artificer
			Ability: types.AbilityIntelligence,
		},
		{
			ID:      3,
			ClassID: 2, // Barbarian
			Ability: types.AbilityStrength,
		},
		{
			ID:      4,
			ClassID: 2, // Barbarian
			Ability: types.AbilityConstitution,
		},
		{
			ID:      5,
			ClassID: 3, // Bard
			Ability: types.AbilityDexterity,
		},
		{
			ID:      6,
			ClassID: 3, // Bard
			Ability: types.AbilityCharisma,
		},
		{
			ID:      7,
			ClassID: 4, // Blood Hunter
			Ability: types.AbilityDexterity,
		},
		{
			ID:      8,
			ClassID: 4, // Blood Hunter
			Ability: types.AbilityIntelligence,
		},
		{
			ID:      9,
			ClassID: 5, // Cleric
			Ability: types.AbilityWisdom,
		},
		{
			ID:      10,
			ClassID: 5, // Cleric
			Ability: types.AbilityCharisma,
		},
		{
			ID:      11,
			ClassID: 6, // Druid
			Ability: types.AbilityIntelligence,
		},
		{
			ID:      12,
			ClassID: 6, // Druid
			Ability: types.AbilityWisdom,
		},
		{
			ID:      13,
			ClassID: 7, // Fighter
			Ability: types.AbilityStrength,
		},
		{
			ID:      14,
			ClassID: 7, // Fighter
			Ability: types.AbilityConstitution,
		},
		{
			ID:      15,
			ClassID: 8, // Monk
			Ability: types.AbilityStrength,
		},
		{
			ID:      16,
			ClassID: 8, // Monk
			Ability: types.AbilityDexterity,
		},
		{
			ID:      17,
			ClassID: 9, // Paladin
			Ability: types.AbilityWisdom,
		},
		{
			ID:      18,
			ClassID: 9, // Paladin
			Ability: types.AbilityCharisma,
		},
		{
			ID:      19,
			ClassID: 10, // Ranger
			Ability: types.AbilityStrength,
		},
		{
			ID:      20,
			ClassID: 10, // Ranger
			Ability: types.AbilityDexterity,
		},
		{
			ID:      21,
			ClassID: 11, // Rogue
			Ability: types.AbilityDexterity,
		},
		{
			ID:      22,
			ClassID: 11, // Rogue
			Ability: types.AbilityIntelligence,
		},
		{
			ID:      23,
			ClassID: 12, // Sorcerer
			Ability: types.AbilityConstitution,
		},
		{
			ID:      24,
			ClassID: 12, // Sorcerer
			Ability: types.AbilityCharisma,
		},
		{
			ID:      25,
			ClassID: 13, // Warlock
			Ability: types.AbilityWisdom,
		},
		{
			ID:      26,
			ClassID: 13, // Warlock
			Ability: types.AbilityCharisma,
		},
		{
			ID:      27,
			ClassID: 14, // Wizard
			Ability: types.AbilityIntelligence,
		},
		{
			ID:      28,
			ClassID: 14, // Wizard
			Ability: types.AbilityWisdom,
		},
	}

	for _, csp := range classSaveProficiencies {
		err := s.DB.Where("id = ?", csp.ID).FirstOrCreate(&csp).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating class save proficiency with id `%v` in seeder: %v", csp.ID, err.Error()))
		}
	}
}
