package repositories

import (
	"dnd-character-dashboard/db"
	"dnd-character-dashboard/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type CampaignRepository struct {
	*Repository
}

func NewCampaignRepository(db *gorm.DB) *CampaignRepository {
	return &CampaignRepository{
		Repository: &Repository{
			Db: db,
		},
	}
}

func (r *CampaignRepository) GetCampaigns(c echo.Context, userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Campaign, int, int) {
	var campaigns []*models.Campaign
	page, pageSize, paginateFunc := db.Paginate(c)

	r.Db.
		Preload("Characters").
		Preload("Characters.ProfilePicture").
		Preload("User").
		Joins("LEFT JOIN characters ON characters.campaign_id = campaigns.id AND characters.user_id = ?", userID).
		Where("campaigns.user_id = ? OR characters.id IS NOT NULL", userID).
		Scopes(paginateFunc).
		Scopes(scopes...).
		Order("campaigns.name ASC").
		Group("campaigns.id").
		Find(&campaigns)

	return campaigns, page, pageSize
}

func (r *CampaignRepository) Count(userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) int {
	var count int64

	r.Db.
		Model(&models.Campaign{}).
		Joins("LEFT JOIN characters ON characters.campaign_id = campaigns.id AND characters.user_id = ?", userID).
		Where("campaigns.user_id = ? OR characters.id IS NOT NULL", userID).
		Scopes(scopes...).
		Group("campaigns.id").
		Count(&count)

	return int(count)
}
