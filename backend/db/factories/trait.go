package factories

import (
	"dnd-character-dashboard/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
)

func NewTrait(db *gorm.DB, trait *models.Trait) {
	fillTraitDefaults(trait)
	err := db.Create(trait).Error
	if err != nil {
		panic("error creating trait in factory: " + err.Error())
	}
}

func fillTraitDefaults(trait *models.Trait) {
	if trait.Name == "" {
		trait.Name = random.String(16)
	}
	if trait.Description == "" {
		trait.Description = random.String(64)
	}
}
