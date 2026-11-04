package factories

import (
	"dnd-character-dashboard/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
)

func NewCampaign(db *gorm.DB, campaign *models.Campaign) {
	fillCampaignDefaults(campaign)
	err := db.Create(campaign).Error
	if err != nil {
		panic("error creating campaign in factory: " + err.Error())
	}
}

func fillCampaignDefaults(campaign *models.Campaign) {
	if campaign.Name == "" {
		campaign.Name = random.String(16)
	}
}
