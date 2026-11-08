package responses

import "dnd-character-dashboard/db/models"

type UserResponse struct {
	*SlimUserResponse
	Email string `json:"email"`

	Roles []RoleResponse `json:"roles"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type SlimUserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func NewUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		SlimUserResponse: NewSlimUserResponse(user),
		Email:            user.Email,
		Roles:            NewRoleResponses(user.Roles),
	}
}

func NewSlimUserResponse(user *models.User) *SlimUserResponse {
	return &SlimUserResponse{
		ID:       user.ID,
		Username: user.Username,
	}
}
