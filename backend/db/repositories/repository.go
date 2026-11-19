package repositories

import "github.com/jinzhu/gorm"

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) Create(object interface{}) error {
	return r.DB.Create(object).Error
}

func (r *Repository) Update(object interface{}) error {
	return r.DB.Save(object).Error
}

func (r *Repository) Delete(object interface{}) error {
	return r.DB.Delete(object).Error
}

func (r *Repository) Restore(object interface{}) error {
	return r.DB.Unscoped().Model(object).Update("deleted_at", nil).Error
}
