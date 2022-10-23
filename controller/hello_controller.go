package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func HelloTest(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"response": "hello naufal",
		"time":     time.Now(),
	})
}
