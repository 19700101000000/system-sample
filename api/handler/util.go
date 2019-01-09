package handler

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo"
	"github.com/mattn/go-slim"
)

/*render slim to html*/
func render(c echo.Context, path string, v slim.Values) error {
	// import template
	tmpl, err := slim.ParseFile(path)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// render slim
	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, v); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// result
	return c.HTML(http.StatusOK, buf.String())
}

/* check auth */
func getAuth(c *gin.Context) (string, bool) {

	name, err := c.Cookie("name")
	if err != nil {
		return "", false
	}
	token, err := c.Cookie("token")
	if err != nil {
		return "", false
	}
	if UserList[name] != token {
		return "", false
	}
	return name, true
}

/* set cookie */
func setCookie(c *gin.Context, name, value string) {
	c.SetCookie(name, value, 86400, "/", "localhost", false, true)
}
