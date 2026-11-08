package responses

import "dnd-character-dashboard/db/models"

type SlimClassResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`

	Tagline          string         `json:"tagline"`
	CardDescription  string         `json:"card_description"`
	PrimaryAbility   string         `json:"primary_ability"`
	HitPointDieValue int            `json:"hit_point_die_value"`
	Save             models.Ability `json:"save"`
	SecondarySave    models.Ability `json:"secondary_save"`

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
		Save:             class.Save,
		SecondarySave:    class.SecondarySave,
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
