package responses

import "dnd/db/models"

type SubclassResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewSubclassResponse(subclass *models.Subclass) *SubclassResponse {
	return &SubclassResponse{
		Name:        subclass.Name,
		Description: subclass.Description,
	}
}

func NewSubclassResponses(subclasses []*models.Subclass) []SubclassResponse {
	subclassResponses := make([]SubclassResponse, len(subclasses))
	for i, subclass := range subclasses {
		subclassResponses[i] = *NewSubclassResponse(subclass)
	}
	
	return subclassResponses
}
