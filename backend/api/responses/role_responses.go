package responses

import "dnd-character-dashboard/db/models"

type RoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewRoleResponse(role *models.Role) *RoleResponse {
	return &RoleResponse{
		ID:   role.ID,
		Name: string(role.Name),
	}
}

func NewRoleResponses(roles []models.Role) []RoleResponse {
	var res []RoleResponse
	for _, role := range roles {
		res = append(res, *NewRoleResponse(&role))
	}

	return res
}
