package factories

import (
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/pkg/logging"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
)

func NewClass(db *gorm.DB, class *models.Class) {
	fillClassDefaults(class)
	err := db.Create(class).Error
	if err != nil {
		logging.Logger.LogWarn().Msg("Error creating class in factory: "+err.Error())
	}
}

func fillClassDefaults(class *models.Class) {
	if class.Name == "" {
		class.Name = random.String(16)
	}
}
