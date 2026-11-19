package seeders

import (
	"dnd/pkg/logging"
	"github.com/jinzhu/gorm"
)

type Seeder struct {
	DB *gorm.DB
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{
		DB: db,
	}
}

func (s *Seeder) Seed() {
	logging.Logger.LogInfo().Msg("Seeding Roles")
	s.SetRoles()
	logging.Logger.LogInfo().Msg("Seeded Roles")
	logging.Logger.LogInfo().Msg("Seeding Users")
	s.SetUsers()
	logging.Logger.LogInfo().Msg("Seeded Users")
}
