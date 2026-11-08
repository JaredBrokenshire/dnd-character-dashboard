package repositories

import (
	"dnd-character-dashboard/db"
	"dnd-character-dashboard/db/models"
	"github.com/jinzhu/gorm"
)

type ClassRepository struct {
	*Repository
}

func NewClassRepository(db *gorm.DB) *ClassRepository {
	return &ClassRepository{
		Repository: &Repository{Db: db},
	}
}

func (r *ClassRepository) GetClasses(page, pageSize int, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Class, int, int) {
	var classes []*models.Class
	pageSize, paginateFunc := db.Paginate(page, pageSize)

	r.Db.
		Preload("Logo").
		Preload("CardBackground").
		Scopes(paginateFunc).
		Scopes(scopes...).
		Order("name ASC").
		Find(&classes)

	return classes, page, pageSize
}

func (r *ClassRepository) Count(scopes []func(db *gorm.DB) *gorm.DB) int {
	var count int64

	r.Db.
		Model(&models.Class{}).
		Scopes(scopes...).
		Count(&count)

	return int(count)
}
