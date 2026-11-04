package api

import (
	"dnd-character-dashboard/db"
	"dnd-character-dashboard/db/repositories"
	"dnd-character-dashboard/pkg/file_storage"
	"dnd-character-dashboard/services"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"time"
)

type Server struct {
	Echo         *echo.Echo
	Db           *gorm.DB
	Repos        *repositories.Repos
	Dependencies *services.Dependencies
}

func NewServer() *Server {
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	time.Local = utc

	fileStore := file_storage.NewLocalStorage("assets/files")

	gormDB, err := db.Init(fileStore)

	s := &Server{
		Echo: echo.New(),
		Db:   gormDB,
	}

	s.Repos = repositories.NewRepos(s.Db)
	s.Dependencies = services.NewDependencies(s.Db, s.Repos)
	s.Dependencies.SetFileStore(fileStore)

	return s
}

// Start runs anything needed for the application before starting the server
func (s *Server) Start(addr string) error {
	return s.Echo.Start(addr)
}
