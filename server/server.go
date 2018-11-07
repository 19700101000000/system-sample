package main

import (
	"github.com/ShikinamiAsuka/ih13/server/handler"
	"github.com/labstack/echo"
)

type server struct {
	e *echo.Echo
}

func newServer() *server {
	return &server{
		e: echo.New(),
	}
}

// set routing this
func (s *server) serverInit() {
	s.e.POST("/auth", handler.Auth)
}
func (s *server) serverRun() {
	s.e.Logger.Fatal(s.e.Start(":8080"))
}
