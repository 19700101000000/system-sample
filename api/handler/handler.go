package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var SqlContactStream chan<- interface{}
var UserList map[string]string

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
