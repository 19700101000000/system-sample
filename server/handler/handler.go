package handler

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/mattn/go-slim"
)

func Auth(c echo.Context) error {
	login := new(LoginRequest)
	if err := c.Bind(login); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	data := AuthData{
		Username:    login.Username,
		AccessToken: "bar",
	}
	result := &AuthResponse{
		Success: true,
		Auth:    data,
	}

	return c.JSON(http.StatusOK, result)
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
