package seeders

import (
	"dnd/pkg/file_storage"
	"dnd/pkg/logging"
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
	logging.Logger.LogInfo().Msg("Seeding Roles")
	s.SetRoles()
	logging.Logger.LogInfo().Msg("Seeded Roles")
	logging.Logger.LogInfo().Msg("Seeding Users")
	s.SetUsers()
	logging.Logger.LogInfo().Msg("Seeded Users")

	// Classes
	logging.Logger.LogInfo().Msg("Seeding Classes")
	s.SetClasses()
	logging.Logger.LogInfo().Msg("Seeded Classes")
	logging.Logger.LogInfo().Msg("Seeding Class Logos")
	s.SetClassLogos()
	logging.Logger.LogInfo().Msg("Seeded Class Logos")
	logging.Logger.LogInfo().Msg("Seeding Class Card Backgrounds")
	s.SetClassCardBackgrounds()
	logging.Logger.LogInfo().Msg("Seeded Class Card Backgrounds")
}
