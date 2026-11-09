package responses

import "dnd-character-dashboard/db/models"

type TraitResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewTraitResponse(trait *models.Trait) *TraitResponse {
	return &TraitResponse{
		ID:          trait.ID,
		Name:        trait.Name,
		Description: trait.Description,
	}
}

func NewTraitResponses(traits []*models.Trait) []TraitResponse {
	var res []TraitResponse
	for _, trait := range traits {
		res = append(res, *NewTraitResponse(trait))
	}

	return res
}
