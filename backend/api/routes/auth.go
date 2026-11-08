package routes

import (
	"dnd-character-dashboard/api"
	"dnd-character-dashboard/api/handlers"
	"dnd-character-dashboard/config"
	"github.com/labstack/echo/v4/middleware"
)

func authRoutes(s *api.Server) {
	authHandler := handlers.NewAuthHandler(s)

	auth := s.Echo.Group("/auth")
	limiterStore := middleware.NewRateLimiterMemoryStore(config.Get().LoginRateLimit)

	auth.POST("/register", authHandler.Register, middleware.RateLimiter(limiterStore))

}
