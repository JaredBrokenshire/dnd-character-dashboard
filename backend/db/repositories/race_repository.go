package repositories

import (
	"dnd/db"
	"dnd/db/models"
	"github.com/jinzhu/gorm"
)

type RaceRepository struct {
	*Repository
}

func NewRaceRepository(db *gorm.DB) *RaceRepository {
	return &RaceRepository{
		Repository: &Repository{DB: db},
	}
}

func (r *RaceRepository) GetRaces(page, pageSize int, scopes db.Scopes) ([]*models.Race, int, int) {
	var races []*models.Race
	pageSize, paginateFunc := db.Paginate(page, pageSize)

	r.DB.
		Preload("Traits").
		Preload("CardBackground").
		Scopes(paginateFunc).
		Scopes(scopes...).
		Order("name ASC").
		Find(&races)

	return races, page, pageSize
}

func (r *RaceRepository) Count(scopes db.Scopes) int {
	var count int64

	r.DB.
		Model(&models.Race{}).
		Scopes(scopes...).
		Count(&count)

	return int(count)
}
