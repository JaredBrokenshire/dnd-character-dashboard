package responses

import (
	"dnd-character-dashboard/db/models"
	"time"
)

type FileResponse struct {
	ID       uint   `json:"id" example:"1"`
	Name     string `json:"name"`
	Location string `json:"location"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewFileResponse(file *models.File) *FileResponse {
	return &FileResponse{
		ID:       file.ID,
		Name:     file.Name,
		Location: file.Location,

		CreatedAt: file.CreatedAt,
		UpdatedAt: file.UpdatedAt,
	}
}
