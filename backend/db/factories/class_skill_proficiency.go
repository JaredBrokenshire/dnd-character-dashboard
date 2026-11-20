package factories

import (
	"dnd/db/models"
	"dnd/db/models/types"
	"github.com/jinzhu/gorm"
)

func NewClassSkillProficiency(cb *gorm.DB, csp *models.ClassSkillProficiency) {
	fillClassSkillProficiencyDefaults(csp)
	err := cb.Create(csp).Error
	if err != nil {
		panic("error creating class skill proficiency in factory: " + err.Error())
	}
}

func fillClassSkillProficiencyDefaults(csp *models.ClassSkillProficiency) {
	if csp.ClassID == 0 {
		csp.ClassID = 1
	}
	if csp.Skill == "" {
		csp.Skill = types.SkillAcrobatics
	}
}
