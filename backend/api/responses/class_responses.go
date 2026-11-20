package responses

import "dnd/db/models"

type ClassResponse struct {
	*SlimClassResponse

	WeaponProficiencies        []ClassWeaponProficiencyResponse `json:"weapon_proficiencies"`
	SkillProficiencies         []ClassSkillProficiencyResponse  `json:"skill_proficiencies"`
	SkillProficiencyMaxOptions uint                             `json:"skill_proficiency_max_options"`
}

func NewClassResponse(class *models.Class) *ClassResponse {
	res := &ClassResponse{
		SlimClassResponse:          NewSlimClassResponse(class),
		SkillProficiencyMaxOptions: class.SkillProficiencyMaxOptions,
	}

	if len(class.WeaponProficiencies) > 0 {
		res.WeaponProficiencies = NewClassWeaponProficiencyResponses(class.WeaponProficiencies)
	}

	if len(class.SkillProficiencies) > 0 {
		res.SkillProficiencies = NewClassSkillProficiencyResponses(class.SkillProficiencies)
	}

	return res
}

type SlimClassResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`

	Tagline           string                         `json:"tagline"`
	CardDescription   string                         `json:"card_description"`
	PrimaryAbility    string                         `json:"primary_ability"`
	HitPointDieValue  int                            `json:"hit_point_die_value"`
	SaveProficiencies []ClassSaveProficiencyResponse `json:"save_proficiencies"`

	Logo           FileResponse `json:"logo"`
	CardBackground FileResponse `json:"card_background"`
}

func NewSlimClassResponse(class *models.Class) *SlimClassResponse {
	res := &SlimClassResponse{
		ID:   class.ID,
		Name: class.Name,

		Tagline:          class.Tagline,
		CardDescription:  class.CardDescription,
		PrimaryAbility:   class.PrimaryAbility,
		HitPointDieValue: class.HitPointDieValue,
	}

	if len(class.SaveProficiencies) > 0 {
		res.SaveProficiencies = NewClassSaveProficiencyResponses(class.SaveProficiencies)
	}

	if class.Logo != nil {
		res.Logo = *NewFileResponse(class.Logo)
	}

	if class.CardBackground != nil {
		res.CardBackground = *NewFileResponse(class.CardBackground)
	}

	return res
}

func NewClassResponses(classes []*models.Class) []SlimClassResponse {
	var res []SlimClassResponse
	for _, class := range classes {
		res = append(res, *NewSlimClassResponse(class))
	}

	return res
}

type PaginatedClassResponse struct {
	Data []SlimClassResponse `json:"data"`
	Meta ResponseMeta        `json:"meta"`
}

func NewPaginatedClassResponse(classes []*models.Class, count, page, pageSize int) *PaginatedClassResponse {
	return &PaginatedClassResponse{
		Data: NewClassResponses(classes),
		Meta: ResponseMeta{
			TotalCount: count,
			Page:       page,
			PageSize:   pageSize,
		},
	}
}
