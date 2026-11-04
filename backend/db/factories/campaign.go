package factories

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/pkg/logging"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
)

func NewCampaign(db *gorm.DB, campaign *models.Campaign) {
	fillCampaignDefaults(campaign)
	err := db.Create(campaign).Error
	if err != nil {
		logging.Logger.LogWarn().Msg("Error creating campaign in factory: " + err.Error())
	}
}

func fillCampaignDefaults(campaign *models.Campaign) {
	if campaign.Name == "" {
		campaign.Name = random.String(16)
	}
}
