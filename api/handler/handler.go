package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	ID    int
	Name  string
	Token string
}

var UserList map[string]UserInfo

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
