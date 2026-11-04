package repositories

import (
	"dnd-character-dashboard/db"
	"dnd-character-dashboard/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type ClassRepository struct {
	*Repository
}

func NewClassRepository(db *gorm.DB) *ClassRepository {
	return &ClassRepository{
		Repository: &Repository{
			Db: db,
		},
	}
}

func (r *ClassRepository) GetClasses(c echo.Context, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Class, int, int) {
	var classes []*models.Class
	page, pageSize, paginateFunc := db.Paginate(c)

	r.Db.
		Preload("Logo").
		Preload("CardBackground").
		Preload("Subclasses").
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
