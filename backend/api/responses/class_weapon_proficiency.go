package responses

import (
	"dnd/db/models"
	"dnd/db/models/types"
)

type ClassWeaponProficiencyResponse struct {
	WeaponType types.WeaponType `json:"weapon_type"`
	Condition  string           `json:"condition"`
}

func NewClassWeaponProficiencyResponse(cwp *models.ClassWeaponProficiency) *ClassWeaponProficiencyResponse {
	return &ClassWeaponProficiencyResponse{
		WeaponType: cwp.WeaponType,
		Condition:  cwp.Condition,
	}
}

func NewClassWeaponProficiencyResponses(cwps []*models.ClassWeaponProficiency) []ClassWeaponProficiencyResponse {
	var res []ClassWeaponProficiencyResponse
	for _, cwp := range cwps {
		res = append(res, *NewClassWeaponProficiencyResponse(cwp))
	}

	return res
}
