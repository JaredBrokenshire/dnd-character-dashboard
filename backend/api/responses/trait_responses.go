package responses

import "dnd-character-dashboard/db/models"

type TraitResponse struct {
	*SlimTraitResponse
}

type SlimTraitResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewSlimTraitResponse(trait *models.Trait) *SlimTraitResponse {
	return &SlimTraitResponse{
		ID:   trait.ID,
		Name: trait.Name,
	}
}

func NewTraitResponses(traits []*models.Trait) []SlimTraitResponse {
	var res []SlimTraitResponse
	for _, trait := range traits {
		res = append(res, *NewSlimTraitResponse(trait))
	}

	return res
}
