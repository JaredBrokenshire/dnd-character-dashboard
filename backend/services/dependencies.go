package services

import (
	"dnd-character-dashboard/db/repositories"
	"github.com/jinzhu/gorm"
)

type Dependencies struct {
	db    *gorm.DB
	repos *repositories.Repos
}

func NewDependencies(db *gorm.DB, repos *repositories.Repos) *Dependencies {
	return &Dependencies{
		db:    db,
		repos: repos,
	}
}

func (d *Dependencies) GetDB() *gorm.DB {
	return d.db
}
