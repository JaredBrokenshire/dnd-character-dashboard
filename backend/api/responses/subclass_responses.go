package responses

import "dnd-character-dashboard/db/models"

type SlimSubclassResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	ClassID uint   `json:"class_id"`
}

func NewSlimSubclassResponse(subclass *models.Subclass) *SlimSubclassResponse {
	return &SlimSubclassResponse{
		ID:      subclass.ID,
		Name:    subclass.Name,
		ClassID: subclass.ClassID,
	}
}

func NewSubclassResponses(subclasses []models.Subclass) []SlimSubclassResponse {
	var res []SlimSubclassResponse
	for _, subclass := range subclasses {
		res = append(res, *NewSlimSubclassResponse(&subclass))
	}
	return res
}
