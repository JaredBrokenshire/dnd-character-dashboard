package responses

import (
	"dnd/db/models"
)

type UserResponse struct {
	*SlimUserResponse

	Roles []RoleResponse `json:"roles"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewUserResponse(user *models.User) *UserResponse {
	res := &UserResponse{
		SlimUserResponse: NewSlimUserResponse(user),
	}

	if len(user.Roles) > 0 {
		res.Roles = NewRoleResponses(user.Roles)
	}

	return res
}

type SlimUserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func NewSlimUserResponse(user *models.User) *SlimUserResponse {
	return &SlimUserResponse{
		ID:       user.ID,
		Username: user.Username,
	}
}
