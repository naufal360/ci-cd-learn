package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloTest(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"response": "hello naufal",
	})
}
