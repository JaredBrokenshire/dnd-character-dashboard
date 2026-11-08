package routes

import (
	"dnd-character-dashboard/api"
	"dnd-character-dashboard/api/handlers"
	"dnd-character-dashboard/config"
	mw "dnd-character-dashboard/pkg/middleware"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func authRoutes(s *api.Server) {
	authHandler := handlers.NewAuthHandler(s)

	auth := s.Echo.Group("/auth")
	limiterStore := middleware.NewRateLimiterMemoryStore(config.Get().LoginRateLimit)

	auth.POST("/register", authHandler.Register, middleware.RateLimiter(limiterStore))
	auth.POST("/login", authHandler.Login, middleware.RateLimiter(limiterStore))
	auth.POST("/refresh", authHandler.Refresh, middleware.RateLimiter(limiterStore))

	protected := auth.Group("")
	protected.Use(mw.JWTAuth(os.Getenv("JWT_SECRET")))

	protected.GET("/me", authHandler.Me)
}
