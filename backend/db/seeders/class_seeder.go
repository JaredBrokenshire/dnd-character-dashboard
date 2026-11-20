package seeders

import (
	"dnd/db/models"
	"dnd/pkg/logging"
	"fmt"
)

func (s *Seeder) SetClasses() {
	classes := []models.Class{
		{
			ID:                         1,
			Name:                       "Artificer",
			Tagline:                    "A Magical Crafter and an Inventor of Wonders",
			CardDescription:            "Artificers use ingenuity and magic to unlock extraordinary capabilities in objects.",
			Description:                "Masters of invention, artificers use ingenuity and magic to unlock extraordinary capabilities in objects. They see magic as a complex system waiting to be decoded and then harnessed in their spells and inventions. You can find everything you need to play one of these inventors in the next few sections.\n\nArtificers use a variety of tools to channel their arcane power. To cast a spell, an artificer might use alchemist's supplies to create a potent elixir, calligrapher's supplies to inscribe a sigil of power, or tinker's tools to craft a temporary charm. The magic of artificers is tied to their tools and their talents, and few other characters can produce the right tool for a job as well as an artificer.",
			PrimaryAbility:             "Intelligence",
			HitPointDieValue:           8,
			SkillProficiencyMaxOptions: 2,
		},
		{
			ID:                         2,
			Name:                       "Barbarian",
			Tagline:                    "A Fierce Warrior of Primal Rage",
			CardDescription:            "Barbarians are mighty warriors who are powered by primal forces of the multiverse that manifest as a Rage.",
			Description:                "Barbarians are mighty warriors who are powered by primal forces of the multiverse that manifest as a Rage. More than a mere emotion—and not limited to anger—this Rage is an incarnation of a predator’s ferocity, a storm’s fury, and a sea’s turmoil.\n\nSome Barbarians personify their Rage as a fierce spirit or revered forebear. Others see it as a connection to the pain and anguish of the world, as an impersonal tangle of wild magic, or as an expression of their own deepest self. For every Barbarian, their Rage is a power that fuels not just battle prowess, but also uncanny reflexes and heightened senses.\n\nBarbarians often serve as protectors and leaders in their communities. They charge headlong into danger so those under their protection don’t have to. Their courage in the face of danger makes Barbarians perfectly suited for adventure.",
			PrimaryAbility:             "Strength",
			HitPointDieValue:           12,
			SkillProficiencyMaxOptions: 2,
		},
		{
			ID:                         3,
			Name:                       "Bard",
			Tagline:                    "An Inspiring Performer of Music, Dance, and Magic",
			CardDescription:            "Bards are experts at inspiring others, soothing hurts, disheartening foes, and creating illusions.",
			Description:                "Invoking magic through music, dance, and verse, Bards are expert at inspiring others, soothing hurts, disheartening foes, and creating illusions. Bards believe the multiverse was spoken into existence and that remnants of its Words of Creation still resound and glimmer on every plane of existence. Bardic magic attempts to harness those words, which transcend any language.\n\nAnything can inspire a new song or tale, so Bards are fascinated by almost everything. They become masters of many things, including performing music, working magic, and making jests.\n\nA Bard’s life is spent traveling, gathering lore, telling stories, and living on the gratitude of audiences, much like any other entertainer. But Bards’ depth of knowledge and mastery of magic sets them apart.",
			PrimaryAbility:             "Charisma",
			HitPointDieValue:           8,
			SkillProficiencyMaxOptions: 3,
		},
		{
			ID:                         4,
			Name:                       "Blood Hunter",
			CardDescription:            "Willing to suffer whatever it takes to achieve victory, these adept warriors have forged themselves into a potent force dedicated to protecting the innocent.",
			Description:                "Blood hunters are clever warriors driven by an unending determination to destroy evils old and new. Armed with rites of secretive blood magic and a willingness to sacrifice their own vitality and humanity for their cause, they protect the realms from the shadows — even as they remain ever vigilant against being drawn to the darkness that consumes the monsters they hunt.",
			PrimaryAbility:             "Strength or Dexterity & Intelligence or Wisdom",
			HitPointDieValue:           10,
			SkillProficiencyMaxOptions: 3,
		},
		{
			ID:                         5,
			Name:                       "Cleric",
			Tagline:                    "A Miraculous Priest of Divine Power",
			CardDescription:            "Clerics can reach out to the divine magic of the Outer Planes and channel it to bolster people and battle foes.",
			Description:                "Clerics draw power from the realms of the gods and harness it to work miracles. Blessed by a deity, a pantheon, or another immortal entity, a Cleric can reach out to the divine magic of the Outer Planes—where gods dwell—and channel it to bolster people and battle foes.\n\nBecause their power is a divine gift, Clerics typically associate themselves with temples dedicated to the deity or other immortal force that unlocked their magic. Harnessing divine magic doesn’t rely on specific training, yet Clerics might learn prayers and rites that help them draw on power from the Outer Planes.\n\nNot every member of a temple or shrine is a Cleric. Some priests are called to a simple life of temple service, carrying out their devotion through prayer and rituals, not through magic. Many mortals claim to speak for the gods, but few can marshal the power of those gods the way a Cleric can.",
			PrimaryAbility:             "Wisdom",
			HitPointDieValue:           8,
			SkillProficiencyMaxOptions: 2,
		},
		{
			ID:                         6,
			Name:                       "Druid",
			Tagline:                    "A Nature Priest of Primal Power",
			CardDescription:            "Druids call on the forces of nature, harnessing magic to heal, transform into animals, and wield elemental destruction.",
			Description:                "Druids belong to ancient orders that call on the forces of nature. Harnessing the magic of animals, plants, and the four elements, Druids heal, transform into animals, and wield elemental destruction.\n\nRevering nature above all, individual Druids gain their magic from nature, a nature deity, or both, and they typically unite with other Druids to perform rites that mark the passage of the seasons and other natural cycles.\n\nDruids are concerned with the delicate ecological balance that sustains plant and animal life and with the need for people to live in harmony with nature. Druids often guard sacred sites or watch over regions of unspoiled nature, but when a significant danger arises, Druids take a more active role as adventurers who combat the threat.",
			PrimaryAbility:             "Wisdom",
			HitPointDieValue:           8,
			SkillProficiencyMaxOptions: 2,
		},
		{
			ID:                         7,
			Name:                       "Fighter",
			Tagline:                    "A Master of All Arms and Armour",
			CardDescription:            "Fighters all share an unparalleled prowess with weapons and armour and are well acquainted with death, both meting it out and defying it.",
			Description:                "Fighters rule many battlefields. Questing knights, royal champions, elite soldiers, and hardened mercenaries—as Fighters, they all share an unparalleled prowess with weapons and armor. And they are well acquainted with death, both meting it out and defying it.\n\nFighters master various weapon techniques, and a well-equipped Fighter always has the right tool at hand for any combat situation. Likewise, a Fighter is adept with every form of armor. Beyond that basic degree of familiarity, each Fighter specializes in certain styles of combat. Some concentrate on archery, some on fighting with two weapons at once, and some on augmenting their martial skills with magic. This combination of broad ability and extensive specialization makes Fighters superior combatants.",
			PrimaryAbility:             "Strength or Dexterity",
			HitPointDieValue:           10,
			SkillProficiencyMaxOptions: 2,
		},
		{
			ID:                         8,
			Name:                       "Monk",
			Tagline:                    "A Martial Artist of Supernatural Focus",
			CardDescription:            "Monks focus their internal reservoirs of power to create extraordinary, even supernatural effects.",
			Description:                "Monks use rigorous combat training and mental discipline to align themselves with the multiverse and focus their internal reservoirs of power. Different Monks conceptualize this power in various ways: as breath, energy, life force, essence, or self, for example. Whether channeled as a striking display of martial prowess or as a subtler manifestation of defense and speed, this power infuses all that a Monk does.\n\nMonks focus their internal power to create extraordinary, even supernatural, effects. They channel uncanny speed and strength into their attacks, with or without the use of weapons. In a Monk’s hands, even the most basic weapons can become sophisticated implements of combat mastery.\n\nMany Monks find that a structured life of ascetic withdrawal helps them cultivate the physical and mental focus they need to harness their power. Other Monks believe that immersing themselves in the vibrant confusion of life helps to fuel their determination and discipline.\n\nMonks generally view adventures as tests of their physical and mental development. They are driven by a desire to accomplish a greater mission than merely slaying monsters and plundering treasure; they strive to turn themselves into living weapons.",
			PrimaryAbility:             "Dexterity & Wisdom",
			HitPointDieValue:           8,
			SkillProficiencyMaxOptions: 2,
		},
		{
			ID:                         9,
			Name:                       "Paladin",
			Tagline:                    "A Devout Warrior of Sacred Oaths",
			CardDescription:            "Paladins live on the front lines of the cosmic struggle, united by their oaths against the forces of annihilation.",
			Description:                "Paladins are united by their oaths to stand against the forces of annihilation and corruption. Whether sworn before a god’s altar, in a sacred glade before nature spirits, or in a moment of desperation and grief with the dead as the only witnesses, a Paladin’s oath is a powerful bond. It is a source of power that turns a devout warrior into a blessed champion.\n\nPaladins train to learn the skills of combat, mastering a variety of weapons and armor. Even so, their martial skills are secondary to the magical power they wield: power to heal the injured, smite their foes, and protect the helpless and those who fight at their side.\n\nAlmost by definition, the life of a Paladin is an adventuring life, for every Paladin lives on the front lines of the cosmic struggle against annihilation. Fighters are rare enough among the ranks of a world’s armies, but even fewer people can claim the calling of a Paladin. When they do receive the call, these blessed folk turn from their former occupations and take up arms and magic.",
			PrimaryAbility:             "Strength & Charisma",
			HitPointDieValue:           10,
			SkillProficiencyMaxOptions: 2,
		},
		{
			ID:                         10,
			Name:                       "Ranger",
			Tagline:                    "A Wandering Warrior Imbued with Primal Magic",
			CardDescription:            "Rangers are honed with deadly focus and harness primal powers to protect the world from the ravages of monsters and tyrants.",
			Description:                "Far from bustling cities, amid the trees of trackless forests and across wide plains, Rangers keep their unending watch in the wilderness. Rangers learn to track their quarry as a predator does, moving stealthily through the wilds and hiding themselves in brush and rubble.\n\nThanks to their connection with nature, Rangers can also cast spells that harness primal powers of the wilderness. A Ranger’s talents and magic are honed with deadly focus to protect the world from the ravages of monsters and tyrants.",
			PrimaryAbility:             "Dexterity & Wisdom",
			HitPointDieValue:           10,
			SkillProficiencyMaxOptions: 3,
		},
		{
			ID:                         11,
			Name:                       "Rogue",
			Tagline:                    "A Dexterous Expert in Stealth and Subterfuge",
			CardDescription:            "Rogues have a knack for finding the solution to just about any problem, prioritising subtle strikes over brute strength.",
			Description:                "Rogues rely on cunning, stealth, and their foes’ vulnerabilities to get the upper hand in any situation. They have a knack for finding the solution to just about any problem. A few even learn magical tricks to supplement their other abilities. Many Rogues focus on stealth and deception, while others refine skills that help them in a dungeon environment, such as climbing, finding and disarming traps, and opening locks.\n\nIn combat, Rogues prioritize subtle strikes over brute strength. They would rather make one precise strike than wear an opponent down with a barrage of blows.\n\nSome Rogues began their careers as criminals, while others used their cunning to fight crime. Whatever a Rogue’s relation to the law, no common criminal or officer of the law can match the subtle brilliance of the greatest Rogues.",
			PrimaryAbility:             "Dexterity",
			HitPointDieValue:           8,
			SkillProficiencyMaxOptions: 4,
		},
		{
			ID:                         12,
			Name:                       "Sorcerer",
			Tagline:                    "A Dazzling Mage Filled with Innate Magic",
			CardDescription:            "Sorcerers harness and channel the raw, roiling power of innate magic that is stamped into their very being.",
			Description:                "Sorcerers wield innate magic that is stamped into their being. Some Sorcerers can’t name the origin of their power, while others trace it to strange events in their personal or family history. The blessing of a dragon or a dryad at a baby’s birth or the strike of lightning from a clear sky might spark a Sorcerer’s gift. So too might the gift of a deity, exposure to the strange magic of another plane of existence, or a glimpse into the inner workings of reality. Whatever the origin, the result is an indelible mark on the Sorcerer, a churning magic that can be passed down through generations.\n\nSorcerers don’t learn magic; the raw, roiling power of magic is part of them. The essential art of a Sorcerer is learning to harness and channel that innate magic, allowing the Sorcerer to discover new and staggering ways to unleash their power. As Sorcerers master their innate magic, they grow more attuned to its origin, developing distinct powers that reflect its source.\n\nSorcerers are rare. Some family lines produce exactly one Sorcerer in every generation, but most of the time, the talents of sorcery appear as a fluke. People who have this magical power soon discover that it doesn’t like to stay quiet. A Sorcerer’s magic wants to be wielded.",
			PrimaryAbility:             "Charisma",
			HitPointDieValue:           6,
			SkillProficiencyMaxOptions: 2,
		},
		{
			ID:                         13,
			Name:                       "Warlock",
			Tagline:                    "An Occultist Empowered by Otherworldly Pacts",
			CardDescription:            "Warlocks quest for knowledge that lies hidden in the fabric of the multiverse, piecing together arcane secrets to bolster their own power.",
			Description:                "Warlocks quest for knowledge that lies hidden in the fabric of the multiverse. They often begin their search for magical power by delving into tomes of forbidden lore, dabbling in invocations meant to attract the power of extraplanar beings, or seeking places of power where the influence of these beings can be felt. In no time, each Warlock is drawn into a binding pact with a powerful patron. Drawing on the ancient knowledge of beings such as angels, archfey, demons, devils, hags, and alien entities of the Far Realm, Warlocks piece together arcane secrets to bolster their own power.\n\nWarlocks view their patrons as resources, as means to the end of achieving magical power. Some Warlocks respect, revere, or even love their patrons; some serve their patrons grudgingly; and some seek to undermine their patrons even as they wield the power their patrons have given them.\n\nOnce a pact is made, a Warlock’s thirst for knowledge and power can’t be slaked with mere study. Most Warlocks spend their days pursuing greater power and deeper knowledge, which typically means some kind of adventure.",
			PrimaryAbility:             "Charisma",
			HitPointDieValue:           8,
			SkillProficiencyMaxOptions: 2,
		},
		{
			ID:                         14,
			Name:                       "Wizard",
			Tagline:                    "A Scholarly Magic-User of Arcane Power",
			CardDescription:            "Wizards cast spells of explosive fire, arcing lightning, subtle deception, and spectacular transformations.",
			Description:                "Wizards are defined by their exhaustive study of magic’s inner workings. They cast spells of explosive fire, arcing lightning, subtle deception, and spectacular transformations. Their magic conjures monsters from other planes of existence, glimpses the future, or forms protective barriers. Their mightiest spells change one substance into another, call meteors from the sky, or open portals to other worlds.\n\nMost Wizards share a scholarly approach to magic. They examine the theoretical underpinnings of magic, particularly the categorization of spells into schools of magic. Renowned Wizards such as Bigby, Tasha, Mordenkainen, and Yolande have built on their studies to invent iconic spells now used across the multiverse.\n\nThe closest a Wizard is likely to come to an ordinary life is working as a sage or lecturer. Other Wizards sell their services as advisers, serve in military forces, or pursue lives of crime or domination.\n\nBut the lure of knowledge calls even the most unadventurous Wizards from the safety of their libraries and laboratories and into crumbling ruins and lost cities. Most Wizards believe that their counterparts in ancient civilizations knew secrets of magic that have been lost to the ages, and discovering those secrets could unlock the path to a power greater than any magic available in the present age.",
			PrimaryAbility:             "Intelligence",
			HitPointDieValue:           6,
			SkillProficiencyMaxOptions: 2,
		},
	}

	for _, class := range classes {
		err := s.DB.Where("id = ?", class.ID).FirstOrCreate(&class).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating class with id `%v` in seeder: %v", class.ID, err.Error()))
		}
	}
}
