package services

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/db/repositories"
	"github.com/jinzhu/gorm"
)

type RaceServiceInterface interface {
	List(page, pageSize int, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Race, int, int)
	Count(scopes []func(db *gorm.DB) *gorm.DB) int
}

type RaceService struct {
	raceRepo *repositories.RaceRepository
}

func NewRaceService(raceRepo *repositories.RaceRepository) *RaceService {
	return &RaceService{raceRepo: raceRepo}
}

func (s *RaceService) List(page, pageSize int, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Race, int, int) {
	return s.raceRepo.GetRaces(page, pageSize, scopes)
}

func (s *RaceService) Count(scopes []func(db *gorm.DB) *gorm.DB) int {
	return s.raceRepo.Count(scopes)
}
