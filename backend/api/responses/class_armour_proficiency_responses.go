package responses

import (
	"dnd/db/models"
	"dnd/db/models/types"
)

type ClassArmourProficiencyResponse struct {
	ArmourType types.ArmourType `json:"armour_type"`
}

func NewClassArmourProficiencyResponse(classArmourProficiency *models.ClassArmourProficiency) *ClassArmourProficiencyResponse {
	return &ClassArmourProficiencyResponse{
		ArmourType: classArmourProficiency.Armour,
	}
}

func NewClassArmourProficiencyResponses(classArmourProficiencies []*models.ClassArmourProficiency) []ClassArmourProficiencyResponse {
	var res []ClassArmourProficiencyResponse
	for _, classArmourProficiency := range classArmourProficiencies {
		res = append(res, *NewClassArmourProficiencyResponse(classArmourProficiency))
	}

	return res
}
