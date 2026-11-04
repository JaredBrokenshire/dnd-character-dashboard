package factories

import (
	"dnd-character-dashboard/db/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
)

func NewFile(db *gorm.DB, file *models.File) {
	setFileDefaults(file)
	err := db.Create(file).Error
	if err != nil {
		panic("error creating file in factory: " + err.Error())
	}
}

func setFileDefaults(file *models.File) {
	if file.Name == "" {
		file.Name = fmt.Sprintf("%v.png", random.String(16))
	}
	if file.Location == "" {
		file.Location = fmt.Sprintf("%v/%v/%v", random.String(16), random.String(16), file.Name)
	}
}
