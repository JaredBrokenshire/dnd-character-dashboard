package routes

import (
	"dnd-character-dashboard/api"
	"dnd-character-dashboard/api/handlers"
	mw "dnd-character-dashboard/pkg/middleware"
	"os"
)

func raceRoutes(s *api.Server) {
	raceHandler := handlers.NewRaceHandler(s)

	race := s.Echo.Group("/races")
	race.Use(mw.JWTAuth(os.Getenv("JWT_SECRET")))

	race.GET("", raceHandler.List)
	race.GET("/:id", raceHandler.Get)
}
