package seeders

import (
	"dnd/db/models"
	"dnd/db/models/types"
	"dnd/pkg/logging"
	"fmt"
)

func (s *Seeder) SetClassSkillProficiencies() {
	clasSkillProficiencies := []models.ClassSkillProficiency{
		{
			ID:      1,
			ClassID: 1, // Artificer
			Skill:   types.SkillArcana,
		},
		{
			ID:      2,
			ClassID: 1, // Artificer
			Skill:   types.SkillHistory,
		},
		{
			ID:      3,
			ClassID: 1, // Artificer
			Skill:   types.SkillInvestigation,
		},
		{
			ID:      4,
			ClassID: 1, // Artificer
			Skill:   types.SkillMedicine,
		},
		{
			ID:      5,
			ClassID: 1, // Artificer
			Skill:   types.SkillNature,
		},
		{
			ID:      6,
			ClassID: 1, // Artificer
			Skill:   types.SkillPerception,
		},
		{
			ID:      7,
			ClassID: 1, // Artificer
			Skill:   types.SkillSleightOfHand,
		},
		{
			ID:      8,
			ClassID: 2, // Barbarian
			Skill:   types.SkillAnimalHandling,
		},
		{
			ID:      9,
			ClassID: 2, // Barbarian
			Skill:   types.SkillAthletics,
		},
		{
			ID:      10,
			ClassID: 2, // Barbarian
			Skill:   types.SkillIntimidation,
		},
		{
			ID:      11,
			ClassID: 2, // Barbarian
			Skill:   types.SkillNature,
		},
		{
			ID:      12,
			ClassID: 2, // Barbarian
			Skill:   types.SkillPerception,
		},
		{
			ID:      13,
			ClassID: 2, // Barbarian
			Skill:   types.SkillSurvival,
		},
		{
			ID:      14,
			ClassID: 3, // Bard
			Skill:   types.SkillAcrobatics,
		},
		{
			ID:      15,
			ClassID: 3, // Bard
			Skill:   types.SkillAnimalHandling,
		},
		{
			ID:      16,
			ClassID: 3, // Bard
			Skill:   types.SkillArcana,
		},
		{
			ID:      17,
			ClassID: 3, // Bard
			Skill:   types.SkillAthletics,
		},
		{
			ID:      18,
			ClassID: 3, // Bard
			Skill:   types.SkillDeception,
		},
		{
			ID:      19,
			ClassID: 3, // Bard
			Skill:   types.SkillHistory,
		},
		{
			ID:      20,
			ClassID: 3, // Bard
			Skill:   types.SkillInsight,
		},
		{
			ID:      21,
			ClassID: 3, // Bard
			Skill:   types.SkillIntimidation,
		},
		{
			ID:      22,
			ClassID: 3, // Bard
			Skill:   types.SkillInvestigation,
		},
		{
			ID:      23,
			ClassID: 3, // Bard
			Skill:   types.SkillMedicine,
		},
		{
			ID:      24,
			ClassID: 3, // Bard
			Skill:   types.SkillNature,
		},
		{
			ID:      25,
			ClassID: 3, // Bard
			Skill:   types.SkillPerception,
		},
		{
			ID:      26,
			ClassID: 3, // Bard
			Skill:   types.SkillPerformance,
		},
		{
			ID:      27,
			ClassID: 3, // Bard
			Skill:   types.SkillPersuasion,
		},
		{
			ID:      28,
			ClassID: 3, // Bard
			Skill:   types.SkillReligion,
		},
		{
			ID:      29,
			ClassID: 3, // Bard
			Skill:   types.SkillSleightOfHand,
		},
		{
			ID:      30,
			ClassID: 3, // Bard
			Skill:   types.SkillStealth,
		},
		{
			ID:      31,
			ClassID: 3, // Bard
			Skill:   types.SkillSurvival,
		},
		{
			ID:      32,
			ClassID: 4, // Blood Hunter
			Skill:   types.SkillAcrobatics,
		},
		{
			ID:      33,
			ClassID: 4, // Blood Hunter
			Skill:   types.SkillArcana,
		},
		{
			ID:      34,
			ClassID: 4, // Blood Hunter
			Skill:   types.SkillAthletics,
		},
		{
			ID:      35,
			ClassID: 4, // Blood Hunter
			Skill:   types.SkillHistory,
		},
		{
			ID:      36,
			ClassID: 4, // Blood Hunter
			Skill:   types.SkillInsight,
		},
		{
			ID:      37,
			ClassID: 4, // Blood Hunter
			Skill:   types.SkillInvestigation,
		},
		{
			ID:      38,
			ClassID: 4, // Blood Hunter
			Skill:   types.SkillReligion,
		},
		{
			ID:      39,
			ClassID: 4, // Blood Hunter
			Skill:   types.SkillSurvival,
		},
		{
			ID:      40,
			ClassID: 5, // Cleric
			Skill:   types.SkillHistory,
		},
		{
			ID:      41,
			ClassID: 5, // Cleric
			Skill:   types.SkillInsight,
		},
		{
			ID:      42,
			ClassID: 5, // Cleric
			Skill:   types.SkillMedicine,
		},
		{
			ID:      43,
			ClassID: 5, // Cleric
			Skill:   types.SkillPersuasion,
		},
		{
			ID:      44,
			ClassID: 5, // Cleric
			Skill:   types.SkillReligion,
		},
		{
			ID:      45,
			ClassID: 6, // Druid
			Skill:   types.SkillAnimalHandling,
		},
		{
			ID:      46,
			ClassID: 6, // Druid
			Skill:   types.SkillArcana,
		},
		{
			ID:      47,
			ClassID: 6, // Druid
			Skill:   types.SkillInsight,
		},
		{
			ID:      48,
			ClassID: 6, // Druid
			Skill:   types.SkillMedicine,
		},
		{
			ID:      49,
			ClassID: 6, // Druid
			Skill:   types.SkillNature,
		},
		{
			ID:      50,
			ClassID: 6, // Druid
			Skill:   types.SkillPerception,
		},
		{
			ID:      51,
			ClassID: 6, // Druid
			Skill:   types.SkillReligion,
		},
		{
			ID:      52,
			ClassID: 6, // Druid
			Skill:   types.SkillSurvival,
		},
		{
			ID:      53,
			ClassID: 7, // Fighter
			Skill:   types.SkillAcrobatics,
		},
		{
			ID:      54,
			ClassID: 7, // Fighter
			Skill:   types.SkillAnimalHandling,
		},
		{
			ID:      55,
			ClassID: 7, // Fighter
			Skill:   types.SkillAthletics,
		},
		{
			ID:      56,
			ClassID: 7, // Fighter
			Skill:   types.SkillHistory,
		},
		{
			ID:      57,
			ClassID: 7, // Fighter
			Skill:   types.SkillInsight,
		},
		{
			ID:      58,
			ClassID: 7, // Fighter
			Skill:   types.SkillIntimidation,
		},
		{
			ID:      59,
			ClassID: 7, // Fighter
			Skill:   types.SkillPersuasion,
		},
		{
			ID:      60,
			ClassID: 7, // Fighter
			Skill:   types.SkillPerception,
		},
		{
			ID:      61,
			ClassID: 7, // Fighter
			Skill:   types.SkillSurvival,
		},
		{
			ID:      62,
			ClassID: 8, // Monk
			Skill:   types.SkillAcrobatics,
		},
		{
			ID:      63,
			ClassID: 8, // Monk
			Skill:   types.SkillAthletics,
		},
		{
			ID:      64,
			ClassID: 8, // Monk
			Skill:   types.SkillHistory,
		},
		{
			ID:      65,
			ClassID: 8, // Monk
			Skill:   types.SkillInsight,
		},
		{
			ID:      66,
			ClassID: 8, // Monk
			Skill:   types.SkillReligion,
		},
		{
			ID:      67,
			ClassID: 8, // Monk
			Skill:   types.SkillStealth,
		},
		{
			ID:      68,
			ClassID: 9, // Paladin
			Skill:   types.SkillAthletics,
		},
		{
			ID:      69,
			ClassID: 9, // Paladin
			Skill:   types.SkillInsight,
		},
		{
			ID:      70,
			ClassID: 9, // Paladin
			Skill:   types.SkillIntimidation,
		},
		{
			ID:      71,
			ClassID: 9, // Paladin
			Skill:   types.SkillMedicine,
		},
		{
			ID:      72,
			ClassID: 9, // Paladin
			Skill:   types.SkillPersuasion,
		},
		{
			ID:      73,
			ClassID: 9, // Paladin
			Skill:   types.SkillReligion,
		},
		{
			ID:      74,
			ClassID: 10, // Ranger
			Skill:   types.SkillAnimalHandling,
		},
		{
			ID:      75,
			ClassID: 10, // Ranger
			Skill:   types.SkillAthletics,
		},
		{
			ID:      76,
			ClassID: 10, // Ranger
			Skill:   types.SkillInsight,
		},
		{
			ID:      77,
			ClassID: 10, // Ranger
			Skill:   types.SkillInvestigation,
		},
		{
			ID:      78,
			ClassID: 10, // Ranger
			Skill:   types.SkillNature,
		},
		{
			ID:      79,
			ClassID: 10, // Ranger
			Skill:   types.SkillPerception,
		},
		{
			ID:      80,
			ClassID: 10, // Ranger
			Skill:   types.SkillStealth,
		},
		{
			ID:      81,
			ClassID: 10, // Ranger
			Skill:   types.SkillSurvival,
		},
		{
			ID:      82,
			ClassID: 11, // Rogue
			Skill:   types.SkillAcrobatics,
		},
		{
			ID:      83,
			ClassID: 11, // Rogue
			Skill:   types.SkillAthletics,
		},
		{
			ID:      84,
			ClassID: 11, // Rogue
			Skill:   types.SkillDeception,
		},
		{
			ID:      85,
			ClassID: 11, // Rogue
			Skill:   types.SkillInsight,
		},
		{
			ID:      86,
			ClassID: 11, // Rogue
			Skill:   types.SkillIntimidation,
		},
		{
			ID:      87,
			ClassID: 11, // Rogue
			Skill:   types.SkillInvestigation,
		},
		{
			ID:      88,
			ClassID: 11, // Rogue
			Skill:   types.SkillPerception,
		},
		{
			ID:      89,
			ClassID: 11, // Rogue
			Skill:   types.SkillPersuasion,
		},
		{
			ID:      90,
			ClassID: 11, // Rogue
			Skill:   types.SkillSleightOfHand,
		},
		{
			ID:      91,
			ClassID: 11, // Rogue
			Skill:   types.SkillStealth,
		},
		{
			ID:      92,
			ClassID: 12, // Sorcerer
			Skill:   types.SkillArcana,
		},
		{
			ID:      93,
			ClassID: 12, // Sorcerer
			Skill:   types.SkillDeception,
		},
		{
			ID:      94,
			ClassID: 12, // Sorcerer
			Skill:   types.SkillInsight,
		},
		{
			ID:      95,
			ClassID: 12, // Sorcerer
			Skill:   types.SkillIntimidation,
		},
		{
			ID:      96,
			ClassID: 12, // Sorcerer
			Skill:   types.SkillPersuasion,
		},
		{
			ID:      97,
			ClassID: 12, // Sorcerer
			Skill:   types.SkillReligion,
		},
		{
			ID:      98,
			ClassID: 13, // Warlock
			Skill:   types.SkillArcana,
		},
		{
			ID:      99,
			ClassID: 13, // Warlock
			Skill:   types.SkillDeception,
		},
		{
			ID:      100,
			ClassID: 13, // Warlock
			Skill:   types.SkillHistory,
		},
		{
			ID:      101,
			ClassID: 13, // Warlock
			Skill:   types.SkillIntimidation,
		},
		{
			ID:      102,
			ClassID: 13, // Warlock
			Skill:   types.SkillInvestigation,
		},
		{
			ID:      103,
			ClassID: 13, // Warlock
			Skill:   types.SkillNature,
		},
		{
			ID:      104,
			ClassID: 13, // Warlock
			Skill:   types.SkillReligion,
		},
		{
			ID:      105,
			ClassID: 14, // Wizard
			Skill:   types.SkillArcana,
		},
		{
			ID:      106,
			ClassID: 14, // Wizard
			Skill:   types.SkillHistory,
		},
		{
			ID:      107,
			ClassID: 14, // Wizard
			Skill:   types.SkillInsight,
		},
		{
			ID:      108,
			ClassID: 14, // Wizard
			Skill:   types.SkillInvestigation,
		},
		{
			ID:      109,
			ClassID: 14, // Wizard
			Skill:   types.SkillMedicine,
		},
		{
			ID:      110,
			ClassID: 14, // Wizard
			Skill:   types.SkillNature,
		},
		{
			ID:      111,
			ClassID: 14, // Wizard
			Skill:   types.SkillReligion,
		},
	}

	for _, csp := range clasSkillProficiencies {
		err := s.DB.Where("id = ?", csp.ID).FirstOrCreate(&csp).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating class skill proficiency with id `%v` in seeder: %v", csp.ID, s.DB.Error))
		}
	}
}
