package repositories

import (
	"dnd/db/models"
	"github.com/jinzhu/gorm"
)

type RoleRepository struct {
	*Repository
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		Repository: &Repository{DB: db},
	}
}

func (r *RoleRepository) GetByName(name models.RoleName) *models.Role {
	role := &models.Role{}
	r.DB.
		Where("name = ?", name).
		First(&role)

	return role
}
