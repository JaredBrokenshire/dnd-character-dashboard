package routes

import (
	"dnd-character-dashboard/api"
	"dnd-character-dashboard/api/handlers"
	mw "dnd-character-dashboard/pkg/middleware"
	"os"
)

func campaignRoutes(s *api.Server) {
	campaignHandler := handlers.NewCampaignHandler(s)

	campaign := s.Echo.Group("/campaigns")
	campaign.Use(mw.JWTAuth(os.Getenv("JWT_SECRET")))

	campaign.GET("", campaignHandler.List)
}
