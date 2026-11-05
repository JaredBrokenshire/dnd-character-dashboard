package handlers

import (
	"dnd-character-dashboard/api"
	"dnd-character-dashboard/api/responses"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ClassHandler struct {
	server *api.Server
}

func NewClassHandler(server *api.Server) *ClassHandler {
	return &ClassHandler{server: server}
}

// List godoc
// @Summary List classes
// @Description List classes (paginated)
// @ID classes-list
// @Tags Class Actions
// @Accept json
// @Produce json
// @Param search query string false "Search classes by name"
// @Param page query int false "The page number"
// @Param page_size query int false "The numbers of items to return. Max 100"
// @Success 200 {object} responses.PaginatedClassResponse
// @Router /classes [get]
func (h *ClassHandler) List(c echo.Context) error {
	var scopes []func(db *gorm.DB) *gorm.DB

	search := c.QueryParam("search")
	if search != "" {
		searchTerm := fmt.Sprintf("%%%s%%", search)
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("name LIKE ?", searchTerm)
		})
	}

	classes, page, pageSize := h.server.Dependencies.GetClassService().List(c, scopes)
	count := h.server.Dependencies.GetClassService().Count(scopes)

	res := responses.NewPaginatedClassResponse(classes, count, page, pageSize)
	return responses.Response(c, http.StatusOK, res)
}
