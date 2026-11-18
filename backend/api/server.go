package api

import (
	"github.com/labstack/echo/v4"
	"time"
)

type Server struct {
	Echo *echo.Echo
}

func NewServer() *Server {
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	time.Local = utc

	return &Server{
		Echo: echo.New(),
	}
}

func (s *Server) Start(address string) error {
	return s.Echo.Start(address)
}
