package handler

import (
	"bytes"
	"github.com/labstack/echo"
	"github.com/mattn/go-slim"
	"net/http"
)

/*render slim to html*/
func render(c echo.Context, fileName string) error {
	tmpl, err := slim.ParseFile(fileName)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, slim.Values{}); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.HTML(http.StatusOK, buf.String())
}
