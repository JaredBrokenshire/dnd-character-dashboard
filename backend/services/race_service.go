package services

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/db/repositories"
	"errors"
	"github.com/jinzhu/gorm"
)

var (
	ErrRaceNotFound = errors.New("race not found")
)

type RaceServiceInterface interface {
	List(page, pageSize int, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Race, int, int)
	Count(scopes []func(db *gorm.DB) *gorm.DB) int
	Get(id interface{}) (*models.Race, error)
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

func (s *RaceService) Get(id interface{}) (*models.Race, error) {
	race := s.raceRepo.GetByID(id)
	if race.ID == 0 {
		return nil, ErrRaceNotFound
	}

	return race, nil
}
