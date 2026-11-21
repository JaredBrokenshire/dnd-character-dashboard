package repositories

import (
	"dnd/db"
	"dnd/db/models"
	"github.com/jinzhu/gorm"
)

type ClassRepository struct {
	*Repository
}

func NewClassRepository(db *gorm.DB) *ClassRepository {
	return &ClassRepository{
		Repository: &Repository{DB: db},
	}
}

func (r *ClassRepository) GetClasses(page, pageSize int, scopes db.Scopes) ([]*models.Class, int, int) {
	var classes []*models.Class
	pageSize, paginateFunc := db.Paginate(page, pageSize)

	r.DB.
		Preload("SaveProficiencies").
		Preload("Logo").
		Preload("CardBackground").
		Scopes(paginateFunc).
		Scopes(scopes...).
		Order("name ASC").
		Find(&classes)

	return classes, page, pageSize
}

func (r *ClassRepository) Count(scopes db.Scopes) int {
	var count int64

	r.DB.
		Model(&models.Class{}).
		Scopes(scopes...).
		Count(&count)

	return int(count)
}

func (r *ClassRepository) GetByID(id interface{}) *models.Class {
	class := &models.Class{}
	r.DB.
		Preload("SaveProficiencies").
		Preload("SkillProficiencies").
		Preload("WeaponProficiencies").
		Preload("ArmourProficiencies").
		Preload("Subclasses").
		Preload("Logo").
		Preload("CardBackground").
		Preload("DetailsImage").
		Where("id = ?", id).
		First(&class)

	return class
}
