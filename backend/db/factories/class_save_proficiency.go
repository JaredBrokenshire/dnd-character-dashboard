package factories

import (
	"dnd/db/models"
	"dnd/db/models/types"
	"github.com/jinzhu/gorm"
	"math/rand"
)

func NewClassSaveProficiency(db *gorm.DB, csp *models.ClassSaveProficiency) {
	fillClassSaveProficiencyDefaults(csp)
	err := db.Create(csp).Error
	if err != nil {
		panic("error creating class save proficiency in factory: " + err.Error())
	}
}

func fillClassSaveProficiencyDefaults(csp *models.ClassSaveProficiency) {
	if csp.ClassID == 0 {
		csp.ClassID = 1
	}
	if csp.Ability == "" {
		abilities := []types.Ability{
			types.AbilityStrength,
			types.AbilityDexterity,
			types.AbilityConstitution,
			types.AbilityIntelligence,
			types.AbilityWisdom,
			types.AbilityCharisma,
		}
		csp.Ability = abilities[rand.Intn(len(abilities))]
	}
}
