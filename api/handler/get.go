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

func GetWorksWanteds(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H(db.WorksWanteds(c.Param("name"), false)))
}
func GetMyWanteds(c *gin.Context) {
	name, ok := getAuth(c)
	if !ok {
		c.String(http.StatusForbidden, "forbidden")
		return
	}

	c.JSON(http.StatusOK, db.WorksWanteds(name, ok))
}
