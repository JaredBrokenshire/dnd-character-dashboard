package models

import "dnd/db/models/types"

type ClassSkillProficiency struct {
	ID      uint
	ClassID uint
	Skill   types.Skill
}
