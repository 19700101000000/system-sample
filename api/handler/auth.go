package handler

import (
	"github.com/19700101000000/system-sample/api/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthCheck(c *gin.Context) {
	resData := map[string]interface{}{
		"status": "failed",
		"name":   nil,
	}
	name, ok := getAuth(c)
	if !ok {
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
	if v, ok := UserList[name]; ok && v.Token == token {
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
	id, name, ok := db.Auth(reqData.Name, reqData.Pass)

	if ok {
		resData["status"] = "ok"
		addUserList(c, id, name)
	} else {
		resData["status"] = "no-user"
	}
	c.JSON(http.StatusOK, gin.H(resData))
}

/* sign up */
func AuthSignup(c *gin.Context) {
	var reqData struct {
		Name string `json:"name"`
		Pass string `json:"password"`
	}
	err := c.Bind(&reqData)
	if err != nil || len(reqData.Name) == 0 || len(reqData.Pass) == 0 {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	id, ok := db.InsertUser(reqData.Name, reqData.Pass)
	if ok {
		addUserList(c, id, reqData.Name)
		c.String(http.StatusOK, "ok")
	} else {
		c.String(http.StatusBadRequest, "bad request")
	}
}
