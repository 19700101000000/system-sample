package main

import (
	"github.com/ShikinamiAsuka/system-sample/server/handler"
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
func (s *server) serverInit() {
	s.e.GET("/", handler.Index)
	s.e.GET("/:name", handler.Index)
}
func (s *server) serverRun() {
	s.e.Logger.Fatal(s.e.Start(":8080"))
}
