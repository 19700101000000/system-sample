package handler

import (
	"github.com/labstack/echo"
	"github.com/mattn/go-slim"
)

func Index(c echo.Context) error {
	return render(c, "template/index.slim", slim.Values{
		"names": []int{1, 2, 3, 4},
	})
}
