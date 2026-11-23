package routes

import (
	"dnd/api"
	"dnd/api/handlers"
	mw "dnd/pkg/middleware"
	"os"
)

func raceRoutes(s *api.Server) {
	raceHandler := handlers.NewRaceHandler(s)

	race := s.Echo.Group("/races")
	race.Use(mw.JWTAuth(os.Getenv("JWT_SECRET")))

	race.GET("", raceHandler.List)
}
