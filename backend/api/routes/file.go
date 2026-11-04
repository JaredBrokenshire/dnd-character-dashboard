package routes

import (
	"dnd-character-dashboard/api"
	"dnd-character-dashboard/api/handlers"
	mw "dnd-character-dashboard/pkg/middleware"
	"os"
)

func fileRoutes(server *api.Server) {
	fileHandler := handlers.NewFileHandler(server)

	file := server.Echo.Group("/files")
	file.Use(mw.JWTAuth(os.Getenv("JWT_SECRET")))

	file.GET("/:filepath", fileHandler.Get)
}
