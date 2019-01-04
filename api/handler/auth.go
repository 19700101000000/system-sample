package handler

import (
	"crypto/sha256"
	"fmt"
	"github.com/19700101000000/system-sample/api/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AuthCheck(c *gin.Context) {
	resData := map[string]interface{}{
		"status": "failed",
		"name":   nil,
	}
	name, err := c.Cookie("name")
	if err != nil {
		c.JSON(http.StatusOK, gin.H(resData))
		return
	}
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusOK, gin.H(resData))
		return
	}
	if UserList[name] != token {
		c.JSON(http.StatusOK, gin.H(resData))
		return
	}
	resData["status"] = "ok"
	resData["name"] = name
	c.JSON(http.StatusOK, gin.H(resData))
}

func AuthSignout(c *gin.Context) {
	resData := map[string]interface{}{
		"status": "ok",
	}
	name, _ := c.Cookie("name")
	token, _ := c.Cookie("token")
	if v, ok := UserList[name]; ok && v == token {
		delete(UserList, name)
	}
	c.JSON(http.StatusOK, gin.H(resData))
}

func AuthSignin(c *gin.Context) {
	var reqData struct {
		Name string `json:"name"`
		Pass string `json:"pass"`
	}
	resData := map[string]interface{}{
		"status": "failed",
	}
	/* parse request data */
	err := c.Bind(&reqData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H(resData))
		return
	}

	/* db connection */
	name := db.Auth(reqData.Name, reqData.Pass)

	if name != nil {
		resData["status"] = "ok"
		token := fmt.Sprintf("%x", sha256.Sum256([]byte(*name+time.Now().String())))
		setCookie(c, "name", *name)
		setCookie(c, "token", token)
		UserList[*name] = token
	} else {
		resData["status"] = "no-user"
	}
	c.JSON(http.StatusOK, gin.H(resData))
}
