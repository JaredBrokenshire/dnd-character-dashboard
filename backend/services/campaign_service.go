package services

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/db/repositories"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type CampaignServiceInterface interface {
	GetCampaigns(c echo.Context, userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Campaign, int, int)
	Count(userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) int
}

type CampaignService struct {
	campaignRepo *repositories.CampaignRepository
}

func NewCampaignService(repo *repositories.CampaignRepository) *CampaignService {
	return &CampaignService{campaignRepo: repo}
}

func (s *CampaignService) GetCampaigns(c echo.Context, userID interface{}, scopes []func(cb *gorm.DB) *gorm.DB) ([]*models.Campaign, int, int) {
	campaigns, page, pageSize := s.campaignRepo.GetCampaigns(c, userID, scopes)

	return campaigns, page, pageSize
}

func (s *CampaignService) Count(userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) int {
	return s.campaignRepo.Count(userID, scopes)
}
