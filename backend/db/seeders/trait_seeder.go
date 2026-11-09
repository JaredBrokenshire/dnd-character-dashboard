package seeders

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/pkg/logging"
	"fmt"
)

func (s *Seeder) SetTraits() {
	traits := []models.Trait{
		{
			ID:   1,
			Name: "Darkvision",
		},
		{
			ID:   2,
			Name: "Celestial Resilience",
		},
		{
			ID:   3,
			Name: "Healing Hands",
		},
		{
			ID:   4,
			Name: "Light Bearer",
		},
		{
			ID:   5,
			Name: "Celestial Revelation",
		},
		{
			ID:   6,
			Name: "Dragon Ancestry",
		},
		{
			ID:   7,
			Name: "Breath Weapon",
		},
		{
			ID:   8,
			Name: "Damage Resistance",
		},
		{
			ID:   9,
			Name: "Draconic Flight",
		},
		{
			ID:   10,
			Name: "Dwarven Resilience",
		},
		{
			ID:   11,
			Name: "Dwarven Toughness",
		},
		{
			ID:   12,
			Name: "Stonecunning",
		},
		{
			ID:   13,
			Name: "Elven Lineage",
		},
		{
			ID:   14,
			Name: "Fey Ancestry",
		},
		{
			ID:   15,
			Name: "Keen Senses",
		},
		{
			ID:   16,
			Name: "Trance",
		},
		{
			ID:   17,
			Name: "Gnomish Cunning",
		},
		{
			ID:   18,
			Name: "Gnomish Lineage",
		},
		{
			ID:   19,
			Name: "Giant Ancestry",
		},
		{
			ID:   20,
			Name: "Large Form",
		},
		{
			ID:   21,
			Name: "Powerful Build",
		},
		{
			ID:   22,
			Name: "Brave",
		},
		{
			ID:   23,
			Name: "Halfling Nimbleness",
		},
		{
			ID:   24,
			Name: "Luck",
		},
		{
			ID:   25,
			Name: "Naturally Stealthy",
		},
		{
			ID:   26,
			Name: "Resourceful",
		},
		{
			ID:   27,
			Name: "Skillful",
		},
		{
			ID:   28,
			Name: "Versatile",
		},
		{
			ID:   29,
			Name: "Adrenaline Rush",
		},
		{
			ID:   30,
			Name: "Relentless Endurance",
		},
		{
			ID:   31,
			Name: "Fiendish Legacy",
		},
		{
			ID:   32,
			Name: "Otherworldly Presence",
		},
		{
			ID:   33,
			Name: "Spider Climb",
		},
		{
			ID:   34,
			Name: "Trace of Undeath",
		},
		{
			ID:   35,
			Name: "Vampiric Bite",
		},
		{
			ID:   36,
			Name: "Construct",
		},
		{
			ID:   37,
			Name: "Construct Resilience",
		},
		{
			ID:   38,
			Name: "Integrated Protection",
		},
		{
			ID:   39,
			Name: "Sentry's Rest",
		},
		{
			ID:   40,
			Name: "Specialised Design",
		},
		{
			ID:   41,
			Name: "Tireless",
		},
		{
			ID:   42,
			Name: "Unending Breath",
		},
		{
			ID:   43,
			Name: "Lightning Resistance",
		},
		{
			ID:   44,
			Name: "Mingle with the Wind",
		},
		{
			ID:   45,
			Name: "Earth Walk",
		},
		{
			ID:   46,
			Name: "Merge with Stone",
		},
		{
			ID:   47,
			Name: "Fire Resistance",
		},
		{
			ID:   48,
			Name: "Reach to the Blaze",
		},
		{
			ID:   49,
			Name: "Draconic Cry",
		},
		{
			ID:   50,
			Name: "Kobold Legacy",
		},
		{
			ID:   51,
			Name: "Acid Resistance",
		},
		{
			ID:   52,
			Name: "Amphibious",
		},
		{
			ID:   53,
			Name: "Call to the Wave",
		},
	}

	for _, trait := range traits {
		err := s.DB.Where("id = ?", trait.ID).FirstOrCreate(&trait).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating trait with id `%v` in seeder: %v", trait.ID, err.Error()))
		}
	}
}
