package responses

import "dnd-character-dashboard/db/models"

type RaceResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewRaceResponse(race *models.Race) *RaceResponse {
	return &RaceResponse{
		ID:   race.ID,
		Name: race.Name,
	}
}
