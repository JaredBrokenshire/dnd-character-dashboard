package handlers

import (
	"dnd-character-dashboard/api"
	"dnd-character-dashboard/api/responses"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CampaignHandler struct {
	server *api.Server
}

func NewCampaignHandler(server *api.Server) *CampaignHandler {
	return &CampaignHandler{server: server}
}

// List godoc
// @Summary List campaigns
// @Description List campaigns (paginated)
// @ID campaigns-list
// @Tags Campaigns Actions
// @Accept json
// @Produce json
// @Param search query string false "Search campaigns by name"
// @Param page query int false "The page number"
// @Param page_size query int false "The numbers of items to return. Max 100"
// @Success 200 {object} responses.PaginatedCampaignResponse
// @Router /campaigns [get]
func (h *CampaignHandler) List(c echo.Context) error {
	userID := c.Get("userID").(uint)

	var scopes []func(db *gorm.DB) *gorm.DB

	search := c.QueryParam("search")
	if search != "" {
		searchTerm := fmt.Sprintf("%%%s%%", search)
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("campaigns.name LIKE ?", searchTerm)
		})
	}

	campaigns, page, pageSize := h.server.Dependencies.GetCampaignService().List(c, userID, scopes)
	count := h.server.Dependencies.GetCampaignService().Count(userID, scopes)

	res := responses.NewPaginatedCampaignResponse(campaigns, count, page, pageSize)
	return responses.Response(c, http.StatusOK, res)
}
