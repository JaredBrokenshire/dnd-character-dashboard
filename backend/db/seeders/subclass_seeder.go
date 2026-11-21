package seeders

import (
	"dnd/db/models"
	"dnd/pkg/logging"
	"fmt"
)

func (s *Seeder) SetSubclasses() {
	subclasses := []models.Subclass{
		{
			ID:          1,
			ClassID:     1, // Artificer
			Name:        "Alchemist",
			Description: "An Alchemist is an expert at combining reagents to produce mystical effects. Alchemists use their creations to give life and to leech it away. Alchemy is the oldest of artificer traditions, and its versatility has long been valued during times of war and peace.",
		},
		{
			ID:          2,
			ClassID:     1, // Artificer
			Name:        "Armourer",
			Description: "An artificer who specializes as an Armorer modifies armor to function almost like a second skin. The armor is enhanced to hone the artificer's magic, unleash potent attacks, and generate a formidable defense. The artificer bonds with this armor, becoming one with it even as they experiment with it and refine its magical capabilities.",
		},
		{
			ID:          3,
			ClassID:     1, // Artificer
			Name:        "Artillerist",
			Description: "An Artillerist specializes in using magic to hurl energy, projectiles, and explosions on a battlefield. This destructive power is valued by armies in the wars on many different worlds. And when war passes, some members of this specialization seek to build a more peaceful world by using their powers to fight the resurgence of strife. The world-hopping gnome artificer Vi has been especially vocal about making things right: \"It's about time we fixed things instead of blowing them all to hell.\"",
		},
		{
			ID:          4,
			ClassID:     1, // Artificer
			Name:        "Battle Smith",
			Description: "Armies require protection, and someone has to put things back together if defenses fail. A combination of protector and medic, a Battle Smith is an expert at defending others and repairing both materiel and personnel. To aid in their work, Battle Smiths are accompanied by a steel defender, a protective companion of their own creation. Many soldiers tell stories of nearly dying before being saved by a Battle Smith and a steel defender.\n\nIn the world of Eberron, Battle Smiths played a key role in House Cannith's work on battle constructs and the original warforged, and after the Last War, these artificers led efforts to aid those who were injured in the war's horrific battles.",
		},
		{
			ID:          5,
			ClassID:     1, // Artificer
			Name:        "Archivist",
			Description: "What are the best ways to store vast amounts of information? That question has occupied chroniclers and librarians for centuries. Scrolls and books were the finest information-storing methods for a long time, but the first Archivists imagined something greater: storing information in a real or an artificial mind. What library could match a mental archive capable of accessing information in seconds? As they work towards this goal, Archivists have become masters of storing knowledge and creating artificial intelligences fueled by magic. These techniques have yet to be perfected and institutionalized, and of all artificers, archivists are on the cutting edge of arcane science.",
		},
		{
			ID:          6,
			ClassID:     2, // Barbarian
			Name:        "Path of the Berserker",
			Description: "Channel Rage into Violent Fury\n\nBarbarians who walk the Path of the Berserker direct their Rage primarily toward violence. Their path is one of untrammeled fury, and they thrill in the chaos of battle as they allow their Rage to seize and empower them.",
		},
		{
			ID:          7,
			ClassID:     2, // Barbarian
			Name:        "Path of the Wild Heart",
			Description: "Walk in Community with the Animal World\n\nBarbarians who follow the Path of the Wild Heart view themselves as kin to animals. These Barbarians learn magical means to communicate with animals, and their Rage heightens their connection to animals as it fills them with supernatural might.",
		},
		{
			ID:          8,
			ClassID:     2, // Barbarian
			Name:        "Path of the World Tree",
			Description: "Trace the Roots and Branches of the Multiverse\n\nBarbarians who follow the Path of the World Tree connect with the cosmic tree Yggdrasil through their Rage. This tree grows among the Outer Planes, connecting them to each other and the Material Plane. These Barbarians draw on the tree’s magic for vitality and as a means of dimensional travel.",
		},
		{
			ID:          9,
			ClassID:     2, // Barbarian
			Name:        "Path of the Zealot",
			Description: "Rage in Ecstatic Union with a God\n\nBarbarians who walk the Path of the Zealot receive boons from a god or pantheon. These Barbarians experience their Rage as an ecstatic episode of divine union that infuses them with power. They are often allies to the priests and other followers of their god or pantheon.",
		},
		{
			ID:          10,
			ClassID:     2, // Barbarian
			Name:        "Path of the Fractured",
			Description: "Two Personalities Are Better than One\n\nBarbarians are defined by their rage, channeling it to unleash brief but potent destruction. Few are brave or unwise enough to study esoteric psychological techniques that split their rage from the rest of their psyche, dividing their identity into two parts: ego and id. When their ego is in control, the Fractured—as these Barbarians are known—possess self-control and cunning beyond most other Barbarians. When they allow their id to take control, their countenance turns monstrous and their bodies swell with the power of rage made physically manifest.",
		},
		{
			ID:          11,
			ClassID:     2, // Barbarian
			Name:        "Path of the Primal Spirit",
			Description: "Rage Alongside a Bestial Spirit\n\nMany Barbarians are in tune with the natural world, but few are as mystically intertwined with it as those who walk the Path of the Primal Spirit. These Barbarians forge powerful connections with beasts and nature spirits, inspiring such ethereal beings to manifest and journey on far-flung adventures.\n\nBarbarians who follow this path have a deep reverence for the cycle of the natural world. Such Barbarians are as likely to accept quests and pleas for aid from local wildlife as they are from people. This respect for animals and spirits doesn’t cross into naivety—few understand the delicate balance between the needs of predator and prey better than those who walk the Path of the Primal Spirit.",
		},
		{
			ID:          12,
			ClassID:     2, // Barbarian
			Name:        "Path of the Wrathful Dead",
			Description: "Channel the Rage of the Unquiet Dead\n\nBarbarians who walk the Path of the Wrathful Dead commune with the spirits of the deceased and channel their spectral powers. These Barbarians are often preoccupied with the dead whispering in their ears, urging them to act upon the whims and emotions of the haunting spirits. Some who walk this path learn to master the powers gifted to them; others find they’ve become servants to the unending chorus clamoring for their attention.",
		},
		{
			ID:          13,
			ClassID:     3, // Bard
			Name:        "College of Dance",
			Description: "Move in Harmony with the Cosmos\n\nBards of the College of Dance know that the Words of Creation can’t be contained within speech or song; the words are uttered by the movements of celestial bodies and flow through the motions of the smallest creatures. These Bards practice a way of being in harmony with the whirling cosmos that emphasizes agility, speed, and grace.",
		},
		{
			ID:          14,
			ClassID:     3, // Bard
			Name:        "College of Glamour",
			Description: "Weave Beguiling Fey Magic\n\nThe College of Glamour traces its origins to the beguiling magic of the Feywild. Bards who study this magic weave threads of beauty and terror into their songs and stories, and the mightiest among them can cloak themselves in otherworldly majesty. Their performances stir up wistful longing for forgotten innocence, evoke unconscious memories of long-held fears, and tug at the emotions of even the most hard-hearted listeners.",
		},
		{
			ID:          15,
			ClassID:     3, // Bard
			Name:        "College of Lore",
			Description: "Plumb the Depths of Magical Knowledge\n\nBards of the College of Lore collect spells and secrets from diverse sources, such as scholarly tomes, mystical rites, and peasant tales. The college’s members gather in libraries and universities to share their lore with one another. They also meet at festivals or affairs of state, where they can expose corruption, unravel lies, and poke fun at self-important figures of authority.",
		},
		{
			ID:          16,
			ClassID:     3, // Bard
			Name:        "College of the Moon",
			Description: "Inspire Allies with Primal Tales\n\nThe College of the Moon traces its origins to the ancient druidic circles of the Moonshae Isles, who entrusted the first Bards of this tradition with chronicling the stories of the islands and their people. Bards of this college draw from the isles’ fey magic and the primal power of the moonwells to bolster their allies, protect the natural world, and inspire their bardic works.",
		},
		{
			ID:          17,
			ClassID:     3, // Bard
			Name:        "College of Valour",
			Description: "Sing the Deeds of Ancient Heroes\n\nBards of the College of Valor are daring storytellers whose tales preserve the memory of the great heroes of the past. These Bards sing the deeds of the mighty in vaulted halls or to crowds gathered around great bonfires. They travel to witness great events firsthand and to ensure that the memory of these events doesn’t pass away. With their songs, they inspire new generations to reach the same heights of accomplishment as the heroes of old.",
		},
		{
			ID:          18,
			ClassID:     3, // Bard
			Name:        "College of Adventurers",
			Description: "Be a Master of Classes\n\nBards from the College of Adventurers learn from heroes of old and stories of legend. Bards are jacks of all trades, and, for the College of Adventurers, this statement rings even truer. They combine all the skills of their companions into one, allowing them to be versatile and supportive. While not having a direct goal, many adventurer-studying Bards live to tell tales of other heroes or seek to create their own. Stories of great deeds, of desperate cunning, of magical anomalies, or of godly might fuel these Bards.",
		},
		{
			ID:          19,
			ClassID:     3, // Bard
			Name:        "College of Fools",
			Description: "Laugh at Your Foes’ Terror\n\nIt’s commonly believed the College of Fools is merely a front for a cult of Beleth. Bards in the College would say their relationship with the Arch Daemon of Fear is much more complicated. These Bards choose to laugh at Etharis’s many horrors rather than be dismayed by them—a whimsical, if ghoulish, philosophy. Their delight in other creatures’ terror encourages their obscene antics and cruel jests. One person’s trick is another person’s terror, after all.",
		},
		{
			ID:          20,
			ClassID:     3, // Bard
			Name:        "College of Requiems",
			Description: "Sing the Songs of the Dead\n\nPerforming macabre melodies filled with grief-stricken chords and mournful refrains, the College of Requiems stirs the very bones of the dead. Requiem Bards weave necromantic magic into their funerary songs to control and empower a host of Undead minions.",
		},
		{
			ID:          21,
			ClassID:     4, // Blood Hunter
			Name:        "Order of the Ghostslayer",
			Description: "The Order of the Ghostslayer is the oldest of the blood hunter orders, its members having originally rediscovered the secrets of hemocraft and refined them for combat against the scourge of undeath. Ghostslayers seek out and study the moment of death, obsessing over the mystery of the transition from life, and the unholy power that can cause the dead to rise once more. These zealous blood hunters make it their life’s work to destroy the scourge of undeath wherever it is found, tuning their abilities to engage undead creatures and those who manipulate the necromancy that creates them.",
		},
		{
			ID:          22,
			ClassID:     4, // Blood Hunter
			Name:        "Order of the Lycan",
			Description: "The ancient curse of lycanthropy is feared by nearly all peoples and cultures, passed through blood and seeding a host with the savage strength and hunger for violence of a wicked beast. The Order of the Lycan is a proud group of blood hunters who undergo “the Taming” — the ceremonial infliction of lycanthropy by a senior member of the order, for those who do not already carry the curse before seeking this path. These hunters then use the magic of their blood to harness the power of the monster they harbor, without losing themselves to it. Using intense will and secret blood magic rituals, members of the Order of the Lycan learn to control and unleash their hybrid forms for short periods of time. Enhanced physical prowess, unnatural resilience, and razor-sharp claws make these warriors a terrible foe to any evil that crosses their path. Yet no training is perfect, and without care and complete focus, even the greatest of blood hunters can temporarily lose themselves to their own hunger.",
		},
		{
			ID:          23,
			ClassID:     4, // Blood Hunter
			Name:        "Order of the Mutant",
			Description: "The process of undertaking the Hunter’s Bane ritual is a painful, scarring, and sometimes fatal experience. Those who survive are irrevocably changed — and not always for the better. Over generations of experimentation, a splinter order of blood hunters honed the way in which hemocraft alters the body, using corrupted alchemy and toxic elixirs to alter their blood even further. Over time, they have modified their capabilities in battle, becoming something beyond what they once were. Calling themselves the Order of the Mutant, these blood hunters now specialize in assessing the strengths and weaknesses of their foes, altering their biology to be best prepared for any conflict.",
		},
		{
			ID:          24,
			ClassID:     4, // Blood Hunter
			Name:        "Order of the Profane Soul",
			Description: "Blood hunters belonging to the Order of the Profane Soul have pushed the limits of hemocraft for use against some of the most terrifying creatures corrupting the world. Ancient fiends and cruel magic-users have long counted on their ability to meld into the background and escape those who hunt them, vanishing into noble courts without a trace, or bending the minds of the most stalwart warriors with but a glance. So the blood hunters who founded this order trusted to their resilience as they delved into the same well of corrupting arcane knowledge, making pacts with lesser evils to better combat the greater threats. And though they might have traded a part of themselves for their power, the members of this order know the benefits of that power far outweigh the price.",
		},
		{
			ID:          25,
			ClassID:     5, // Cleric
			Name:        "Knowledge Domain",
			Description: "Unearth Secrets and Master the Mind\n\nThe Knowledge Domain values learning and understanding above all. Clerics who tap into this domain study esoteric lore, collect old tomes, delve into secret places, and examine the processes of the mind. To them, knowledge is more valuable than material wealth, and learning is an act of worship. Libraries, universities, and other educational institutions also draw on the power of the Knowledge Domain.\n\nGods of knowledge vary from masters of arcane magic to patrons of craft and invention. In Faerûn, Clerics of the Knowledge Domain worship deities of learning and ingenuity such as Oghma and Gond, among many others.",
		},
		{
			ID:          26,
			ClassID:     5, // Cleric
			Name:        "Life Domain",
			Description: "Soothe the Hurts of the World\n\nThe Life Domain focuses on the positive energy that helps sustain all life in the multiverse. Clerics who tap into this domain are masters of healing, using that life force to cure many hurts.\n\nExistence itself relies on the positive energy associated with this domain, so a Cleric of almost any religious tradition might choose it. This domain is particularly associated with agricultural deities, gods of healing or endurance, and gods of home and community. Religious orders of healing also seek the magic of this domain.",
		},
		{
			ID:          27,
			ClassID:     5, // Cleric
			Name:        "Light Domain",
			Description: "Bring Light to Banish Darkness\n\nThe Light Domain emphasizes the divine power to bring about blazing fire and revelation. Clerics who wield this power are enlightened souls infused with radiance and the power of their deities’ discerning vision, charged with chasing away lies and burning away darkness.\n\nThe Light Domain is associated with gods of truth, vigilance, beauty, insight, and renewal. Some of these gods are identified with the sun or as charioteers who guide the sun across the sky. Others are sentinels who pierce deception. Some are deities of beauty and artistry who teach that art is a vehicle for the soul’s improvement.",
		},
		{
			ID:          28,
			ClassID:     5, // Cleric
			Name:        "Trickery Domain",
			Description: "Make Mischief and Challenge Authority\n\nThe Trickery Domain offers magic of deception, illusion, and stealth. Clerics who wield this magic are a disruptive force in the world, puncturing pride, mocking tyrants, freeing captives, and flouting hollow traditions. They prefer subterfuge and pranks to direct confrontation.\n\nGods of trickery are mischief-makers and instigators who stand as a constant challenge to the accepted order among both gods and mortals. They embody the forces of change and social upheaval, and they’re patrons of thieves, scoundrels, gamblers, rebels, and liberators. Religious orders that operate in secret, especially those that seek to undermine oppressive governments or hierarchies, also draw on the power of the Trickery Domain.",
		},
		{
			ID:          29,
			ClassID:     5, // Cleric
			Name:        "War Domain",
			Description: "Inspire Valor and Smite Foes\n\nWar has many manifestations. It can make heroes of ordinary people. It can be desperate and horrific, with acts of cruelty and cowardice eclipsing instances of excellence and courage. Clerics who tap into the magic of the War Domain excel in battle, inspiring others to fight the good fight or offering acts of violence as prayers.\n\nGods of the War Domain watch over warriors and reward them for their great deeds. They include champions of honor and chivalry as well as gods of destruction and pillage. Other war gods take a more neutral stance, promoting war in all its manifestations and supporting warriors in any circumstance.",
		},
		{
			ID:          30,
			ClassID:     5, // Cleric
			Name:        "Eldritch Domain",
			Description: "Wield the Power that Killed the Gods\n\nEven the most benevolent divine beings operate outside the bounds of mortal comprehension. These eldritch forces sing a siren song that calls to mortal worshippers through lucid dreams and terrible whispers. The Eldritch Domain empowers the followers of the unknown and distant forces of chaos, divine entities of eldritch oblivion, and dead gods.",
		},
		{
			ID:          31,
			ClassID:     5, // Cleric
			Name:        "Inquisition Domain",
			Description: "Cleanse the Heretics\n\nThe Inquisition Domain reflects the order of the multiverse and the rejection of tainted magic—so far as certain celestial powers see it. Only the divine casters are pure and fit for use.\n\nSince arcane magic is strong enough to challenge the gods, divine beings of this domain, such as the Arch Seraph Empyreus, demand magic-using mortals are kept in check. Most zealots root out all arcanists, while some strike fragile truces when complete removal isn’t feasible.",
		},
		{
			ID:          32,
			ClassID:     5, // Cleric
			Name:        "Purification Domain",
			Description: "Purge Sickness with Suffering\n\nThe Purification Domain believes sickness, mutation, and other forms of bodily corruption are signs of a corresponding spiritual malady. Virtue can’t exist alongside impurity. The wickedness of humanity and other mortal, sentient creatures is a contagion that manifests plagues such as the Weeping Pox. When there’s a risk of such corruption spreading, there’s only one surefire way to contain the rot and atone for the sins that led to this point: Burn it out.",
		},
		{
			ID:          33,
			ClassID:     6, // Druid
			Name:        "Circle of the Land",
			Description: "Celebrate Connection to the Natural World\n\nThe Circle of the Land comprises mystics and sages who safeguard ancient knowledge and rites. These Druids meet within sacred circles of trees or standing stones to whisper primal secrets in Druidic. The circle’s wisest members preside as the chief priests of their communities.",
		},
		{
			ID:          34,
			ClassID:     6, // Druid
			Name:        "Circle of the Moon",
			Description: "Adopt Animal Forms to Guard the Wilds\n\nDruids of the Circle of the Moon draw on lunar magic to transform themselves. Their order gathers under the moon to share news and perform rituals.\n\nChangeable as the moon, a Druid of this circle might prowl as a great cat one night, soar over the treetops as an eagle the next day, and then crash through undergrowth as a bear to drive off a trespassing monster. The wild is in the Druid’s blood.",
		},
		{
			ID:          35,
			ClassID:     6, // Druid
			Name:        "Circle of the Sea",
			Description: "Become One with Tides and Storms\n\nDruids of the Circle of the Sea draw on the tempestuous forces of oceans and storms. Some view themselves as embodiments of nature’s wrath, seeking vengeance against those who despoil nature. Others seek mystical unity with nature by attuning themselves to the ebb and flow of the tides, following the rush of currents and waves and listening to the inscrutable whispers and roars of the winds.",
		},
		{
			ID:          36,
			ClassID:     6, // Druid
			Name:        "Circle of the Stars",
			Description: "Harness Secrets Hidden in Constellations\n\nThe Circle of the Stars has tracked heavenly patterns since time immemorial, discovering secrets hidden amid the constellations. By understanding these secrets, the Druids of this circle seek to harness the powers of the cosmos.",
		},
		{
			ID:          37,
			ClassID:     6, // Druid
			Name:        "Circle of Blood",
			Description: "Wreak Bloody Havoc\n\nThe Circle of Blood is a keeper of the old ways. They remember how ancient Druids performed sacrificial rituals under a blood-red moon to appease the uncaring forces of nature. These Druids trade blood for life in a delicate balance to bolster their allies and destroy their enemies.",
		},
		{
			ID:          38,
			ClassID:     6, // Druid
			Name:        "Circle of Entropy",
			Description: "Become an Instrument of Decay and Ruin\n\nDruids of the Circle of Entropy see doom as a constant of the natural world. Lives end, works erode, and the civilizations of Etharis ultimately come to ruin. The inevitable end of the old brings opportunity for the new. These fatalistic Druids see themselves as agents of ruin, ending those whose time has passed to make space for those whose time has yet to come.",
		},
		{
			ID:          39,
			ClassID:     6, // Druid
			Name:        "Circle of Mutation",
			Description: "Druids of the Circle of Mutation believe that nature should be improved in order to compete with the unnatural world. Their order hides in the darkest corners of swamps and forests, conducting experiments to warp the structure of their domain. They go unheard from for months, emerging from the wilds with twisted vines and mutated creatures at their side. These Druids have earned the scorn of other circles, discredited as those who have lost their way. In the eyes of these Druids, those who wish to preserve nature as it is simply fear what change brings.",
		},
		{
			ID:          40,
			ClassID:     7, // Fighter
			Name:        "Banneret",
			Description: "Rally Fellow Heroes with Inspiring Leadership\n\nBannerets are paragons of valor and leadership who protect the innocent and rally fellow adventurers to the causes of justice and freedom. Many are knights serving in Cormyr, the Silver Marches, Damara, Chessenta, or other lands across Faerûn. They wander the realms as knights errant, taking the fight against evil beyond their kingdom’s borders.\n\nA Banneret relies on judgment, bravery, and fidelity to the code of chivalry to guide them in defeating evildoers. A lone Banneret is a skilled warrior, but when leading a band of allies one of these warriors can transform even a poorly equipped militia into a ferocious war band.",
		},
		{
			ID:          41,
			ClassID:     7, // Fighter
			Name:        "Battle Master",
			Description: "Master Sophisticated Battle Maneuvers\n\nBattle Masters are students of the art of battle, learning martial techniques passed down through generations. The most accomplished Battle Masters are well-rounded figures who combine their carefully honed combat skills with academic study in the fields of history, theory, and the arts.",
		},
		{
			ID:          42,
			ClassID:     7, // Fighter
			Name:        "Champion",
			Description: "Pursue Physical Excellence in Combat\n\nA Champion focuses on the development of martial prowess in a relentless pursuit of victory. Champions combine rigorous training with physical excellence to deal devastating blows, withstand peril, and garner glory. Whether in athletic contests or bloody battle, Champions strive for the crown of the victor.",
		},
		{
			ID:          43,
			ClassID:     7, // Fighter
			Name:        "Eldritch Knight",
			Description: "Support Combat Skills with Arcane Magic\n\nEldritch Knights combine the martial mastery common to all Fighters with a careful study of magic. Their spells both complement and extend their combat skills, providing additional protection to shore up their armor and also allowing them to engage many foes at once with explosive magic.",
		},
		{
			ID:          44,
			ClassID:     7, // Fighter
			Name:        "Psi Warrior",
			Description: "Augment Physical Might with Psionic Power\n\nPsi Warriors awaken the power of their minds to augment their physical might. They harness this psionic power to infuse their weapon strikes, lash out with telekinetic energy, and create barriers of mental force.",
		},
		{
			ID:          45,
			ClassID:     7, // Fighter
			Name:        "Sorrow Knight",
			Description: "Channel Despair and Darkness to Protect Others\n\nSorrow knights suffuse their opponents with grief and misery from the gloomy plane called the Shadowfell. These fighters harness their past despair to forge a connection with others and protect themselves. A sorrow knight’s intense loss—whether the loss of love, faith, or something equally profound—is the weapon they wield against their foes and their motivator to ensure that no one experiences that same pain.",
		},
		{
			ID:          46,
			ClassID:     7, // Fighter
			Name:        "Bulwark Warrior",
			Description: "Those who learn the fighting prowess of the Bulwark Warriors protect their allies from certain death. Bulwark Warriors guard their allies by provoking their enemies. With endurance and strength as their primary skills, these warriors are prepared to take lethal strikes for their compatriots.",
		},
		{
			ID:          47,
			ClassID:     7, // Fighter
			Name:        "Living Crucible",
			Description: "Fighters who become Living Crucibles have discovered an external means of power. While most Fighters train in martial traditions or study the art of war, Living Crucibles hone their craft of alchemy to endure compounds poisonous to others. In exchange for this rigorous physical and mental preparation, these Fighters are able to temporarily push their bodies past their natural limits. Under the influence of their alchemical compounds, these warriors can see in darkness, enhance their speed, inure themselves to magical attacks, and more.",
		},
		{
			ID:          48,
			ClassID:     7, // Fighter
			Name:        "Nightwatcher",
			Description: "Nightwatchers come from humble origins, often toiling as guards or investigators. Undaunted by dangers in the darkness, they stand firm against the growing night. Ever vigilant, they use the darkness to surprise and overwhelm foes much greater than themselves.",
		},
		{
			ID:          49,
			ClassID:     8, // Monk
			Name:        "Warrior of Mercy",
			Description: "Manipulate Forces of Life and Death\n\nWarriors of Mercy manipulate the life force of others. These Monks are wandering physicians, but they bring a swift end to their enemies. They often wear masks, presenting themselves as faceless bringers of life and death.",
		},
		{
			ID:          50,
			ClassID:     8, // Monk
			Name:        "Warrior of Shadow",
			Description: "Harness Shadow Power for Stealth and Subterfuge\n\nWarriors of Shadow practice stealth and subterfuge, harnessing the power of the Shadowfell. They are at home in darkness, able to draw gloom around themselves to hide, leap from shadow to shadow, and take on a wraithlike form.",
		},
		{
			ID:          51,
			ClassID:     8, // Monk
			Name:        "Warrior of the Elements",
			Description: "Wield Strikes and Bursts of Elemental Power\n\nWarriors of the Elements tap into the power of the Elemental Planes. Harnessing their supernatural focus, these Monks momentarily tame the energy of the Elemental Chaos to empower themselves in and out of battle.",
		},
		{
			ID:          52,
			ClassID:     8, // Monk
			Name:        "Warrior of the Open Hand",
			Description: "Master Unarmed Combat Techniques\n\nWarriors of the Open Hand are masters of unarmed combat. They learn techniques to push and trip their opponents and manipulate their own energy to protect themselves from harm.",
		},
		{
			ID:          53,
			ClassID:     8, // Monk
			Name:        "Warrior of Pride",
			Description: "Monks who become Warriors of Pride value themselves above all. These monks focus on mastering their form, ego, and their destructive power. They display monastic aesthetics publicly with the intention of gaining the respect and admiration of “lesser beings.”\n\nThe prideful traditions of this order include adorning themselves with jewelry, hiding their scars, and commanding respect from those who would oppose them.",
		},
		{
			ID:          54,
			ClassID:     8, // Monk
			Name:        "Warrior of Regret",
			Description: "No one is born intent on bringing evil into this world, yet many lives are transformed by the suffering inflicted upon them. Warriors of Regret are haunted by the shadows of evil deeds done.\n\nThrough training and reflection, Warriors of Regret accept their past can’t be erased. Instead, they focus on forging remorse and grief into a weapon against their foes Though their misdeeds follow their every step, Warriors of Regret are not powerless to atone for them.",
		},
		{
			ID:          55,
			ClassID:     8, // Monk
			Name:        "Warrior of the Leaden Crown",
			Description: "In contrast to the esoteric ideals of other orders, Warriors of the Leaden Crown have practical aims: the self-governance of all people. These Monks see history as a series of clashes between groups of powerful beings where Humanoids are at best collateral damage and at worst disposable pawns. To break this cycle of dependence and destruction, Warriors of the Leaden Crown master mental abilities to fight against the otherworldly powers of the multiverse and protect Humanoid sovereignty.\n\nOne component of this plan involves training to do battle with powerful planar foes. The other equally important component is ensuring that Humanoid societies are prepared to overthrow those already dominating them. To that end, these Monks seek political positions that place them in or near decision-making roles where they can influence others to fight back against Arch Seraphs, Arch Daemons, and Primordials.",
		},
		{
			ID:          56,
			ClassID:     9, // Paladin
			Name:        "Oath of Devotion",
			Description: "Uphold the Ideals of Justice and Order\n\nThe Oath of Devotion binds Paladins to the ideals of justice and order. These Paladins meet the archetype of the knight in shining armor. They hold themselves to the highest standards of conduct, and some—for better or worse—hold the rest of the world to the same standards.\n\nMany who swear this oath are devoted to gods of law and good and use their gods’ tenets as the measure of personal devotion. Others hold angels as their ideals and incorporate images of angelic wings into their helmets or coats of arms.",
		},
		{
			ID:          57,
			ClassID:     9, // Paladin
			Name:        "Oath of Glory",
			Description: "Strive for the Heights of Heroism\n\nPaladins who take the Oath of Glory believe they and their companions are destined to achieve glory through deeds of heroism. They train diligently and encourage their companions, so they’re all ready when destiny calls.\n\nThese paladins share the following tenets:\n\n\t• Endeavor to be known by your deeds.\n\t• Face hardships with courage.\n\t• Inspire others to strive for glory.\n",
		},
		{
			ID:          58,
			ClassID:     9, // Paladin
			Name:        "Oath of the Ancients",
			Description: "Preserve Life and Light in the World\n\nThe Oath of the Ancients is as old as the first elves. Paladins who swear this oath cherish the light; they love the beautiful and life-giving things of the world more than any principles of honor, courage, and justice. They often adorn their armor and clothing with images of growing things—leaves, antlers, or flowers—to reflect their commitment to preserving life and light.\n\nThese paladins share the following tenets:\n\n\t• Kindle the light of hope.\n\t• Shelter life.\n\t• Delight in art and laughter.\n",
		},
		{
			ID:          59,
			ClassID:     9, // Paladin
			Name:        "Oath of the Noble Genies",
			Description: "Brandish the Elemental Splendor of Genies\n\nPaladins sworn to the Oath of the Noble Genies revere the forces of the Elemental Planes. Through taking this oath, Paladins draw power from the four different types of genies—dao, masters of earth; djinn, masters of air; efreet, masters of fire; and marids, masters of water. In Faerûn, many Paladins who swear this oath hail from Calimshan, a land teeming with genies.\n\nPaladins who swear this oath often undertake quests that take them all over the Realms and across the multiverse—including the Elemental Planes. These paladins share the following tenets:\n\n\t• Sow the seeds of creation amid the ashes of destruction.\n\t• Lead with splendor and grace.\n\t• Respect the elements, and fear their wrath.\n",
		},
		{
			ID:          60,
			ClassID:     9, // Paladin
			Name:        "Oath of Vengeance",
			Description: "Punish Evildoers at Any Cost\n\nThe Oath of Vengeance is a solemn commitment to punish those who have committed grievously evil acts. When evil armies slaughter helpless villagers, when a tyrant defies the will of the gods, when a thieves’ guild grows too violent, when a dragon rampages through the countryside—at times like these, paladins arise and swear an Oath of Vengeance to set right what has gone wrong.\n\nThese paladins share the following tenets:\n\n\t• Show the wicked no mercy.\n\t• Fight injustice and its causes.\n\t• Aid those harmed by injustice.\n",
		},
		{
			ID:          61,
			ClassID:     9, // Paladin
			Name:        "Oath of Pestilence",
			Description: "Gain Strength Through Suffering\n\nClad in grime-soaked armor and wielding rusting weapons, Oath of Pestilence Paladins spread corruption, disease, and filth. Bound by an oath that infests their bodies with all manner of plagues, these heralds of decay lumber forward with unholy toughness and grim resolve.\n\nThe Paladins share the following tenets:\n\n\t• Spread plague to foster strength in the survivors.\n\t• There is nothing unnatural or immoral about the ending of life.\n\t• Might makes right.\n",
		},
		{
			ID:          62,
			ClassID:     9, // Paladin
			Name:        "Oath of Slaughter",
			Description: "Revel in Unbridled Violence\n\nThe only truth of war is death. The only virtue is to kill. The only victory is to survive. Heroism, honor, and glory are illusions. This is the creed of Tormach, Arch Daemon of Wrath, but this philosophy rings true for all warriors who devote themselves, body and soul, to the thrill of combat and the delight of bloodshed. Such champions find themselves at the forefront of armies, using the most violent methods at their disposal to bring battles to a swift close … or drawing out the war to sate their own appetites.\n\nPaladins who take the Oath of Slaughter share the following tenets:\n\n\t• Violence is the end, not the means.\n\t• Break the body to break the spirit.\n\t• Revel in bloodshed, whether friend or foe.\n",
		},
		{
			ID:          63,
			ClassID:     9, // Paladin
			Name:        "Oath of Zeal",
			Description: "Eliminate Threats to Ideological Purity\n\nThe Oath of Zeal is taken by paladins consumed by hatred for a specific group or ideology. Zealots, as these Paladins are sometimes called, pursue an inquisition against their enemies at all costs. They abandon compassion and honor as impediments to the more important work of ridding the world of those they deem dangerous or heretical.\n\nThese Paladins share the following tenets:\n\n\t• Reveal heresy and destroy those who practice it.\n\t• Show unwavering conviction and unflinching resolve.\n\t• Perform any sacrifice necessary to stop wickedness.\n",
		},
		{
			ID:          64,
			ClassID:     10, // Ranger
			Name:        "Beast Master",
			Description: "Bond with a Primal Beast\n\nA Beast Master forms a mystical bond with a special animal, drawing on primal magic and a deep connection to the natural world.",
		},
		{
			ID:          65,
			ClassID:     10, // Ranger
			Name:        "Fey Wanderer",
			Description: "Wield Fey Mirth and Fury\n\nA fey mystique surrounds you, thanks to the boon of an archfey or a location in the Feywild that transformed you. However you gained fey magic, you are now a Fey Wanderer. Your joyful laughter brightens the hearts of the downtrodden, and your martial prowess strikes terror in your foes, for great is the mirth of the fey and dreadful is their fury.",
		},
		{
			ID:          66,
			ClassID:     10, // Ranger
			Name:        "Gloom Stalker",
			Description: "Draw on Shadow Magic to Fight Your Foes\n\nGloom Stalkers are at home in the darkest places, wielding magic drawn from the Shadowfell to combat enemies that lurk in darkness.",
		},
		{
			ID:          67,
			ClassID:     10, // Ranger
			Name:        "Hunter",
			Description: "Protect Nature and People from Destruction\n\nYou stalk prey in the wilds and elsewhere, using your abilities as a Hunter to protect nature and people everywhere from forces that would destroy them.",
		},
		{
			ID:          68,
			ClassID:     10, // Ranger
			Name:        "Winter Walker",
			Description: "Withstand the Horrors of Frigid Wastelands\n\nWinter Walkers hone their craft in the bleak and frozen wilds of places like Icewind Dale. These ruthless, rimed Rangers hunt monsters that haunt arctic wastelands, eventually becoming frigid terrors themselves. Winter Walkers are well versed in the phenomena of Icewind Dale, including the latent magic of fallen Netherese cities, endemic monsters like yetis and crag cats, and the rising threat of Underdark invaders. Due to their cold pragmatism, terrifying magic, and mastery of the region, Winter Walkers are regarded with equal parts respect and fear. Ten-Towns citizens say that Winter Walkers’ frequent exposure to malignant entities gives them their fearsome powers. Many Reghed nomads, on the other hand, believe that nature spirits bestow on Winter Walkers a unique curse.",
		},
		{
			ID:          69,
			ClassID:     10, // Ranger
			Name:        "Green Reaper",
			Description: "Slay with Nature’s Venom\n\nGreen Reapers are assassins who specialize in harnessing the toxic elements of flora and fauna. These Rangers often work as killers for hire utilizing their extensive knowledge to end lives discreetly or with gory panache, depending on the poison used and the client’s wishes. Green Reapers exhibit a morbid curiosity when encountering a toxin they’re unfamiliar with, typically followed by an enthusiastic application of the substance on their next foe.",
		},
		{
			ID:          70,
			ClassID:     10, // Ranger
			Name:        "Primordial Archer",
			Description: "Channel the Wrath of the Wilds\n\nIn the areas of Etharis where nature still holds sway, many cling to the ancient customs of woodland hunters and witches. Marking themselves with wild plants and bone needles, theirs is a tradition of cunning, cruelty, and reprisal against all trespassers. Unlike the Druids, these wood witches twist natural magic into hexes, imbuing their weapons with nature’s fury.",
		},
		{
			ID:          71,
			ClassID:     10, // Ranger
			Name:        "Vermin Lord",
			Description: "Grow Powerful from Strength in Numbers\n\nVermin Lords cultivate loyal hordes of disease-bearing rodents to help them battle greater evils. For many of these Rangers, that greater evil is some kind of societal ill, but a rare few see their fellow Humanoids as an infestation that threatens the natural world. Regardless of their goal, they make their homes in sewers, slums, and other forgotten places where they are free to plot against their enemies and tend to their verminkin.",
		},
		{
			ID:          72,
			ClassID:     11, // Rogue
			Name:        "Arcane Trickster",
			Description: "Enhance Stealth with Arcane Spells\n\nSome Rogues enhance their fine-honed skills of stealth and agility with spells, learning magical tricks to aid them in their trade. Some Arcane Tricksters use their talents as pickpockets and burglars, while others are pranksters.",
		},
		{
			ID:          73,
			ClassID:     11, // Rogue
			Name:        "Assassin",
			Description: "Practice the Grim Art of Death\n\nAn Assassin’s training focuses on using stealth, poison, and disguise to eliminate foes with deadly efficiency. While some Rogues who follow this path are hired killers, spies, or bounty hunters, the capabilities of this subclass are equally useful for adventurers facing a variety of monstrous enemies.",
		},
		{
			ID:          74,
			ClassID:     11, // Rogue
			Name:        "Scion of the Three",
			Description: "Become a Gruesome Agent of Malice\n\nA Scion of the Three draws power from a group of malevolent gods known in Baldur’s Gate as the Dead Three: Bane, a god of tyranny; Bhaal, a god of violence and murder; and Myrkul, a god of death. While some Rogues of this subclass pledge themselves ardently to those three macabre gods, others are thrust on this path by a curse. Either way, a scion’s power manifests as various occult gifts, as well as an uncanny talent for striking and terrifying foes.\n\nScions of the Three are most common in Baldur’s Gate, where the Dead Three lived as mortals before ascending to godhood. Underground cults to Bane, Bhaal, and Myrkul often count Scions of the Three among their most useful agents. Outside Baldur’s Gate, secular thieves’ guilds such as the Shadow Thieves of Amn or the Xanathar Guild in Waterdeep might call on a Scion of the Three to undertake an especially violent contract.",
		},
		{
			ID:          75,
			ClassID:     11, // Rogue
			Name:        "Soulknife",
			Description: "Strike Foes with Psionic Blades\n\nA Soulknife strikes with the mind, cutting through barriers both physical and psychic. These Rogues discover psionic power within themselves and channel it to do their roguish work. As a Soulknife, your psionic abilities might have haunted you since childhood, revealing their full potential only as you experienced the stress of adventure. Or you might have sought out an order of psychic adepts and spent years learning how to manifest your power.",
		},
		{
			ID:          76,
			ClassID:     11, // Rogue
			Name:        "Thief",
			Description: "Hunt for Treasure as a Classic Adventurer\n\nA mix of burglar, treasure hunter, and explorer, you are the epitome of an adventurer. In addition to improving your agility and stealth, you gain abilities useful for delving into ruins and getting maximum benefit from the magic items you find there.",
		},
		{
			ID:          77,
			ClassID:     11, // Rogue
			Name:        "Highway Rider",
			Description: "A Steed Makes the Best Partner in Crime\n\nStalking the backroads, the Highway Rider strikes fear into the heart of every traveler and penny-pinching merchant. They run down their prize astride a swift and loyal steed— and then make a quick getaway.",
		},
		{
			ID:          78,
			ClassID:     11, // Rogue
			Name:        "Misfortune Bringer",
			Description: "Curse Those You’re About to Strike\n\nYou’ve matched your penchant for illicit activities with the ability to mark your enemies for mishap and misfortune. Whether you were imbued with magic by spending time among the Fey or you learned the art of cursing from a long line of hedge Wizards, you are a Misfortune Bringer. Although not universal, many Misfortune Bringers possess heterochromia iridum, eyes of two different colors, and they use only one when glaring at targets they intend to curse.",
		},
		{
			ID:          79,
			ClassID:     11, // Rogue
			Name:        "Sanguine Thief",
			Description: "In dark alleyways across Etharis, deals are made in coin and blood, but a few enterprising Rogues have posed the question: “Why not both?” Sanguine Thieves are assassins who harness the power of Sangromancy to fuel their abilities, finding new and innovative ways to utilize the blood they see all too frequently in their line of work.\n\nSanguine Thieves learn to draw upon the vita of their victims, extending their ability to use blood magic beyond the limits of their own vitality. When cornered, Sanguine Thieves can unleash this power in a hail of crimson darts or fade before a foe’s blade, leaving nothing behind but a red mist.",
		},
		{
			ID:          80,
			ClassID:     12, // Sorcerer
			Name:        "Aberrant Sorcery",
			Description: "Wield Unnatural Psionic Power\n\nAn alien influence has wrapped its tendrils around your mind, giving you psionic power. You can now touch other minds with that power and alter the world around you. Will this power shine from you as a hopeful beacon to others? Or will you be a terror to those who feel the stab of your mind?\n\nPerhaps a psychic wind from the Astral Plane carried psionic energy to you, or you were exposed to the Far Realm’s warping influence. Alternatively, you were implanted with a mind flayer tadpole, but your transformation into a mind flayer never occurred; now the tadpole’s psionic power is yours. However you acquired this power, your mind is aflame with it.",
		},
		{
			ID:          81,
			ClassID:     12, // Sorcerer
			Name:        "Clockwork Sorcery",
			Description: "Channel Cosmic Forces of Order\n\nThe cosmic force of order has suffused you with magic. That power arises from Mechanus or a realm like it—a plane of existence shaped entirely by clockwork efficiency. You or someone from your lineage might have become entangled in the machinations of modrons, the orderly beings who inhabit Mechanus. Perhaps your ancestor even took part in the Great Modron March. Whatever its origin within you, the power of order can seem strange to others, but for you, it’s part of a vast and glorious system.",
		},
		{
			ID:          82,
			ClassID:     12, // Sorcerer
			Name:        "Draconic Sorcery",
			Description: "Breathe the Magic of Dragons\n\nYour innate magic comes from the gift of a dragon. Perhaps an ancient dragon facing death bequeathed some of its magical power to you or your ancestor. You might have absorbed magic from a site infused with dragons’ power. Or perhaps you handled a treasure taken from a dragon’s hoard that was steeped in draconic power. Or you might have a dragon for an ancestor.",
		},
		{
			ID:          83,
			ClassID:     12, // Sorcerer
			Name:        "Spellfire Sorcery",
			Description: "Wield Raw Magic\n\nYour innate power stems from the source of magic itself: the Weave. This connection manifests as a rare ability known as spellfire, and you surge with radiant bursts of this raw magic. Your talent with spellfire allows you to heal allies, sear enemies, and absorb powerful spells.\n\nWielders of spellfire tend to have a penchant for wandering. Many travel between cosmopolitan settlements, such as those along the Sword Coast, and wield their magic in service of the common good. Others realize their own strange powers by roaming equally strange lands, from the magic-blasted wastes of the desert of Anauroch to the godtouched wilds of the Old Empires.\n\nWherever they go in the Realms, spellfire Sorcerers are courted by factions with interests in the arcane arts, such as the Harpers, Cult of the Dragon, and Red Wizards.",
		},
		{
			ID:          84,
			ClassID:     12, // Sorcerer
			Name:        "Wild Magic Sorcery",
			Description: "Unleash Chaotic Magic\n\nYour innate magic stems from the forces of chaos that underlie the order of creation. You or an ancestor might have endured exposure to raw magic, perhaps through a planar portal leading to Limbo or the Elemental Planes. Perhaps you were blessed by a fey being or marked by a demon. Or your magic could be a fluke with no apparent cause. Whatever its source, this magic churns within you, waiting for any outlet.",
		},
		{
			ID:          85,
			ClassID:     12, // Sorcerer
			Name:        "Apocalypse Sorcery",
			Description: "Hasten or Stall the End of the World\n\nApocalypse Sorcerers have seen the shape of things to come. They may have visions of the Great Beast ushering in the end of Etharis or dreams of Gormadraug waking and devouring the land.\n\nThese prophecies besiege Apocalypse Sorcerers, plaguing and empowering in equal measure. Some find relief in putting these revelations into writing, producing texts that read as incoherent but slowly seep into the mind as sensible, even inevitable. When the apocalypse-bringer is freed to travel across Etharis unburdened, undoing the works of gods and men, magic will swell and pool, ripe for the taking. This sign of the end is also their gift.",
		},
		{
			ID:          86,
			ClassID:     12, // Sorcerer
			Name:        "Haunted Sorcery",
			Description: "Commune with the Dead\n\nUnlike other Sorcerers, the circumstances of your birth were mundane and you have no arcane inheritance passed down from strange ancestors. Instead, you gained your sorcerous powers after you survived an experience that left you near death. Since that experience, you’ve had a preternatural sense for danger and a ghostly companion that either can’t, or won’t, leave you alone. Some Haunted, as Sorcerers who share your origin are called, develop cordial relationships with their phantom, while others find their spectral companion to be a relentless nuisance.",
		},
		{
			ID:          87,
			ClassID:     12, // Sorcerer
			Name:        "Wretched Bloodline",
			Description: "Wield Your Curse Like a Weapon\n\nWith every promise whispered to a fairy, every contract signed with a devil, and every pact made with the dead, there’s a chance someone doesn’t keep their end of the bargain. The vengeance of immortal beings lasts much longer than one lifetime, and the lingering magic of these curses can affect the mortal’s descendants.\n\nThese inherited magical afflictions may manifest as a plague, deformity, or aversion to the sun. In such families, children may be born who learn to master the latent magic within their inherited curses, turning their banes into personal boons—these sorcerers are known collectively as the Wretched.",
		},
		{
			ID:          88,
			ClassID:     13, // Warlock
			Name:        "Archfey Patron",
			Description: "Bargain with Whimsical Fey\n\nYour pact draws on the power of the Feywild. When you choose this subclass, you might make a deal with an archfey, such as the Prince of Frost; the Queen of Air and Darkness, ruler of the Gloaming Court; Titania of the Summer Court; or an ancient hag. Or you might call on a spectrum of Fey, weaving a web of favors and debts. Whoever they are, your patron is often inscrutable and whimsical.",
		},
		{
			ID:          89,
			ClassID:     13, // Warlock
			Name:        "Celestial Patron",
			Description: "Call on the Power of the Heavens\n\nYour pact draws on the Upper Planes, the realms of everlasting bliss. You might enter an agreement with an empyrean, a couatl, a sphinx, a unicorn, or another heavenly entity. Or you might call on numerous such beings as you pursue goals aligned with theirs. Your pact allows you to experience a hint of the holy light that illuminates the multiverse.",
		},
		{
			ID:          90,
			ClassID:     13, // Warlock
			Name:        "Fiend Patron",
			Description: "Make a Deal with the Lower Planes\n\nYour pact draws on the Lower Planes, the realms of perdition. You might forge a bargain with a demon lord such as Demogorgon or Orcus; an archdevil such as Asmodeus; or a pit fiend, balor, yugoloth, or night hag that is especially mighty. That patron’s aims are evil—the corruption or destruction of all things, ultimately including you—and your path is defined by the extent to which you strive against those aims.",
		},
		{
			ID:          91,
			ClassID:     13, // Warlock
			Name:        "Great Old One Patron",
			Description: "Unearth Forbidden Lore of Ineffable Beings\n\nWhen you choose this subclass, you might bind yourself to an unspeakable being from the Far Realm or an elder god—a being such as Tharizdun, the Chained God; Zargon, the Returner; Hadar, the Dark Hunger; or Great Cthulhu. Or you might invoke several entities without yoking yourself to one. The motives of these beings are incomprehensible, and the Great Old One might be indifferent to your existence. But the secrets you’ve learned nevertheless allow you to draw strange magic from it.",
		},
		{
			ID:          92,
			ClassID:     13, // Warlock
			Name:        "Parasite Patron",
			Description: "Become One with a Cosmic Being\n\nYour patron is a cosmic parasite, draining the vitality of entire peoples and worlds. Upon forging a pact with such a patron, you become host to one of its nascent offspring. As you grow in power, the boundary between your identity and the parasitic larva within blurs.",
		},
		{
			ID:          93,
			ClassID:     13, // Warlock
			Name:        "The Coven",
			Description: "Curse and Befuddle with Hag Magic\n\nHags dwell in the deep places of Etharis, from the primordial forests of Valika to the putrid streets of Liesech. These creatures, though rarely benevolent, are willing to bargain with mortals in exchange for furthering their own capricious goals.\n\nWarlocks often find themselves with a single Hag as a sponsor. Should this Hag have a coven, it can lead to the peculiar situation of Warlock agents acting against other Hags or their Warlocks within the coven. But such is the politics of the Realm of Faerie.",
		},
		{
			ID:          94,
			ClassID:     13, // Warlock
			Name:        "The First Vampire Patron",
			Description: "Wield the Power of Undeath\n\nYou have made a pact with a powerful vampire cursed by the gods or born as a creature of the night. Hundreds of haunted beings serve this vampire, so why enter a pact with you? Because you can act in the light of day unhindered? To cull weaker vampires?",
		},
		{
			ID:          95,
			ClassID:     14, // Wizard
			Name:        "Abjurer",
			Description: "Shield Companions and Banish Foes\n\nYour study of magic is focused on spells that block, banish, or protect—ending harmful effects, banishing evil influences, and protecting the weak. Abjurers are sought when baleful spirits require exorcism, when locations must be guarded against magical spying, and when portals to other planes of existence must be closed. Adventuring parties value Abjurers for the protection they provide against a variety of hostile magic and other attacks.",
		},
		{
			ID:          96,
			ClassID:     14, // Wizard
			Name:        "Bladesinger",
			Description: "Wield Weapon and Wizardry in Elegant Tandem\n\nBladesingers master a tradition of wizardry that incorporates swordplay and dance. In combat, a Bladesinger uses intricate, elegant maneuvers that fend off harm and allow the Bladesinger to channel magic into devastating attacks and a cunning defense. Many who have observed a Bladesinger at work remember the display as one of the more beautiful experiences in their life— a glorious dance accompanied by a singing blade.\n\nBladesinging is associated with the ancient elven societies that first mastered the art and coined the term. Even today, most Bladesingers still hail from old elven realms, such as Myth Drannor, or from non-elven societies that share land and history with elves, such as the Silver Marches. Wherever they hail from, Bladesingers take their talents all across the Realms to help common people and perform heroic deeds. Most communities greet the arrival of a Bladesinger as a good omen.",
		},
		{
			ID:          97,
			ClassID:     14, // Wizard
			Name:        "Diviner",
			Description: "Learn the Secrets of the Multiverse\n\nThe counsel of a Diviner is sought by those who want a clearer understanding of the past, present, and future. As a Diviner, you strive to part the veils of space, time, and consciousness. You work to master spells of discernment, remote viewing, supernatural knowledge, and foresight.",
		},
		{
			ID:          98,
			ClassID:     14, // Wizard
			Name:        "Evoker",
			Description: "Create Explosive Elemental Effects\n\nYour studies focus on magic that creates powerful elemental effects such as bitter cold, searing flame, rolling thunder, crackling lightning, and burning acid. Some Evokers find employment in military forces, serving as artillery to blast armies from afar. Others use their power to protect others, while some seek their own gain.",
		},
		{
			ID:          99,
			ClassID:     14, // Wizard
			Name:        "Illusionist",
			Description: "Weave Subtle Spells of Deception\n\nYou specialize in magic that dazzles the senses and tricks the mind, and the illusions you craft make the impossible seem real.",
		},
		{
			ID:          100,
			ClassID:     14, // Wizard
			Name:        "Daemonologist",
			Description: "Through theological study, or perhaps an induction into a secretive cabal, Daemonologists have learned to cast divine magic without the need for faith or bargains. Their spellbooks are scribed with prayer fragments and blasphemous hymns, enabling them to siphon scraps of power from both Arch Seraphs and Arch Daemons. From their heretical research, they've come to believe such otherworldly beings are two sides of the same divine kin.\n\nDaemonology is rarely practiced. If the nature of their abilities is discovered, they risk being persecuted not only by superstitious folk but also by celestial and infernal powers.",
		},
		{
			ID:          101,
			ClassID:     14, // Wizard
			Name:        "Plague Doctor",
			Description: "Blending spellcasting with science, you distill your magic into concoctions that harm or heal. Plague Doctors often wear grotesque masks protecting them from their toxic ingredients. Many fearfully regard the mask as a sign of pestilence, making Plague Doctors a source of both hope and trepidation.",
		},
		{
			ID:          102,
			ClassID:     14, // Wizard
			Name:        "Sangromancer",
			Description: "You study an uncommon school of magic known as Sangromancy or blood magic. Despite its dark reputation, there's nothing inherently evil about the practice of Sangromancy, though its demands on its practitioners are gruesome. As a Sangromancer, your chosen magic demands more than knowledge: it demands sacrifice. Other Wizards may view you with skepticism or even outright hostility, but none can deny the potency of your art.",
		},
	}

	for _, subclass := range subclasses {
		err := s.DB.Where("id = ?", subclass.ID).FirstOrCreate(&subclass).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating subclass with id `%v` in seeder: %v", subclass.ID, s.DB.Error))
		}
	}
}
