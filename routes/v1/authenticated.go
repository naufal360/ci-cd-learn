package v1

import (
	"os"
	"praktikum/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserRouteAuthenticated(routes *echo.Echo, api *controller.UserController) {
	authUser := routes.Group("/v1")
	authUser.Use(middleware.JWT([]byte(os.Getenv("SECRET_KEY"))))

	{
		authUser.GET("/users", api.GetAllUsers)
	}
}
