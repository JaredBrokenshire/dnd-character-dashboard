package repositories

import (
	"github.com/jinzhu/gorm"
)

type Repos struct {
	DB *gorm.DB
}

func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		DB: db,
	}
}
