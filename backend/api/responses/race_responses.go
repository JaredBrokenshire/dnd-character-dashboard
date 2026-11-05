package responses

import "dnd-character-dashboard/db/models"

type RaceResponse struct {
	*SlimRaceResponse
}

type SlimRaceResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`

	CardBackground FileResponse `json:"card_background"`

	CardDescription string `json:"card_description"`

	Traits []SlimTraitResponse `json:"traits"`
}

type PaginatedRaceResponse struct {
	Data []SlimRaceResponse `json:"data"`
	Meta ResponseMeta       `json:"meta"`
}

func NewSlimRaceResponse(race *models.Race) *SlimRaceResponse {
	res := &SlimRaceResponse{
		ID:              race.ID,
		Name:            race.Name,
		CardDescription: race.CardDescription,
	}

	if race.CardBackground != nil {
		res.CardBackground = *NewFileResponse(race.CardBackground)
	}
	if len(race.Traits) > 0 {
		res.Traits = NewTraitResponses(race.Traits)
	}

	return res
}

func NewRaceResponses(races []*models.Race) []SlimRaceResponse {
	var res []SlimRaceResponse
	for _, race := range races {
		res = append(res, *NewSlimRaceResponse(race))
	}

	return res
}

func NewPaginatedRaceResponse(races []*models.Race, count, page, pageSize int) *PaginatedRaceResponse {
	return &PaginatedRaceResponse{
		Data: NewRaceResponses(races),
		Meta: ResponseMeta{
			TotalCount: count,
			Page:       page,
			PageSize:   pageSize,
		},
	}
}
