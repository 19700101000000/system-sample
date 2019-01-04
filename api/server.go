package main

import (
	"github.com/19700101000000/system-sample/api/db"
	"github.com/19700101000000/system-sample/api/handler"
	"github.com/gin-gonic/gin"
)

type server struct {
	router *gin.Engine
}

func newServer() *server {
	return &server{
		router: gin.Default(),
	}
}
func (s *server) serverInit() {
	db.InitDB()
	handler.UserList = make(map[string]string)

	s.router.GET("/", handler.Index)

	s.router.GET("/auth/check", handler.AuthCheck)
	s.router.GET("/auth/signout", handler.AuthSignout)
	s.router.POST("/auth/signin", handler.AuthSignin)

	/* OAuth2 from GOOGLE */
	// s.router.GET("/auth/google/signin", handler.AuthGoogleSignin)
	// s.router.GET("/auth/google/callback", handler.AuthGoogleCallback)
}
func (s *server) serverRun() {
	defer db.CloseDB()
	s.router.Run(":8080")
}
