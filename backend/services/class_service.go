package services

import (
	"dnd/db"
	"dnd/db/models"
	"dnd/db/repositories"
)

type ClassServiceInterface interface {
	List(page, pageSize int, scopes db.Scopes) ([]*models.Class, int, int)
	Count(scopes db.Scopes) int
}

type ClassService struct {
	classRepo *repositories.ClassRepository
}

func NewClassService(classRepo *repositories.ClassRepository) *ClassService {
	return &ClassService{classRepo: classRepo}
}

func (s *ClassService) List(page, pageSize int, scopes db.Scopes) ([]*models.Class, int, int) {
	return s.classRepo.GetClasses(page, pageSize, scopes)
}

func (s *ClassService) Count(scopes db.Scopes) int {
	return s.classRepo.Count(scopes)
}
