package handler

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo"
	"github.com/mattn/go-slim"
	"net/http"
)

func Index(c echo.Context) error {
	name := echo.Context.Param(c, "name")
	if name == "" {
		name = "index"
	}

	tmpl, err := slim.ParseFile(fmt.Sprintf("template/%s.slim", name))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, slim.Values{}); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.HTML(http.StatusOK, buf.String())
}
