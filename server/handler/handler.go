package handler

import (
	"bytes"
	"github.com/labstack/echo"
	"github.com/mattn/go-slim"
	"net/http"
)

func Index(c echo.Context) error {
	tmpl, err := slim.ParseFile("template/index.slim")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, slim.Values{}); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.HTML(http.StatusOK, buf.String())
}
