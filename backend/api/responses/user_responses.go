package responses

import "dnd-character-dashboard/db/models"

type UserResponse struct {
	ID       uint           `json:"id"`
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Roles    []RoleResponse `json:"roles"`
}

type SlimUserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func NewUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Roles:    NewRoleResponses(user.Roles),
	}
}

func NewSlimUserResponse(user *models.User) *SlimUserResponse {
	return &SlimUserResponse{
		ID:       user.ID,
		Username: user.Username,
	}
}
