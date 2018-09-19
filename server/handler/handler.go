package handler

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/mattn/go-slim"
)

func Index(c echo.Context) error {
	return slimRender(c, "template/index.slim")
}

func Static(c echo.Context) error {
	name := echo.Context.Param(c, "name")
	return slimRender(c, fmt.Sprintf("template/%s.slim", name))
}

/*slim's render*/
func slimRender(c echo.Context, filepath string) error {
	tmpl, err := slim.ParseFile(filepath)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, slim.Values{"test10": make([]string, 10)}); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.HTML(http.StatusOK, buf.String())
}
