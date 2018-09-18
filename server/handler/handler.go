package handler

import (
	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	return render(c, "template/index.slim")
}
