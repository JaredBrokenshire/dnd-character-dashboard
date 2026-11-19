package routes

import (
	"dnd/api"
	"dnd/api/handlers"
	mw "dnd/pkg/middleware"
	"os"
)

func classRoutes(s *api.Server) {
	classHandler := handlers.NewClassHandler(s)

	class := s.Echo.Group("/classes")
	class.Use(mw.JWTAuth(os.Getenv("JWT_SECRET")))

	class.GET("", classHandler.List)
}
