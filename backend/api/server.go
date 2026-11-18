package api

import (
	"dnd/db"
	"dnd/db/repositories"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"time"
)

type Server struct {
	Echo  *echo.Echo
	DB    *gorm.DB
	Repos *repositories.Repos
}

func NewServer() *Server {
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	time.Local = utc

	database, err := db.Init()
	if err != nil {
		panic(err)
	}

	s := &Server{
		Echo: echo.New(),
		DB:   database,
	}

	s.Repos = repositories.NewRepos(s.DB)

	return s
}

func (s *Server) Start(address string) error {
	return s.Echo.Start(address)
}
