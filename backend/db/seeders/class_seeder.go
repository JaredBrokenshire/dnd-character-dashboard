package seeders

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/pkg/logging"
	"fmt"
)

func (s *Seeder) SetClasses() {
	classes := []models.Class{
		{
			ID:               1,
			Name:             "Artificer",
			Tagline:          "A Magical Crafter and an Inventor of Wonders",
			CardDescription:  "Artificers use ingenuity and magic to unlock extraordinary capabilities in objects.",
			PrimaryAbility:   "Intelligence",
			HitPointDieValue: 8,
			Save:             models.Constitution,
			SecondarySave:    models.Intelligence,
		},
		{
			ID:               2,
			Name:             "Barbarian",
			Tagline:          "A Fierce Warrior of Primal Rage",
			CardDescription:  "Barbarians are mighty warriors who are powered by primal forces of the multiverse that manifest as a Rage.",
			PrimaryAbility:   "Strength",
			HitPointDieValue: 12,
			Save:             models.Strength,
			SecondarySave:    models.Constitution,
		},
		{
			ID:               3,
			Name:             "Bard",
			Tagline:          "An Inspiring Performer of Music, Dance, and Magic",
			CardDescription:  "Bards are experts at inspiring others, soothing hurts, disheartening foes, and creating illusions.",
			PrimaryAbility:   "Charisma",
			HitPointDieValue: 8,
			Save:             models.Dexterity,
			SecondarySave:    models.Charisma,
		},
		{
			ID:               4,
			Name:             "Blood Hunter",
			CardDescription:  "Willing to suffer whatever it takes to achieve victory, these adept warriors have forged themselves into a potent force dedicated to protecting the innocent.",
			PrimaryAbility:   "Strength or Dexterity & Intelligence or Wisdom",
			HitPointDieValue: 10,
			Save:             models.Dexterity,
			SecondarySave:    models.Intelligence,
		},
		{
			ID:               5,
			Name:             "Cleric",
			Tagline:          "A Miraculous Priest of Divine Power",
			CardDescription:  "Clerics can reach out to the divine magic of the Outer Planes and channel it to bolster people and battle foes.",
			PrimaryAbility:   "Wisdom",
			HitPointDieValue: 8,
			Save:             models.Wisdom,
			SecondarySave:    models.Charisma,
		},
		{
			ID:               6,
			Name:             "Druid",
			Tagline:          "A Nature Priest of Primal Power",
			CardDescription:  "Druids call on the forces of nature, harnessing magic to heal, transform into animals, and wield elemental destruction.",
			PrimaryAbility:   "Wisdom",
			HitPointDieValue: 8,
			Save:             models.Intelligence,
			SecondarySave:    models.Wisdom,
		},
		{
			ID:               7,
			Name:             "Fighter",
			Tagline:          "A Master of All Arms and Armour",
			CardDescription:  "Fighters all share an unparalleled prowess with weapons and armour and are well acquainted with death, both meting it out and defying it.",
			PrimaryAbility:   "Strength or Dexterity",
			HitPointDieValue: 10,
			Save:             models.Strength,
			SecondarySave:    models.Constitution,
		},
		{
			ID:               8,
			Name:             "Monk",
			Tagline:          "A Martial Artist of Supernatural Focus",
			CardDescription:  "Monks focus their internal reservoirs of power to create extraordinary, even supernatural effects.",
			PrimaryAbility:   "Dexterity & Wisdom",
			HitPointDieValue: 8,
			Save:             models.Strength,
			SecondarySave:    models.Dexterity,
		},
		{
			ID:               9,
			Name:             "Paladin",
			Tagline:          "A Devout Warrior of Sacred Oaths",
			CardDescription:  "Paladins live on the front lines of the cosmic struggle, united by their oaths against the forces of annihilation.",
			PrimaryAbility:   "Strength & Charisma",
			HitPointDieValue: 10,
			Save:             models.Wisdom,
			SecondarySave:    models.Charisma,
		},
		{
			ID:               10,
			Name:             "Ranger",
			Tagline:          "A Wandering Warrior Imbued with Primal Magic",
			CardDescription:  "Rangers are honed with deadly focus and harness primal powers to protect the world from the ravages of monsters and tyrants.",
			PrimaryAbility:   "Dexterity & Wisdom",
			HitPointDieValue: 10,
			Save:             models.Strength,
			SecondarySave:    models.Dexterity,
		},
		{
			ID:               11,
			Name:             "Rogue",
			Tagline:          "A Dexterous Expert in Stealth and Subterfuge",
			CardDescription:  "Rogues have a knack for finding the solution to just about any problem, prioritising subtle strikes over brute strength.",
			PrimaryAbility:   "Dexterity",
			HitPointDieValue: 8,
			Save:             models.Dexterity,
			SecondarySave:    models.Intelligence,
		},
		{
			ID:               12,
			Name:             "Sorcerer",
			Tagline:          "A Dazzling Mage Filled with Innate Magic",
			CardDescription:  "Sorcerers harness and channel the raw, roiling power of innate magic that is stamped into their very being.",
			PrimaryAbility:   "Charisma",
			HitPointDieValue: 6,
			Save:             models.Constitution,
			SecondarySave:    models.Charisma,
		},
		{
			ID:               13,
			Name:             "Warlock",
			Tagline:          "An Occultist Empowered by Otherworldly Pacts",
			CardDescription:  "Warlocks quest for knowledge that lies hidden in the fabric of the multiverse, piecing together arcane secrets to bolster their own power.",
			PrimaryAbility:   "Charisma",
			HitPointDieValue: 8,
			Save:             models.Wisdom,
			SecondarySave:    models.Charisma,
		},
		{
			ID:               14,
			Name:             "Wizard",
			Tagline:          "A Scholarly Magic-User of Arcane Power",
			CardDescription:  "Wizards cast spells of explosive fire, arcing lightning, subtle deception, and spectacular transformations.",
			PrimaryAbility:   "Intelligence",
			HitPointDieValue: 6,
			Save:             models.Intelligence,
			SecondarySave:    models.Wisdom,
		},
	}

	for _, class := range classes {
		err := s.DB.Where("id = ?", class.ID).FirstOrCreate(&class).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating class with id `%v` in seeder: %v", class.ID, err.Error()))
		}
	}
}
