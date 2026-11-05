package services

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/db/repositories"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type ClassServiceInterface interface {
	List(c echo.Context, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Class, int, int)
	Count(scopes []func(db *gorm.DB) *gorm.DB) int
}

type ClassService struct {
	classRepo *repositories.ClassRepository
}

func NewClassService(repo *repositories.ClassRepository) *ClassService {
	return &ClassService{classRepo: repo}
}

func (s *ClassService) List(c echo.Context, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Class, int, int) {
	return s.classRepo.GetClasses(c, scopes)
}

func (s *ClassService) Count(scopes []func(db *gorm.DB) *gorm.DB) int {
	return s.classRepo.Count(scopes)
}
