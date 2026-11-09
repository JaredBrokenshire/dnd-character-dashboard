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
				Description:     "Aasimar (pronounced AH-sih-mar) are mortals who carry a spark of the Upper Planes within their souls. Whether descended from an angelic being or infused with celestial power, they can fan that spark to bring light, healing, and heavenly fury.\n\nAasimar can arise among any population of mortals. They resemble their parents, but they live for up to 160 years and have features that hint at their celestial heritage, such as metallic freckles, luminous eyes, a halo, or the skin color of an angel (silver, opalescent green, or coppery red). These features start subtle and become obvious when the aasimar learns to reveal their full celestial nature.",
				CreatureType:    "Humanoid",
				Size:            "Medium (about 4-7 ft. tall) or Small (about 2-4 ft. tall)",
				BaseSpeed:       30,
			},
			traitIds: []uint{1, 2, 3, 4, 5},
		},
		{
			race: models.Race{
				ID:              2,
				Name:            "Dragonborn",
				CardDescription: "The ancestors of Dragonborn hatched from the eggs of chromatic and metallic dragons.",
				Description:     "The ancestors of dragonborn hatched from the eggs of chromatic and metallic dragons. One story holds that these eggs were blessed by the dragon gods Bahamut and Tiamat, who wanted to populate the multiverse with people created in their image. Another story claims that dragons created the first dragonborn without the gods’ blessings. Whatever their origin, dragonborn have made homes for themselves on the Material Plane.\n\nDragonborn look like wingless, bipedal dragons— scaly, bright-eyed, and thick-boned with horns on their heads—and their coloration and other features are reminiscent of their draconic ancestors.",
				CreatureType:    "Humanoid",
				Size:            "Medium (about 5-7 ft. tall)",
				BaseSpeed:       30,
			},
			traitIds: []uint{6, 7, 8, 1, 9},
		},
		{
			race: models.Race{
				ID:              3,
				Name:            "Dwarf",
				CardDescription: "Dwarves were raised from the earth in the elder days by a deity of the forge.",
				Description:     "Dwarves were raised from the earth in the elder days by a deity of the forge. Called by various names on different worlds—Moradin, Reorx, and others—that god gave dwarves an affinity for stone and metal and for living underground. The god also made them resilient like the mountains, with a life span of about 350 years.\n\nSquat and often bearded, the original dwarves carved cities and strongholds into mountainsides and under the earth. Their oldest legends tell of conflicts with the monsters of mountaintops and the Underdark, whether those monsters were towering giants or subterranean horrors. Inspired by those tales, dwarves of any culture often sing of valorous deeds—especially of the little overcoming the mighty.\n\nOn some worlds in the multiverse, the first settlements of dwarves were built in hills or mountains, and the families who trace their ancestry to those settlements call themselves hill dwarves or mountain dwarves, respectively. The Greyhawk and Dragonlance settings have such communities.",
				CreatureType:    "Humanoid",
				Size:            "Medium (about 4-5 ft. tall)",
				BaseSpeed:       30,
			},
			traitIds: []uint{1, 10, 11, 12},
		},
		{
			race: models.Race{
				ID:              4,
				Name:            "Elf",
				CardDescription: "The elves' curiosity led many of them to explore other planes of existence.",
				Description:     "Created by the god Corellon, the first elves could change their forms at will. They lost this ability when Corellon cursed them for plotting with the deity Lolth, who tried and failed to usurp Corellon’s dominion. When Lolth was cast into the Abyss, most elves renounced her and earned Corellon’s forgiveness, but that which Corellon had taken from them was lost forever.\n\nNo longer able to shape-shift at will, the elves retreated to the Feywild, where their sorrow was deepened by that plane’s influence. Over time, curiosity led many of them to explore other planes of existence, including worlds in the Material Plane.\n\nElves have pointed ears and lack facial and body hair. They live for around 750 years, and they don’t sleep but instead enter a trance when they need to rest. In that state, they remain aware of their surroundings while immersing themselves in memories and meditations.\n\nAn environment subtly transforms elves after they inhabit it for a millennium or more, and it grants them certain kinds of magic. Drow, high elves, and wood elves are examples of elves who have been transformed thus.",
				CreatureType:    "Humanoid",
				Size:            "Medium (about 5-6 ft. tall)",
				BaseSpeed:       30,
			},
			traitIds: []uint{1, 13, 14, 15, 16},
		},
		{
			race: models.Race{
				ID:              5,
				Name:            "Gnome",
				CardDescription: "Gnomes are magical folk created by gods of invention, illusions, and life underground.",
				Description:     "Gnomes are magical folk created by gods of invention, illusions, and life underground. The earliest gnomes were seldom seen by other folk due to the gnomes’ secretive nature and their propensity for living in forests and burrows. What they lacked in size, they made up for in cleverness. They confounded predators with traps and labyrinthine tunnels. They also learned magic from gods like Garl Glittergold, Baervan Wildwanderer, and Baravar Cloakshadow, who visited them in disguise. That magic eventually created the lineages of forest gnomes and rock gnomes.\n\nGnomes are petite folk with big eyes and pointed ears, who live around 425 years. Many gnomes like the feeling of a roof over their head, even if that “roof” is nothing more than a hat.",
				CreatureType:    "Humanoid",
				Size:            "Small (about 3-4 ft. tall)",
				BaseSpeed:       30,
			},
			traitIds: []uint{1, 17, 18},
		},
		{
			race: models.Race{
				ID:              6,
				Name:            "Goliath",
				CardDescription: "Goliaths are distant descendants of giants and seek heights above those reached by their ancestors.",
				Description:     "Towering over most folk, goliaths are distant descendants of giants. Each goliath bears the favors of the first giants—favors that manifest in various supernatural boons, including the ability to quickly grow and temporarily approach the height of goliaths’ gigantic kin.\n\nGoliaths have physical characteristics that are reminiscent of the giants in their family lines. For example, some goliaths look like stone giants, while others resemble fire giants. Whatever giants they count as kin, goliaths have forged their own path in the multiverse—unencumbered by the internecine conflicts that have ravaged giantkind for ages—and seek heights above those reached by their ancestors.",
				CreatureType:    "Humanoid",
				Size:            "Medium (7-8 ft. tall)",
				BaseSpeed:       35,
			},
			traitIds: []uint{19, 20, 21},
		},
		{
			race: models.Race{
				ID:              7,
				Name:            "Halfling",
				CardDescription: "Halflings possess a brave and adventurous spirit that leads them on journeys of discovery.",
				Description:     "Cherished and guided by gods who value life, home, and hearth, halflings gravitate toward bucolic havens where family and community help shape their lives. That said, many halflings possess a brave and adventurous spirit that leads them on journeys of discovery, affording them the chance to explore a bigger world and make new friends along the way. Their size—similar to that of a human child—helps them pass through crowds unnoticed and slip through tight spaces.\n\nAnyone who has spent time around halflings, particularly halfling adventurers, has likely witnessed the storied “luck of the halflings” in action. When a halfling is in mortal danger, an unseen force seems to intervene on the halfling’s behalf. Many halflings believe in the power of luck, and they attribute their unusual gift to one or more of their benevolent gods, including Yondalla, Brandobaris, and Charmalaine. The same gift might contribute to their robust life spans (about 150 years).\n\nHalfling communities come in all varieties. For every sequestered shire tucked away in an unspoiled part of the world, there’s a crime syndicate like the Boromar Clan in the Eberron setting or a territorial mob of halflings like those in the Dark Sun setting.\n\nHalflings who prefer to live underground are sometimes called strongheart halflings or stouts. Nomadic halflings, as well as those who live among humans and other tall folk, are sometimes called lightfoot halflings or tallfellows.",
				CreatureType:    "Humanoid",
				Size:            "Small (about 2-3 ft. tall)",
				BaseSpeed:       30,
			},
			traitIds: []uint{22, 23, 24, 25},
		},
		{
			race: models.Race{
				ID:              8,
				Name:            "Human",
				CardDescription: "Found throughout the multiverse, humans are as varied as they are numerous.",
				Description:     "Found throughout the multiverse, humans are as varied as they are numerous, and they endeavor to achieve as much as they can in the years they are given. Their ambition and resourcefulness are commended, respected, and feared on many worlds.\n\nHumans are as diverse in appearance as the people of Earth, and they have many gods. Scholars dispute the origin of humanity, but one of the earliest known human gatherings is said to have occurred in Sigil, the torus-shaped city at the center of the multiverse and the place where the Common language was born. From there, humans could have spread to every part of the multiverse, bringing the City of Doors’ cosmopolitanism with them.",
				CreatureType:    "Humanoid",
				Size:            "Medium (about 4-7 ft. tall) or Small (about 2-4 ft. tall)",
				BaseSpeed:       30,
			},
			traitIds: []uint{26, 27, 28},
		},
		{
			race: models.Race{
				ID:              9,
				Name:            "Orc",
				CardDescription: "Orcs are equipped with gifts to help them wander great plains, vast caverns, and churning seas.",
				Description:     "Orcs trace their creation to Gruumsh, a powerful god who roamed the wide open spaces of the Material Plane. Gruumsh equipped his children with gifts to help them wander great plains, vast caverns, and churning seas and to face the monsters that lurk there. Even when they turn their devotion to other gods, orcs retain Gruumsh’s gifts: endurance, determination, and the ability to see in darkness.\n\nOrcs are, on average, tall and broad. They have gray skin, ears that are sharply pointed, and prominent lower canines that resemble small tusks. Orc youths on some worlds are told about their ancestors’ great travels and travails. Inspired by those tales, many of those orcs wonder when Gruumsh will call on them to match the heroic deeds of old and if they will prove worthy of his favor. Other orcs are happy to leave old tales in the past and find their own way.",
				CreatureType:    "Humanoid",
				Size:            "Medium (about 6-7 ft. tall)",
				BaseSpeed:       30,
			},
			traitIds: []uint{29, 1, 30},
		},
		{
			race: models.Race{
				ID:              10,
				Name:            "Tiefling",
				CardDescription: "Tieflings are either born in the Lower Planes or have fiendish ancestors who originated there.",
				Description:     "Tieflings are either born in the Lower Planes or have fiendish ancestors who originated there. A tiefling (pronounced TEE-fling) is linked by blood to a devil, a demon, or some other Fiend. This connection to the Lower Planes is the tiefling’s fiendish legacy, which comes with the promise of power yet has no effect on the tiefling’s moral outlook.\n\nA tiefling chooses whether to embrace or lament their fiendish legacy. The three legacies are described below.",
				CreatureType:    "Humanoid",
				Size:            "Medium (about 4-6 ft. tall) or Small (about 3-4 ft. tall)",
				BaseSpeed:       30,
			},
			traitIds: []uint{1, 31, 32},
		},
		{
			race: models.Race{
				ID:              11,
				Name:            "Half-Elf",
				CardDescription: "Half-elves combine what some say are the best qualities of their elf and human parents.",
				Description:     "Walking in two worlds but truly belonging to neither, half-elves combine what some say are the best qualities of their elf and human parents: human curiosity, inventiveness, and ambition tempered by the refined senses, love of nature, and artistic tastes of the elves. Some half-elves live among humans, set apart by their emotional and physical differences, watching friends and loved ones age while time barely touches them. Others live with the elves, growing restless as they reach adulthood in the timeless elven realms, while their peers continue to live as children. Many half-elves, unable to fit into either society, choose lives of solitary wandering or join with other misfits and outcasts in the adventuring life.",
				CreatureType:    "Humanoid",
				Size:            "Medium (about 5-6 ft. tall)",
				BaseSpeed:       30,
			},
			traitIds: []uint{1, 14, 28},
		},
		{
			race: models.Race{
				ID:              12,
				Name:            "Half-Orc",
				CardDescription: "Some half-orcs rise to become proud leaders or orc communities. Some venture into the world to prove their worth. Many of these become adventurers, achieving greatness for their mighty deeds.",
				Description:     "Whether united under the leadership of a mighty warlock or having fought to a standstill after years of conflict, orc and human communities, sometimes form alliances. When these alliances are sealed by marriages, half-orcs are born. Some half-orcs rise to become proud leaders of orc communities. Some venture into the world to prove their worth. Many of these become adventurers, achieving greatness for their mighty deeds.",
				CreatureType:    "Humanoid",
				Size:            "Medium (about 5-7 ft. tall)",
				BaseSpeed:       30,
			},
			traitIds: []uint{1, 30},
		},
		{
			race: models.Race{
				ID:              13,
				Name:            "Air Genasi",
				CardDescription: "Embodying many of the airy traits of their otherworldly ancestors, air genasi can draw upon their connection to the winds.",
				Description:     "Air genasi are descended from djinn, the genies of the Elemental Plane of Air. Embodying many of the airy traits of their otherworldly ancestors, air genasi can draw upon their connection to the winds.\n\nAir genasi’s skin tones include many shades of blue, along with the full range of human skin tones, with bluish or ashen casts. Sometimes their skin is marked by lines that seem like cracks with bluish-white energy spilling out. An air genasi’s hair might blow in a phantom wind or be made entirely of clouds or vapor.",
				CreatureType:    "Humanoid",
				Size:            "Medium or Small",
				BaseSpeed:       35,
			},
			traitIds: []uint{1, 42, 43, 44},
		},
		{
			race: models.Race{
				ID:              14,
				Name:            "Earth Genasi",
				CardDescription: "Tracing their ancestry to dao, the genies of the Elemental Plane, earth genasi inherit dao's steadfast strength and control over earth.",
				Description:     "Tracing their ancestry to dao, the genies of the Elemental Plane of Earth, earth genasi inherit dao’s steadfast strength and control over earth.\n\nAn earth genasi’s skin can be the colors of stone and earth or a human skin tone with glittering sparkles like gem dust. Some earth genasi have lines marking their skin like cracks, either showing glimmering gemlike veins or a dim, yellowish glow. Earth genasi hair can appear carved of stone or crystal or resemble strands of spun metal.",
				CreatureType:    "Humanoid",
				Size:            "Medium or Small",
				BaseSpeed:       30,
			},
			traitIds: []uint{1, 45, 46},
		},
		{
			race: models.Race{
				ID:              15,
				Name:            "Fire Genasi",
				CardDescription: "Descended from efreet, the genies of the Elemental Plane of Fire, fire genasi channel the flamboyant and often destructive nature of flame.",
				Description:     "Descended from efreet, the genies of the Elemental Plane of Fire, fire genasi channel the flamboyant and often destructive nature of flame. They show their heritage in their skin tones, which can range from deep charcoal to shades of red and orange. Some bear skin tones common to humanity but with fiery marks, such as slowly swirling lights under their skin that resemble embers or glowing red lines tracing over their bodies like cracks. Fire genasi hair can resemble threads of fire or sooty smoke.",
				CreatureType:    "Humanoid",
				Size:            "Medium or Small",
				BaseSpeed:       30,
			},
			traitIds: []uint{1, 47, 48},
		},
		{
			race: models.Race{
				ID:              16,
				Name:            "Water Genasi",
				CardDescription: "Water genasi are perfectly suited to life underwater and carry the power of the waves inside themselves.",
				Description:     "Water genasi descend from marids, aquatic genies from the Elemental Plane of Water. Water genasi are perfectly suited to life underwater and carry the power of the waves inside themselves.\n\nTheir skin is often shades of blue or green, sometimes a blend of the two. If they have a human skin tone, there is a glistening texture that catches the light, like water droplets or nearly invisible fish scales. Their hair can resemble seaweed, waving as if in a current, or it can even be like water itself.",
				CreatureType:    "Humanoid",
				Size:            "Medium or Small",
				BaseSpeed:       30,
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
				Description:     "Some of the smallest draconic creatures in the multiverse, kobolds display their draconic ancestry in the glint of their scales and in their roars. Legends tell of the first kobolds emerging from the Underdark near the lairs of the earliest dragons. In some lands, kobolds serve chromatic or metallic dragons—even worshiping them as divine beings. In other places, kobolds know too well how dangerous those dragons can be and help others defend against draconic destruction.\n\nWhatever their relationship to dragons, kobold scales tend to be rust colored, although the occasional kobold sports a scale color more akin to that of a chromatic or a metallic dragon. A kobold’s cry can express a range of emotion: anger, resolve, elation, fear, and more. Regardless of the emotion expressed, their cry resonates with draconic power.",
				CreatureType:    "Humanoid",
				Size:            "Small",
				BaseSpeed:       30,
			},
			traitIds: []uint{1, 49, 50},
		},
		{
			race: models.Race{
				ID:              19,
				Name:            "Dhampir",
				CardDescription: "Dhampirs are living people who possess vampiric prowess but are cursed with macabre hunger.",
				Description:     "Poised between the worlds of the living and the dead, dhampirs retain their grip on life yet are endlessly tested by vicious hungers. Their ties to the undead grant dhampirs a taste of a vampire’s deathless prowess in the form of increased speed, darkvision, and a life-draining bite.\n\nWith unique insights into the nature of the undead, many dhampirs become adventurers and monster hunters. Their reasons are often deeply personal. Some seek danger, imagining monsters as personifications of their own hungers. Others pursue revenge against whatever turned them into a dhampir. And still others embrace the solitude of the hunt, striving to distance themselves from those who’d tempt their hunger.",
				CreatureType:    "Humanoid",
				Size:            "Medium or Small",
				BaseSpeed:       35,
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
