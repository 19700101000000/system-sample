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

/* set cookie */
func setCookie(c *gin.Context, name, value string) {
	c.SetCookie(name, value, 86400, "/", "localhost", false, true)
}
