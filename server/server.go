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
	s.e.POST("/display/:listtype", handler.DisplayOrdersTable)
	s.e.POST("/display/db", handler.DisplayOrdersTable)
	s.e.POST("/item/clients", handler.ItemClients)
	s.e.POST("/item/employees", handler.ItemEmployees)

	s.e.GET("/display/:listtype", handler.DisplayOrdersTable)
	s.e.GET("/display/db", handler.DisplayOrdersTable)
	s.e.GET("/sample/:num", handler.Sample)
	s.e.GET("/sample/db", handler.SampleDB)
	s.e.GET("/sample/:num", handler.Sample)
}
func (s *server) serverRun() {
	s.e.Logger.Fatal(s.e.Start(":8080"))
}
