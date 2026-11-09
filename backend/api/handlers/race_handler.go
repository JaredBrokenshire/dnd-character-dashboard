package handlers

import (
	"dnd-character-dashboard/api"
	"dnd-character-dashboard/api/responses"
	"dnd-character-dashboard/pkg/logging"
	"dnd-character-dashboard/services"
	"errors"
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

// Get godoc
// @Summary Get race by ID
// @Description Get race by ID
// @ID races-get
// @Tags Race Actions
// @Accept json
// @Produce json
// @Param id path int true "Race ID"
// @Error 404 {object} responses.ErrorResponse
// @Success 200 {object} responses.RaceResponse
// @Router /races/{id} [get]
func (h *RaceHandler) Get(c echo.Context) error {
	id := c.Param("id")

	race, err := h.server.Dependencies.GetRaceService().Get(id)
	switch {
	case errors.Is(err, services.ErrRaceNotFound):
		logging.Logger.LogWarn().Msg("race/get race not found" + err.Error())
		return responses.ErrorResponse(c, http.StatusNotFound, "Race not found")
	case err != nil:
		logging.Logger.LogError().Msg("race/get error: " + err.Error())
		return responses.ErrorResponse(c, http.StatusInternalServerError, "There was an error getting the race")
	}

	res := responses.NewRaceResponse(race)
	return responses.Response(c, http.StatusOK, res)
}
