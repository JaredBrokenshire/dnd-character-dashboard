package routes

import (
	"dnd-character-dashboard/api"
	"dnd-character-dashboard/api/handlers"
	mw "dnd-character-dashboard/pkg/middleware"
	"os"
)

func characterRoutes(s *api.Server) {
	characterHandler := handlers.NewCharacterHandler(s)

	character := s.Echo.Group("/characters")
	character.Use(mw.JWTAuth(os.Getenv("JWT_SECRET")))

	character.GET("", characterHandler.List)
	character.GET("/:id", characterHandler.Get)
}
