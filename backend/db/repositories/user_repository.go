package repositories

import (
	"dnd/db/models"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	*Repository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repository: &Repository{DB: db},
	}
}

func (r *UserRepository) GetById(id interface{}) *models.User {
	user := &models.User{}
	r.DB.
		Preload("Roles").
		Where("id = ?", id).
		First(user)

	return user
}

func (r *UserRepository) GetByUsername(username string) *models.User {
	user := &models.User{}
	r.DB.
		Preload("Roles").
		Where("username = ?", username).
		First(&user)

	return user
}
