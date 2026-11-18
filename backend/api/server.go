package api

import (
	"dnd/db"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"time"
)

type Server struct {
	Echo *echo.Echo
	DB   *gorm.DB
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

	return &Server{
		Echo: echo.New(),
		DB:   database,
	}
}

func (s *Server) Start(address string) error {
	return s.Echo.Start(address)
}
