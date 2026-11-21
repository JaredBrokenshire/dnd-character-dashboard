package models

import "dnd/db/models/types"

type ClassWeaponProficiency struct {
	ID         uint
	ClassID    uint
	WeaponType types.WeaponType
	Condition  string
}
