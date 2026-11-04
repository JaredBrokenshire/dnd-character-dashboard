package factories

import (
	"dnd-character-dashboard/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
)

func NewSubclass(db *gorm.DB, subclass *models.Subclass) {
	fillSubclassDefaults(subclass)
	err := db.Create(subclass).Error
	if err != nil {
		panic("error creating subclass in factory: " + err.Error())
	}
}

func fillSubclassDefaults(subclass *models.Subclass) {
	if subclass.Name == "" {
		subclass.Name = random.String(16)
	}
	if subclass.ClassID == 0 {
		subclass.ClassID = 1
	}
}
