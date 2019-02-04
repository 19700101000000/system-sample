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

/* upload image */
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

/* upload wanted */
func UploadWanted(c *gin.Context) {
	// auth check
	name, ok := getAuth(c)
	if !ok {
		c.String(http.StatusForbidden, "forbidden")
		return
	}

	// get wanted
	var reqData struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Price       int    `json:"price"`
	}
	err := c.Bind(&reqData)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	ok = db.InsertWanted(UserList[name].ID, db.StructWanted{
		Title:       reqData.Title,
		Description: reqData.Description,
		Price:       reqData.Price,
	})

	c.JSON(http.StatusOK, gin.H{})
}

/* upload request */
func UploadRequest(c *gin.Context) {
	name, ok := getAuth(c)
	if !ok {
		c.String(http.StatusForbidden, "forbidden")
		return
	}

	// get request
	var reqData struct {
		OwnerName   string `json:"ownername"`
		WantedID    int    `json:"wanted"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Price       int    `json:"price"`
	}
	err := c.Bind(&reqData)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	// insert
	ok = db.InsertRequest(UserList[name].ID, db.StructRequest{
		Wanted: db.StructWanted{
			Username: reqData.OwnerName,
			Number:   reqData.WantedID,
		},
		Title:       reqData.Title,
		Description: reqData.Description,
		Price:       reqData.Price,
	})
	if !ok {
		c.String(http.StatusInternalServerError, "internal server error")
		return
	}

	c.String(http.StatusOK, "ok")
}
