package main

import (
	"github.com/19700101000000/system-sample/server/handler"
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
	s.e.GET("/go", handler.Index)
}
func (s *server) serverRun() {
	s.e.Logger.Fatal(s.e.Start(":8080"))
}
