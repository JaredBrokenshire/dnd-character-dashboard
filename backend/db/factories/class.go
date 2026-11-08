package factories

import (
	"dnd-character-dashboard/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"math/rand"
)

func NewClass(db *gorm.DB, class *models.Class) {
	fillClassDefaults(class)
	err := db.Create(class).Error
	if err != nil {
		panic("error creating class in factory: " + err.Error())
	}
}

func fillClassDefaults(class *models.Class) {
	if class.Name == "" {
		class.Name = random.String(16)
	}
	if class.Tagline == "" {
		class.Tagline = random.String(16)
	}
	if class.CardDescription == "" {
		class.CardDescription = random.String(64)
	}
	if class.PrimaryAbility == "" {
		class.PrimaryAbility = random.String(16)
	}
	if class.HitPointDieValue == 0 {
		class.HitPointDieValue = rand.Intn(20)
	}
	if class.Save == "" {
		class.Save = models.Strength
	}
	if class.SecondarySave == "" {
		class.SecondarySave = models.Dexterity
	}
}
