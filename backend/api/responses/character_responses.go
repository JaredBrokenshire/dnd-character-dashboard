package responses

import "dnd-character-dashboard/db/models"

type CharacterResponse struct {
	*SlimCharacterResponse
}

type SlimCharacterResponse struct {
	ID             uint                   `json:"id"`
	Name           string                 `json:"name"`
	Level          int                    `json:"level"`
	ProfilePicture FileResponse           `json:"profile_picture"`
	Race           RaceResponse           `json:"race"`
	Classes        []SlimClassResponse    `json:"classes"`
	Subclasses     []SlimSubclassResponse `json:"subclasses"`
	Campaign       SlimCampaignResponse   `json:"campaign"`
}

type PaginatedCharacterResponse struct {
	Data []SlimCharacterResponse `json:"data"`
	Meta ResponseMeta            `json:"meta"`
}

func NewCharacterResponse(character *models.Character) *CharacterResponse {
	return &CharacterResponse{
		SlimCharacterResponse: NewSlimCharacterResponse(character),
	}
}

func NewSlimCharacterResponse(character *models.Character) *SlimCharacterResponse {
	res := &SlimCharacterResponse{
		ID:    character.ID,
		Name:  character.Name,
		Level: character.Level,
	}

	if character.ProfilePicture != nil {
		res.ProfilePicture = *NewFileResponse(character.ProfilePicture)
	}

	if character.Race != nil {
		res.Race = *NewRaceResponse(character.Race)
	}
	if len(character.Classes) > 0 {
		res.Classes = NewClassResponses(character.Classes)
	}
	if len(character.Subclasses) > 0 {
		res.Subclasses = NewSubclassResponses(character.Subclasses)
	}
	if character.Campaign != nil {
		res.Campaign = *NewSlimCampaignResponse(character.Campaign)
	}

	return res
}

func NewCharacterResponses(characters []*models.Character) []SlimCharacterResponse {
	var res []SlimCharacterResponse
	for _, character := range characters {
		res = append(res, *NewSlimCharacterResponse(character))
	}
	return res
}

func NewPaginatedCharacterResponse(characters []*models.Character, count, page, pageSize int) *PaginatedCharacterResponse {
	return &PaginatedCharacterResponse{
		Data: NewCharacterResponses(characters),
		Meta: ResponseMeta{
			TotalCount: count,
			Page:       page,
			PageSize:   pageSize,
		},
	}
}
