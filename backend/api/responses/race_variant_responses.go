package responses

import "dnd-character-dashboard/db/models"

type RaceVariantResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewRaceVariantResponse(raceVariant *models.RaceVariant) *RaceVariantResponse {
	return &RaceVariantResponse{
		ID:          raceVariant.ID,
		Name:        raceVariant.Name,
		Description: raceVariant.Description,
	}
}

func NewRaceVariantResponses(raceVariants []*models.RaceVariant) []RaceVariantResponse {
	var res []RaceVariantResponse
	for _, raceVariant := range raceVariants {
		res = append(res, *NewRaceVariantResponse(raceVariant))
	}

	return res
}
