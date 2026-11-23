package responses

import "dnd/db/models"

type TraitResponse struct {
	Name string `json:"name"`
}

func NewTraitResponse(trait *models.Trait) *TraitResponse {
	return &TraitResponse{
		Name: trait.Name,
	}
}

func NewTraitResponses(traits []*models.Trait) []TraitResponse {
	var res []TraitResponse
	for _, trait := range traits {
		res = append(res, *NewTraitResponse(trait))
	}

	return res
}
