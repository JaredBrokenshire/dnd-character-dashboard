package factories

import (
	"dnd/db/models"
	"dnd/db/models/types"
	"github.com/jinzhu/gorm"
)

func NewClassArmourProficiency(db *gorm.DB, classArmourProficiency *models.ClassArmourProficiency) {
	fillClassArmourProficiencyDefaults(classArmourProficiency)
	err := db.Create(classArmourProficiency).Error
	if err != nil {
		panic("error creating class armour proficiency in factory: " + err.Error())
	}
}

func fillClassArmourProficiencyDefaults(classArmourProficiency *models.ClassArmourProficiency) {
	if classArmourProficiency.ClassID == 0 {
		classArmourProficiency.ClassID = 1
	}
	if classArmourProficiency.Armour == "" {
		classArmourProficiency.Armour = types.ArmourTypeLight
	}
}
