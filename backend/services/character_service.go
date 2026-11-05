package services

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/db/repositories"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

var ErrCharacterNotFound = errors.New("character not found")

type CharacterServiceInterface interface {
	List(c echo.Context, userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Character, int, int)
	Count(userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) int
	Get(userID, id interface{}) (*models.Character, error)
}

type CharacterService struct {
	characterRepo *repositories.CharacterRepository
}

func NewCharacterService(repo *repositories.CharacterRepository) *CharacterService {
	return &CharacterService{characterRepo: repo}
}

func (s *CharacterService) List(c echo.Context, userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Character, int, int) {
	characters, page, pageSize := s.characterRepo.GetCharacters(c, userID, scopes)

	return characters, page, pageSize
}

func (s *CharacterService) Count(userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) int {
	return s.characterRepo.Count(userID, scopes)
}

func (s *CharacterService) Get(userID, id interface{}) (*models.Character, error) {
	character := s.characterRepo.GetByID(userID, id)
	if character.ID == 0 {
		return nil, ErrCharacterNotFound
	}

	return character, nil
}
