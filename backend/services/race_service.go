package services

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/db/repositories"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type RaceServiceInterface interface {
	List(c echo.Context, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Race, int, int)
	Count(scopes []func(db *gorm.DB) *gorm.DB) int
}

type RaceService struct {
	raceRepo *repositories.RaceRepository
}

func NewRaceService(repo *repositories.RaceRepository) *RaceService {
	return &RaceService{raceRepo: repo}
}

func (s *RaceService) List(c echo.Context, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Race, int, int) {
	return s.raceRepo.GetRaces(c, scopes)
}

func (s *RaceService) Count(scopes []func(db *gorm.DB) *gorm.DB) int {
	return s.raceRepo.Count(scopes)
}
