package services

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/db/repositories"
	"github.com/jinzhu/gorm"
)

type ClassServiceInterface interface {
	List(page, pageSize int, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Class, int, int)
	Count(scopes []func(db *gorm.DB) *gorm.DB) int
}

type ClassService struct {
	classRepo *repositories.ClassRepository
}

func NewClassService(classRepo *repositories.ClassRepository) *ClassService {
	return &ClassService{classRepo: classRepo}
}

func (s *ClassService) List(page, pageSize int, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Class, int, int) {
	return s.classRepo.GetClasses(page, pageSize, scopes)
}

func (s *ClassService) Count(scopes []func(db *gorm.DB) *gorm.DB) int {
	return s.classRepo.Count(scopes)
}
