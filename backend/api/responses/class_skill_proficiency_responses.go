package responses

import (
	"dnd/db/models"
	"dnd/db/models/types"
)

type ClassSkillProficiencyResponse struct {
	Skill types.Skill `json:"skill"`
}

func NewClassSkillProficiencyResponse(csp *models.ClassSkillProficiency) *ClassSkillProficiencyResponse {
	return &ClassSkillProficiencyResponse{
		Skill: csp.Skill,
	}
}

func NewClassSkillProficiencyResponses(csps []*models.ClassSkillProficiency) []ClassSkillProficiencyResponse {
	var res []ClassSkillProficiencyResponse
	for _, csp := range csps {
		res = append(res, *NewClassSkillProficiencyResponse(csp))
	}

	return res
}
