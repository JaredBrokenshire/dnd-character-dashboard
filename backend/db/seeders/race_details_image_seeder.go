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

const raceDetailsImageSeedDir = "db/seeders/assets/races/details-images"

func (s *Seeder) SetRaceDetailsImages() {
	if s.FileStore == nil {
		logging.Logger.LogWarn().Msg("seeder: skipping SetRaceDetailsImages — no file store configured")
		return
	}

	var races []models.Race
	if err := s.DB.Find(&races).Error; err != nil {
		logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: SetRaceDetailsImages could not load races: %v", err.Error()))
		return
	}

	// Resolve the seed asset directory, walking up from cwd if needed.
	assetDir := resolveDir(raceDetailsImageSeedDir)
	if assetDir == "" {
		logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error finding asset dir %q - skipping", raceDetailsImageSeedDir))
		return
	}

	for _, race := range races {
		// Skip if this class already has a card background file.
		if race.DetailsImageFileID != 0 {
			continue
		}

		filename := fmt.Sprintf("%v.png", race.ID)
		srcPath := filepath.Join(assetDir, filename)

		data, err := os.ReadFile(srcPath)
		if err != nil {
			// No seed image for this race — not an error, just skip.
			logging.Logger.LogInfo().Msg(fmt.Sprintf("Seeder: no seed details image for race %q (%v) - skipping", race.Name, srcPath))
			continue
		}

		storagePath := "/races/details-images"
		newFileName := fmt.Sprintf("%v.png", random.String(16))
		if err := s.FileStore.Save(storagePath, bytes.NewReader(data), newFileName); err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error saving details image for race %q: %v", race.Name, err.Error()))
			continue
		}

		file := models.File{
			Name:     newFileName,
			Location: fmt.Sprintf("%v/%v", storagePath, newFileName),
		}
		if err := s.DB.Create(&file).Error; err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error creating file record for race details image %q: %v", race.Name, err.Error()))
			continue
		}

		if err := s.DB.Model(&race).Update("details_image_file_id", file.ID).Error; err != nil {
			logging.Logger.LogWarn().Msg(fmt.Sprintf("Seeder: error updating details_image_file_id for race %q: %v", race.Name, err.Error()))
		}
	}
}
