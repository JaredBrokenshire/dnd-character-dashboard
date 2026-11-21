package models

import "dnd/db/models/types"

type ClassSaveProficiency struct {
	ID      uint
	ClassID uint
	Ability types.Ability
}
