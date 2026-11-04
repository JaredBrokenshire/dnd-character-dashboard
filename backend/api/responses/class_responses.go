package responses

import "dnd-character-dashboard/db/models"

type ClassResponse struct {
	*SlimClassResponse
}

type SlimClassResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`

	Logo           FileResponse `json:"logo"`
	CardBackground FileResponse `json:"card_background"`

	Tagline          string         `json:"tagline"`
	CardDescription  string         `json:"card_description"`
	PrimaryAbility   string         `json:"primary_ability"`
	HitPointDieValue int            `json:"hit_point_die_value"`
	Save             models.Ability `json:"save"`
	SecondarySave    models.Ability `json:"secondary_save"`
}

type PaginatedClassResponse struct {
	Data []SlimClassResponse `json:"data"`
	Meta ResponseMeta        `json:"meta"`
}

func NewSlimClassResponse(class *models.Class) *SlimClassResponse {
	res := &SlimClassResponse{
		ID:   class.ID,
		Name: class.Name,
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
