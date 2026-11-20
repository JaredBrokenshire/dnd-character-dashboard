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
	res := &ClassWeaponProficiencyResponse{
		WeaponType: cwp.WeaponType,
	}

	if cwp.Condition != nil {
		res.Condition = *cwp.Condition
	}

	return res
}

func NewClassWeaponProficiencyResponses(cwps []*models.ClassWeaponProficiency) []ClassWeaponProficiencyResponse {
	var res []ClassWeaponProficiencyResponse
	for _, cwp := range cwps {
		res = append(res, *NewClassWeaponProficiencyResponse(cwp))
	}

	return res
}
