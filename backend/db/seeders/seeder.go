package seeders

import (
	"dnd-character-dashboard/pkg/file_storage"
	"github.com/jinzhu/gorm"
)

type Seeder struct {
	DB        *gorm.DB
	FileStore file_storage.Store
}

func NewSeeder(db *gorm.DB, fileStore file_storage.Store) *Seeder {
	return &Seeder{
		DB:        db,
		FileStore: fileStore,
	}
}

func (s *Seeder) Seed() {
	s.SetRoles()
	s.SetUsers()

	// Classes
	s.SetClasses()
	s.SetClassLogos()
	s.SetClassCardBackgrounds()

	// Races
	s.SetTraits()
	s.SetRaces()
	s.SetRaceCardBackgrounds()
	s.SetRaceDetailsImages()
	s.SetRaceVariants()
}
