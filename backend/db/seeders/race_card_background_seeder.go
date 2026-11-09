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

const raceCardBackgroundSeedDir = "db/seeders/assets/races/card-backgrounds"

func (s *Seeder) SetRaceCardBackgrounds() {
	if s.FileStore == nil {
		logging.Logger.LogWarn().Msg("seeder: skipping SetRaceCardBackgrounds — no file store configured")
		return
	}

	var races []models.Race
	if err := s.DB.Find(&races).Error; err != nil {
		logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: SetRaceCardBackgrounds could not load races: %v", err.Error()))
		return
	}

	// Resolve the seed asset directory, walking up from cwd if needed.
	assetDir := resolveDir(raceCardBackgroundSeedDir)
	if assetDir == "" {
		logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error finding asset dir %q - skipping", raceCardBackgroundSeedDir))
		return
	}

	for _, race := range races {
		// Skip if this class already has a card background file.
		if race.CardBackgroundFileID != 0 {
			continue
		}

		filename := fmt.Sprintf("%v.png", race.ID)
		srcPath := filepath.Join(assetDir, filename)

		data, err := os.ReadFile(srcPath)
		if err != nil {
			// No seed image for this race — not an error, just skip.
			logging.Logger.LogInfo().Msg(fmt.Sprintf("Seeder: no seed card background for race %q (%v) - skipping", race.Name, srcPath))
			continue
		}

		storagePath := "/races/card-backgrounds"
		newFileName := fmt.Sprintf("%v.png", random.String(16))
		if err := s.FileStore.Save(storagePath, bytes.NewReader(data), newFileName); err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error saving card background for race %q: %v", race.Name, err.Error()))
			continue
		}

		file := models.File{
			Name:     newFileName,
			Location: fmt.Sprintf("%v/%v", storagePath, newFileName),
		}
		if err := s.DB.Create(&file).Error; err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error creating file record for race card background %q: %v", race.Name, err.Error()))
			continue
		}

		if err := s.DB.Model(&race).Update("card_background_file_id", file.ID).Error; err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error updating card_background_file_id for race %q: %v", race.Name, err.Error()))
		}
	}
}
