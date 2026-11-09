package repositories

import (
	"dnd-character-dashboard/db"
	"dnd-character-dashboard/db/models"
	"github.com/jinzhu/gorm"
)

type RaceRepository struct {
	*Repository
}

func NewRaceRepository(db *gorm.DB) *RaceRepository {
	return &RaceRepository{
		Repository: &Repository{Db: db},
	}
}

func (r *RaceRepository) GetRaces(page, pageSize int, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Race, int, int) {
	var races []*models.Race
	pageSize, paginateFunc := db.Paginate(page, pageSize)

	r.Db.
		Preload("DetailsImage").
		Preload("CardBackground").
		Preload("Traits").
		Preload("Variants").
		Scopes(paginateFunc).
		Scopes(scopes...).
		Order("name ASC").
		Find(&races)

	return races, page, pageSize
}

func (r *RaceRepository) Count(scopes []func(db *gorm.DB) *gorm.DB) int {
	var count int64
	r.Db.
		Model(&models.Race{}).
		Scopes(scopes...).
		Count(&count)

	return int(count)
}
