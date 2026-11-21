package services

import (
	"dnd/db"
	"dnd/db/models"
	"dnd/db/repositories"
	"errors"
)

type ClassServiceInterface interface {
	List(page, pageSize int, scopes db.Scopes) ([]*models.Class, int, int)
	Count(scopes db.Scopes) int
	Get(id interface{}) (*models.Class, error)
}

var (
	ErrClassNotFound = errors.New("class not found")
)

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

func (s *ClassService) Get(id interface{}) (*models.Class, error) {
	class := s.classRepo.GetByID(id)
	if class.ID == 0 {
		return nil, ErrClassNotFound
	}

	return class, nil
}
