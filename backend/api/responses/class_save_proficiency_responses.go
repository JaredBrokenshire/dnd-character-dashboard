package responses

import (
	"dnd/db/models"
	"dnd/db/models/types"
)

type ClassSaveProficiencyResponse struct {
	Save types.Ability `json:"save"`
}

func NewClassSaveProficiencyResponse(csp *models.ClassSaveProficiency) *ClassSaveProficiencyResponse {
	return &ClassSaveProficiencyResponse{
		Save: csp.Ability,
	}
}

func NewClassSaveProficiencyResponses(csps []*models.ClassSaveProficiency) []ClassSaveProficiencyResponse {
	var res []ClassSaveProficiencyResponse
	for _, csp := range csps {
		res = append(res, *NewClassSaveProficiencyResponse(csp))
	}

	return res
}
