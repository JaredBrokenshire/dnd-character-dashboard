package seeders

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/pkg/logging"
	"fmt"
)

type characterSeed struct {
	character   models.Character
	classIDs    []uint
	subclassIDs []uint
}

func (s *Seeder) SetCharacters() {
	seeds := []characterSeed{
		{
			character: models.Character{
				ID:     1,
				UserID: 1,
				Name:   "PeeWee McAnkle-Biter",
				Level:  5,
				RaceID: 3,
			},
			classIDs:    []uint{2},
			subclassIDs: []uint{4},
		},
		{
			character: models.Character{
				ID:         2,
				UserID:     1,
				Name:       "Ira Veomakute",
				Level:      3,
				RaceID:     6,
				CampaignID: 2,
			},
			classIDs:    []uint{4},
			subclassIDs: []uint{3},
		},
		{
			character: models.Character{
				ID:         3,
				UserID:     1,
				Name:       "Kael Drakeshield",
				Level:      7,
				RaceID:     2,
				CampaignID: 1,
			},
			classIDs:    []uint{7},
			subclassIDs: []uint{2},
		},
		{
			character: models.Character{
				ID:     4,
				UserID: 1,
				Name:   "Axel Claystride",
				Level:  5,
				RaceID: 14,
			},
			classIDs:    []uint{3},
			subclassIDs: []uint{5},
		},
		{
			character: models.Character{
				ID:         5,
				UserID:     1,
				Name:       "Faelan Haversham",
				Level:      10,
				RaceID:     7,
				CampaignID: 1,
			},
			classIDs:    []uint{4},
			subclassIDs: []uint{1},
		},
		{
			character: models.Character{
				ID:         6,
				UserID:     3,
				Name:       "Blade",
				Level:      10,
				RaceID:     17,
				CampaignID: 1,
			},
			classIDs:    []uint{9},
			subclassIDs: []uint{6},
		},
		{
			character: models.Character{
				ID:         7,
				UserID:     3,
				Name:       "Nikoline Nightstriker",
				Level:      7,
				RaceID:     19,
				CampaignID: 1,
			},
			classIDs:    []uint{10, 11},
			subclassIDs: []uint{11},
		},
		{
			character: models.Character{
				ID:         8,
				UserID:     4,
				Name:       "Idris",
				Level:      7,
				RaceID:     8,
				CampaignID: 1,
			},
			classIDs:    []uint{1},
			subclassIDs: []uint{7},
		},
		{
			character: models.Character{
				ID:         9,
				UserID:     4,
				Name:       "Meddw",
				Level:      10,
				RaceID:     3,
				CampaignID: 1,
			},
			classIDs:    []uint{8},
			subclassIDs: []uint{10},
		},
		{
			character: models.Character{
				ID:         10,
				UserID:     5,
				Name:       "Karma",
				Level:      7,
				RaceID:     10,
				CampaignID: 1,
			},
			classIDs:    []uint{12},
			subclassIDs: []uint{8},
		},
		{
			character: models.Character{
				ID:         11,
				UserID:     6,
				Name:       "Lycia",
				Level:      7,
				RaceID:     18,
				CampaignID: 1,
			},
			classIDs:    []uint{5},
			subclassIDs: []uint{9},
		},
		{
			character: models.Character{
				ID:         12,
				UserID:     7,
				Name:       "Vashta",
				Level:      7,
				RaceID:     12,
				CampaignID: 1,
			},
			classIDs:    []uint{2},
			subclassIDs: []uint{12},
		},
	}

	for _, seed := range seeds {
		existing := models.Character{}
		if err := s.DB.Where("id = ?", seed.character.ID).First(&existing).Error; err == nil {
			// Already exists — skip to keep seeder idempotent.
			continue
		}

		if err := s.DB.Create(&seed.character).Error; err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating character with id `%v` in seeder: %v", seed.character.ID, err.Error()))
			continue
		}

		// Build ID-only slices for the join tables. Using Association.Append
		classes := make([]models.Class, len(seed.classIDs))
		for i, id := range seed.classIDs {
			classes[i] = models.Class{ID: id}
		}

		subclasses := make([]models.Subclass, len(seed.subclassIDs))
		for i, id := range seed.subclassIDs {
			subclasses[i] = models.Subclass{ID: id}
		}

		if err := s.DB.Model(&seed.character).Association("Classes").Append(classes).Error; err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error assigning classes to character `%v` in seeder: %v", seed.character.ID, err.Error()))
		}

		if err := s.DB.Model(&seed.character).Association("Subclasses").Append(subclasses).Error; err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error assigning subclasses to character `%v` in seeder: %v", seed.character.ID, err.Error()))
		}
	}
}
