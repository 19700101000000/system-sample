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
	handler.UserList = make(map[string]handler.UserInfo)

	// s.router.GET("/", handler.Index)
	s.router.GET("/get/categories", handler.GetCategories)
	s.router.GET("/get/galleries", handler.GetGalleries)
	s.router.GET("/get/galleries/:name", handler.GetMyGalleries)
	s.router.GET("/get/user/:name", handler.GetUser)
	s.router.GET("/get/works/wanteds", handler.GetMyWanteds)
	s.router.GET("/get/works/wanteds/:name", handler.GetWorksWanteds)
	s.router.GET("/get/works/requests", handler.GetMyRequests)
	s.router.GET("/get/works/requests/:wanted", handler.GetWorksRequests)
	s.router.GET("/get/works/info", handler.GetWorksInfo)
	s.router.GET("/get/rates/:name", handler.GetRates)

	/* auth */
	s.router.GET("/auth/check", handler.AuthCheck)
	s.router.GET("/auth/signout", handler.AuthSignout)
	s.router.POST("/auth/signin", handler.AuthSignin)
	s.router.POST("/auth/signup", handler.AuthSignup)

	/* galleries */
	s.router.Static("/images", "./public/images")
	s.router.POST("/upload/image", handler.UploadImage)

	/* wanted and requests */
	s.router.POST("/upload/wanted", handler.UploadWanted)
	s.router.POST("/upload/request", handler.UploadRequest)

	s.router.POST("/update/wanted/status", handler.UpdateWantedStatus)
	s.router.POST("/update/request/checked", handler.UpdateRequestChecked)
	s.router.POST("/update/request/status", handler.UpdateRequestStatus)

	/* evaluation */
	s.router.POST("/update/evaluate", handler.UpdateEvaluate)

	/* OAuth2 from GOOGLE */
	// s.router.GET("/auth/google/signin", handler.AuthGoogleSignin)
	// s.router.GET("/auth/google/callback", handler.AuthGoogleCallback)
}
func (s *server) serverRun() {
	defer db.CloseDB()
	s.router.Run(":8080")
}
