package responses

import "dnd-character-dashboard/db/models"

type RaceResponse struct {
	*SlimRaceResponse

	Description string `json:"description"`

	CreatureType string `json:"creature_type"`
	Size         string `json:"size"`
	BaseSpeed    int    `json:"base_speed"`

	Variants []RaceVariantResponse `json:"variants"`

	DetailsImage FileResponse `json:"details_image"`
}

func NewRaceResponse(race *models.Race) *RaceResponse {
	res := &RaceResponse{
		SlimRaceResponse: NewSlimRaceResponse(race),

		Description: race.Description,

		CreatureType: race.CreatureType,
		Size:         race.Size,
		BaseSpeed:    race.BaseSpeed,
	}

	if len(race.Variants) > 0 {
		res.Variants = NewRaceVariantResponses(race.Variants)
	}

	if race.DetailsImage != nil {
		res.DetailsImage = *NewFileResponse(race.DetailsImage)
	}

	return res
}

type SlimRaceResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`

	CardDescription string `json:"card_description"`

	Traits []TraitResponse `json:"traits"`

	CardBackground FileResponse `json:"card_background"`
}

func NewSlimRaceResponse(race *models.Race) *SlimRaceResponse {
	res := &SlimRaceResponse{
		ID:              race.ID,
		Name:            race.Name,
		CardDescription: race.CardDescription,
	}

	if len(race.Traits) > 0 {
		res.Traits = NewTraitResponses(race.Traits)
	}

	if race.CardBackground != nil {
		res.CardBackground = *NewFileResponse(race.CardBackground)
	}

	return res
}

func NewSlimRaceResponses(races []*models.Race) []SlimRaceResponse {
	var res []SlimRaceResponse
	for _, race := range races {
		res = append(res, *NewSlimRaceResponse(race))
	}

	return res
}

type PaginatedRaceResponse struct {
	Data []SlimRaceResponse `json:"data"`
	Meta ResponseMeta       `json:"meta"`
}

func NewPaginatedRaceResponse(races []*models.Race, count, page, pageSize int) PaginatedRaceResponse {
	return PaginatedRaceResponse{
		Data: NewSlimRaceResponses(races),
		Meta: ResponseMeta{
			TotalCount: count,
			Page:       page,
			PageSize:   pageSize,
		},
	}
}
