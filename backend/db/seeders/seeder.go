package seeders

import (
	"dnd/pkg/file_storage"
	"dnd/pkg/logging"
	"github.com/jinzhu/gorm"
	"os"
	"path/filepath"
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
	logging.Logger.LogInfo().Msg("Seeding Class Details Images")
	s.SetClassDetailsImages()
	logging.Logger.LogInfo().Msg("Seeded Class Details Images")
	logging.Logger.LogInfo().Msg("Seeding Class Save Proficiencies")
	s.SetClassSaveProficiencies()
	logging.Logger.LogInfo().Msg("Seeded Class Save Proficiencies")
	logging.Logger.LogInfo().Msg("Seeding Class Skill Proficiencies")
	s.SetClassSkillProficiencies()
	logging.Logger.LogInfo().Msg("Seeded Class Skill Proficiencies")
	logging.Logger.LogInfo().Msg("Seeding Class Weapon Proficiencies")
	s.SetClassWeaponProficiencies()
	logging.Logger.LogInfo().Msg("Seeded Class Weapon Proficiencies")
	logging.Logger.LogInfo().Msg("Seeding Class Armour Proficiencies")
	s.SetClassArmourProficiencies()
	logging.Logger.LogInfo().Msg("Seeded Class Armour Proficiencies")
	logging.Logger.LogInfo().Msg("Seeding Subclasses")
	s.SetSubclasses()
	logging.Logger.LogInfo().Msg("Seeded Subclasses")
}

// resolveDir returns the first path (trying cwd, then ../cwd, then ../../cwd)
// where the directory exists, or empty string if none found.
func resolveDir(path string) string {
	prefixes := []string{"", "../", "../../"}
	for _, prefix := range prefixes {
		candidate := filepath.Join(prefix, path)
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
	}
	return ""
}
