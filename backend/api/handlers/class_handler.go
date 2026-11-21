package handlers

import (
	"dnd/api"
	"dnd/api/responses"
	"dnd/db"
	"dnd/pkg/logging"
	"dnd/services"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))

	var scopes db.Scopes

	search := c.QueryParam("search")
	if search != "" {
		searchTerm := fmt.Sprintf("%%%s%%", search)
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("name LIKE ?", searchTerm)
		})
	}

	classes, page, pageSize := h.server.Dependencies.GetClassService().List(page, pageSize, scopes)
	count := h.server.Dependencies.GetClassService().Count(scopes)

	res := responses.NewPaginatedClassResponse(classes, count, page, pageSize)
	return responses.Response(c, http.StatusOK, res)
}

// Get godoc
// @Summary Get classes
// @Description Get class by ID
// @ID classes-get
// @Tags Class Actions
// @Accept json
// @Produce json
// @Param id path int true "Class ID"
// @Success 200 {object} responses.ClassResponse
// @Router /classes/:id [get]
func (h *ClassHandler) Get(c echo.Context) error {
	id := c.Param("id")

	class, err := h.server.Dependencies.GetClassService().Get(id)
	switch {
	case errors.Is(err, services.ErrClassNotFound):
		logging.Logger.LogWarn().Msg("classes/get class not found")
		return responses.ErrorResponse(c, http.StatusNotFound, "Class not found")
	case err != nil:
		logging.Logger.LogWarn().Msg(fmt.Sprintf("classes/get error: %v", err.Error()))
		return responses.ErrorResponse(c, http.StatusInternalServerError, "There was an error getting the class")
	}

	res := responses.NewClassResponse(class)
	return responses.Response(c, http.StatusOK, res)
}
