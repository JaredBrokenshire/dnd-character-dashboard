package services

import (
	"dnd/db"
	"dnd/db/models"
	"dnd/db/repositories"
)

type RaceServiceInterface interface {
	List(page, pageSize int, scopes db.Scopes) ([]*models.Race, int, int)
	Count(scopes db.Scopes) int
}

type RaceService struct {
	raceRepo *repositories.RaceRepository
}

func NewRaceService(raceRepo *repositories.RaceRepository) *RaceService {
	return &RaceService{raceRepo: raceRepo}
}

func (s *RaceService) List(page, pageSize int, scopes db.Scopes) ([]*models.Race, int, int) {
	return s.raceRepo.GetRaces(page, pageSize, scopes)
}

func (s *RaceService) Count(scopes db.Scopes) int {
	return s.raceRepo.Count(scopes)
}
