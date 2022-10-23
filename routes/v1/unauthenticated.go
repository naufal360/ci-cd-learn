package v1

import (
	"praktikum/controller"

	"github.com/labstack/echo/v4"
)

func UserRouteUnathenticated(routes *echo.Echo, api *controller.UserController) {
	user := routes.Group("/v1")
	{
		user.POST("/users", api.CreateUser)
	}
}

func UserLogin(routes *echo.Echo, api *controller.UserController) {
	user := routes.Group("/v1")
	{
		user.POST("/login", api.LoginUser)
	}
}

func GetHello(routes *echo.Echo) {
	user := routes.Group("/v1")
	{
		user.GET("/hello", controller.HelloTest)
	}
}
