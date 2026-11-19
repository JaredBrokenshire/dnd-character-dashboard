package routes

import (
	"dnd/api"
	"dnd/api/handlers"
	mw "dnd/pkg/middleware"
	"os"
)

func fileRoutes(server *api.Server) {
	fileHandler := handlers.NewFileHandler(server)

	file := server.Echo.Group("/files")
	file.Use(mw.JWTAuth(os.Getenv("JWT_SECRET")))

	file.GET("/:filepath", fileHandler.Get)
}
