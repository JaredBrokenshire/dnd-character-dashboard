package factories

import (
	"dnd/db/models"
	"dnd/db/models/types"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
)

func NewClassWeaponProficiency(db *gorm.DB, cwp *models.ClassWeaponProficiency) {
	fillClassWeaponProficiencyDefaults(cwp)
	err := db.Create(cwp).Error
	if err != nil {
		panic("error creating class weapon proficiency in factory: " + err.Error())
	}
}

func fillClassWeaponProficiencyDefaults(cwp *models.ClassWeaponProficiency) {
	if cwp.ClassID == 0 {
		cwp.ClassID = 1
	}
	if cwp.WeaponType == "" {
		cwp.WeaponType = types.WeaponTypeSimple
	}
	if cwp.Condition == nil {
		randomString := random.String(16)
		cwp.Condition = &randomString
	}
}
