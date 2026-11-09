package seeders

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/pkg/logging"
	"fmt"
)

func (s *Seeder) SetRaceVariants() {
	raceVariants := []models.RaceVariant{
		{
			ID:          1,
			RaceID:      4,
			Name:        "Drow",
			Description: "Drow typically dwell in the Underdark and have been shaped by it. Some drow individuals and societies avoid the Underdark altogether yet carry its magic. In the Eberron setting, for example, drow dwell in rainforests and cyclopean ruins on the continent of Xen’drik.",
		},
		{
			ID:          2,
			RaceID:      4,
			Name:        "High Elves",
			Description: "High elves have been infused with the magic of crossings between the Feywild and the Material Plane. On some worlds, high elves refer to themselves by other names. For example, they call themselves sun or moon elves in the Forgotten Realms setting, Silvanesti and Qualinesti in the Dragonlance setting, and Aereni in the Eberron setting.",
		},
		{
			ID:          3,
			RaceID:      4,
			Name:        "Wood Elves",
			Description: "Wood elves carry the magic of primeval forests within themselves. They are known by many other names, including wild elves, green elves, and forest elves. Grugach are reclusive wood elves of the Greyhawk setting, while the Kagonesti and the Tairnadal are wood elves of the Dragonlance and Eberron settings, respectively.",
		},
		{
			ID:          4,
			RaceID:      10,
			Name:        "Abyssal",
			Description: "The entropy of the Abyss, the chaos of Pandemonium, and the despair of Carceri call to tieflings who have the abyssal legacy. Horns, fur, tusks, and peculiar scents are common physical features of such tieflings, most of whom have the blood of demons coursing through their veins.",
		},
		{
			ID:          5,
			RaceID:      10,
			Name:        "Chthonic",
			Description: "Tieflings who have the chthonic legacy feel not only the tug of Carceri but also the greed of Gehenna and the gloom of Hades. Some of these tieflings look cadaverous. Others possess the unearthly beauty of a succubus, or they have physical features in common, with a night hag, a yugoloth, or some other Neutral Evil fiendish ancestor.",
		},
		{
			ID:          6,
			RaceID:      10,
			Name:        "Infernal",
			Description: "The infernal legacy connects tieflings not only to Gehenna but also the Nine Hells and the raging battlefields of Acheron. Horns, spines, tails, golden eyes, and a faint odor of sulfur or smoke are common physical features of such tieflings, most of who trace their ancestry to devils.",
		},
		{
			ID:          7,
			RaceID:      2,
			Name:        "Chromatic Dragonborn",
			Description: "Dragonborn with chromatic ancestry claim the raw elemental power of chromatic dragons. The vibrant colors of black, blue, green, red, and white dragons gleam in those dragonborn’s scaled skin and in the deadly energy of their breath weapons. Theirs is the raw elemental fury of the volcano, of biting arctic winds, and of raging lightning storms, as well as the subtle whisper of swamp and forest, toxic and corrosive.",
		},
	}

	for _, raceVariant := range raceVariants {
		err := s.DB.Where("id = ?", raceVariant.ID).FirstOrCreate(&raceVariant).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating race variant with id `%v` in seeder: %v", raceVariant.ID, err.Error()))
		}
	}
}
