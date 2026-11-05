package seeders

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/pkg/logging"
	"fmt"
)

type raceSeed struct {
	race     models.Race
	traitIds []uint
}

func (s *Seeder) SetRaces() {
	races := []raceSeed{
		{
			race: models.Race{
				ID:              1,
				Name:            "Aasimar",
				CardDescription: "Aasimar are mortals who carry a spark of the Upper Planes within their souls.",
			},
			traitIds: []uint{1, 2, 3, 4, 5},
		},
		{
			race: models.Race{
				ID:              2,
				Name:            "Dragonborn",
				CardDescription: "The ancestors of Dragonborn hatched from the eggs of chromatic and metallic dragons.",
			},
			traitIds: []uint{6, 7, 8, 1, 9},
		},
		{
			race: models.Race{
				ID:              3,
				Name:            "Dwarf",
				CardDescription: "Dwarves were raised from the earth in the elder days by a deity of the forge.",
			},
			traitIds: []uint{1, 10, 11, 12},
		},
		{
			race: models.Race{
				ID:              4,
				Name:            "Elf",
				CardDescription: "The elves' curiosity led many of them to explore other planes of existence.",
			},
			traitIds: []uint{1, 13, 14, 15, 16},
		},
		{
			race: models.Race{
				ID:              5,
				Name:            "Gnome",
				CardDescription: "Gnomes are magical folk created by gods of invention, illusions, and life underground.",
			},
			traitIds: []uint{1, 17, 18},
		},
		{
			race: models.Race{
				ID:              6,
				Name:            "Goliath",
				CardDescription: "Goliaths are distant descendants of giants and seek heights above those reached by their ancestors.",
			},
			traitIds: []uint{19, 20, 21},
		},
		{
			race: models.Race{
				ID:              7,
				Name:            "Halfling",
				CardDescription: "Halflings possess a brave and adventurous spirit that leads them on journeys of discovery.",
			},
			traitIds: []uint{22, 23, 24, 25},
		},
		{
			race: models.Race{
				ID:              8,
				Name:            "Human",
				CardDescription: "Found throughout the multiverse, humans are as varied as they are numerous.",
			},
			traitIds: []uint{26, 27, 28},
		},
		{
			race: models.Race{
				ID:              9,
				Name:            "Orc",
				CardDescription: "Orcs are equipped with gifts to help them wander great plains, vast caverns, and churning seas.",
			},
			traitIds: []uint{29, 1, 30},
		},
		{
			race: models.Race{
				ID:              10,
				Name:            "Tiefling",
				CardDescription: "Tieflings are either born in the Lower Planes or have fiendish ancestors who originated there.",
			},
			traitIds: []uint{1, 31, 32},
		},
		{
			race: models.Race{
				ID:              11,
				Name:            "Half-Elf",
				CardDescription: "Half-elves combine what some say are the best qualities of their elf and human parents.",
			},
			traitIds: []uint{1, 14, 28},
		},
		{
			race: models.Race{
				ID:              12,
				Name:            "Half-Orc",
				CardDescription: "Some half-orcs rise to become proud leaders or orc communities. Some venture into the world to prove their worth. Many of these become adventurers, achieving greatness for their mighty deeds.",
			},
			traitIds: []uint{1, 30},
		},
		{
			race: models.Race{
				ID:              13,
				Name:            "Air Genasi",
				CardDescription: "Embodying many of the airy traits of their otherworldly ancestors, air genasi can draw upon their connection to the winds.",
			},
			traitIds: []uint{1, 42, 43, 44},
		},
		{
			race: models.Race{
				ID:              14,
				Name:            "Earth Genasi",
				CardDescription: "Tracing their ancestry to dao, the genies of the Elemental Plane, earth genasi inherit dao's steadfast strength and control over earth.",
			},
			traitIds: []uint{1, 45, 46},
		},
		{
			race: models.Race{
				ID:              15,
				Name:            "Fire Genasi",
				CardDescription: "Descended from efreet, the genies of the Elemental Plane of Fire, fire genasi channel the flamboyant and often destructive nature of flame.",
			},
			traitIds: []uint{1, 47, 48},
		},
		{
			race: models.Race{
				ID:              16,
				Name:            "Water Genasi",
				CardDescription: "Water genasi are perfectly suited to life underwater and carry the power of the waves inside themselves.",
			},
			traitIds: []uint{51, 52, 53, 1},
		},
		{
			race: models.Race{
				ID:              17,
				Name:            "Warforged",
				CardDescription: "Warforged are mechanical beings built as weapons to fight in the Last War.",
			},
			traitIds: []uint{36, 37, 38, 39, 40, 41},
		},
		{
			race: models.Race{
				ID:              18,
				Name:            "Kobold",
				CardDescription: "Some of the smallest draconic creatures in the multiverse, kobolds display their draconic ancestry in the glint of their scales and in their roars.",
			},
			traitIds: []uint{1, 49, 50},
		},
		{
			race: models.Race{
				ID:              19,
				Name:            "Dhampir",
				CardDescription: "Dhampirs are living people who possess vampiric prowess but are cursed with macabre hunger.",
			},
			traitIds: []uint{1, 33, 34, 35},
		},
	}

	for _, seed := range races {
		existing := models.Race{}
		if err := s.DB.Where("id = ?", seed.race.ID).First(&existing).Error; err == nil {
			// Already exists — skip to keep seeder idempotent.
			continue
		}

		if err := s.DB.Create(&seed.race).Error; err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating race with id `%v` in seeder: %v", seed.race.ID, err.Error()))
			continue
		}

		// Build ID-only slices for the join tables. Using Association.Append
		traits := make([]models.Trait, len(seed.traitIds))
		for i, id := range seed.traitIds {
			traits[i] = models.Trait{ID: id}
		}

		if err := s.DB.Model(&seed.race).Association("Traits").Append(traits).Error; err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error assigning traits to race `%v` in seeder: %v", seed.race.ID, err.Error()))
		}

	}
}
