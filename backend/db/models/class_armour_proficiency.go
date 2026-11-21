package models

import "dnd/db/models/types"

type ClassArmourProficiency struct {
	ID      uint
	ClassID uint
	Armour  types.ArmourType
}
