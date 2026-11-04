package seeders

import (
	"bytes"
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/pkg/logging"
	"fmt"
	"github.com/labstack/gommon/random"
	"os"
	"path/filepath"
)

// seedAssetsDir is resolved relative to the binary's working directory.
// The LocalStorage path resolver uses the same "../" walk, so both converge
// on the same root regardless of whether the app is run from backend/ or
// backend/cmd/api/.
const classCardBackgroundSeedDir = "db/seeders/assets/classes/card-backgrounds"

func (s *Seeder) SetClassCardBackgrounds() {
	if s.FileStore == nil {
		logging.Logger.LogWarn().Msg("seeder: skipping SetClassCardBackgrounds — no file store configured")
		return
	}

	var classes []models.Class
	if err := s.DB.Find(&classes).Error; err != nil {
		logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: SetClassCardBackgrounds could not load classes: %v", err.Error()))
		return
	}

	// Resolve the seed asset directory, walking up from cwd if needed.
	assetDir := resolveDir(classCardBackgroundSeedDir)
	if assetDir == "" {
		logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error finding asset dir %q - skipping", classCardBackgroundSeedDir))
		return
	}

	for _, class := range classes {
		// Skip if this class already has a logo.
		if class.CardBackgroundFileID != 0 {
			continue
		}

		filename := fmt.Sprintf("%v.png", class.ID)
		srcPath := filepath.Join(assetDir, filename)

		data, err := os.ReadFile(srcPath)
		if err != nil {
			// No seed image for this class — not an error, just skip.
			logging.Logger.LogInfo().Msg(fmt.Sprintf("Seeder: no seed card background for class %q (%v) - skipping", class.Name, srcPath))
			continue
		}

		storagePath := "/classes/card-backgrounds"
		newFileName := fmt.Sprintf("%v.png", random.String(16))
		if err := s.FileStore.Save(storagePath, bytes.NewReader(data), newFileName); err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error saving card background for class %q: %v", class.Name, err.Error()))
			continue
		}

		file := models.File{
			Name:     newFileName,
			Location: fmt.Sprintf("%v/%v", storagePath, newFileName),
		}
		if err := s.DB.Create(&file).Error; err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error creating file record for class card background %q: %v", class.Name, err.Error()))
			continue
		}

		if err := s.DB.Model(&class).Update("card_background_file_id", file.ID).Error; err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error updating card_background_id for class %q: %v", class.Name, err.Error()))
		}
	}
}
