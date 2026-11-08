package routes

import (
	"dnd-character-dashboard/api"
	"dnd-character-dashboard/api/handlers"
	mw "dnd-character-dashboard/pkg/middleware"
	"os"
)

func classRoutes(s *api.Server) {
	classHandler := handlers.NewClassHandler(s)

	class := s.Echo.Group("/classes")
	class.Use(mw.JWTAuth(os.Getenv("JWT_SECRET")))

	class.GET("", classHandler.List)
}
