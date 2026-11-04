package seeders

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/pkg/logging"
	"fmt"
	"time"
)

func (s *Seeder) SetCampaigns() {
	campaigns := []models.Campaign{
		{
			ID:        1,
			UserID:    8,
			Name:      "Marblethorn",
			CreatedAt: time.Date(2024, time.May, 20, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:        2,
			UserID:    3,
			Name:      "Dungeoneers",
			CreatedAt: time.Date(2023, time.February, 25, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:     3,
			UserID: 1,
			Name:   "Heist of the Sapphire Vault",
		},
	}

	for _, campaign := range campaigns {
		err := s.DB.Where("id = ?", campaign.ID).FirstOrCreate(&campaign).Error
		if err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Error creating campaign with id `%v` in seeder: %v", campaign.ID, err.Error()))
		}
	}
}
