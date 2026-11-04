package services

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/db/repositories"
)

type UserServiceInterface interface {
	GetUserById(id interface{}) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
}

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserById(id interface{}) (*models.User, error) {
	user := s.userRepo.GetById(id)
	if user.ID == 0 {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	user := s.userRepo.GetByUsername(username)
	if user.ID == 0 {
		return nil, ErrUserNotFound
	}

	return user, nil
}
