package handler

import (
	"github.com/19700101000000/system-sample/api/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H(db.Categories()))
}

func GetGalleries(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H(db.Galleries()))
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H(db.User(c.Param("name"))))
}
