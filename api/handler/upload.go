package handler

import (
	"crypto/sha256"
	"fmt"
	"github.com/19700101000000/system-sample/api/db"
	"github.com/gin-gonic/gin"
	"net/http"
	// "path/filepath"
	"time"
)

func UploadImage(c *gin.Context) {
	name, ok := getAuth(c)
	if !ok {
		c.String(http.StatusForbidden, "forbidden")
		return
	}
	file, err := c.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	categories := c.PostFormArray("categories")

	// TODO resiging image
	contentType := file.Header["Content-Type"][0]
	nowTime := time.Now().String()

	if contentType != "image/jpeg" && contentType != "image/png" {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	// filename := filepath.Base(file.Filename)
	filename := fmt.Sprintf("%x%x", sha256.Sum256([]byte(name)), sha256.Sum256([]byte(nowTime)))
	if err := c.SaveUploadedFile(file, fmt.Sprintf("./public/images/%s", filename)); err != nil {
		fmt.Printf("error uploadImage: %v\n", err)
	}
	ok = db.InsertGallery(UserList[name].ID, filename, categories)
	if !ok {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
