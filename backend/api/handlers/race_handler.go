package handlers

import (
	"dnd-character-dashboard/api"
	"dnd-character-dashboard/api/responses"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type RaceHandler struct {
	server *api.Server
}

func NewRaceHandler(server *api.Server) *RaceHandler {
	return &RaceHandler{server: server}
}

// List godoc
// @Summary List races
// @Description List races (paginated)
// @ID races-list
// @Tags Race Actions
// @Accept json
// @Produce json
// @Param search query string false "Search races by name"
// @Param page query int false "The page number"
// @Param page_size query int false "The numbers of items to return. Max 100"
// @Success 200 {object} responses.PaginatedRaceResponse
// @Router /races [get]
func (h *RaceHandler) List(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))

	var scopes []func(db *gorm.DB) *gorm.DB

	search := c.QueryParam("search")
	if search != "" {
		searchTerm := fmt.Sprintf("%%%s%%", search)
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("name LIKE ?", searchTerm)
		})
	}

	races, page, pageSize := h.server.Dependencies.GetRaceService().List(page, pageSize, scopes)
	count := h.server.Dependencies.GetRaceService().Count(scopes)

	res := responses.NewPaginatedRaceResponse(races, count, page, pageSize)
	return responses.Response(c, http.StatusOK, res)
}
