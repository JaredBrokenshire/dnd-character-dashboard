package routes

import (
	"dnd-character-dashboard/api"
	mw "dnd-character-dashboard/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"os"
	"strings"
)

func RegisterRoutes(s *api.Server) {
	// ── Global middleware ──────────────────────────────────────────────────────
	s.Echo.Use(mw.LoggingMiddleware())
	s.Echo.Use(middleware.Recover())

	// CORS origins are configured via the CORS_ORIGINS env var (comma-separated).
	// Falls back to the development default so existing local setups keep working.
	rawOrigins := os.Getenv("CORS_ORIGINS")
	if rawOrigins == "" {
		rawOrigins = "http://localdnd.com:8080"
	}
	allowedOrigins := strings.Split(rawOrigins, ",")
	for i, o := range allowedOrigins {
		allowedOrigins[i] = strings.TrimSpace(o)
	}

	s.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: allowedOrigins,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
		AllowCredentials: true,
	}))

	// ── Swagger UI ────────────────────────────────────────────────────────────
	s.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	authRoutes(s)
	fileRoutes(s)

	characterRoutes(s)
	campaignRoutes(s)
}
