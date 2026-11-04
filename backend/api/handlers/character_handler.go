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
)

type CharacterHandler struct {
	server *api.Server
}

func NewCharacterHandler(server *api.Server) *CharacterHandler {
	return &CharacterHandler{server: server}
}

// List godoc
// @Summary List characters
// @Description List characters (paginated)
// @ID characters-list
// @Tags Character Actions
// @Accept json
// @Produce json
// @Param search query string false "Search characters by name"
// @Param page query int false "The page number"
// @Param page_size query int false "The numbers of items to return. Max 100"
// @Success 200 {object} responses.PaginatedCharacterResponse
// @Router /characters [get]
func (h *CharacterHandler) List(c echo.Context) error {
	userID := c.Get("userID").(uint)

	var scopes []func(db *gorm.DB) *gorm.DB

	search := c.QueryParam("search")
	if search != "" {
		searchTerm := fmt.Sprintf("%%%s%%", search)
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("name LIKE ?", searchTerm)
		})
	}

	characters, page, pageSize := h.server.Dependencies.GetCharacterService().GetCharacters(c, userID, scopes)
	count := h.server.Dependencies.GetCharacterService().Count(userID, scopes)

	res := responses.NewPaginatedCharacterResponse(characters, count, page, pageSize)
	return responses.Response(c, http.StatusOK, res)
}

// Get godoc
// @Summary Get character by ID
// @Description Get character by ID
// @ID characters-get
// @Tags Character Actions
// @Accept json
// @Produce json
// @Param id path int true "Character ID"
// @Error 404 {object} responses.ErrorResponse
// @Success 200 {object} responses.CharacterResponse
// @Router /characters/{id} [get]
func (h *CharacterHandler) Get(c echo.Context) error {
	userID := c.Get("userID").(uint)
	id := c.Param("id")

	character, err := h.server.Dependencies.GetCharacterService().GetCharacterByID(userID, id)
	switch {
	case errors.Is(err, services.ErrCharacterNotFound):
		logging.Logger.LogWarn().Msg("character/get character not found" + err.Error())
		return responses.ErrorResponse(c, http.StatusNotFound, "Character not found")
	case err != nil:
		logging.Logger.LogError().Msg("character/get error: " + err.Error())
		return responses.ErrorResponse(c, http.StatusInternalServerError, "There was an error getting the character")
	}

	res := responses.NewCharacterResponse(character)
	return responses.Response(c, http.StatusOK, res)
}
