package services

import (
	"dnd/db/models"
	"dnd/db/repositories"
)

type UserServiceInterface interface {
	Get(id interface{}) (*models.User, error)
}

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Get(id interface{}) (*models.User, error) {
	user := s.userRepo.GetById(id)
	if user.ID == 0 {
		return nil, ErrUserNotFound
	}

	return user, nil
}
