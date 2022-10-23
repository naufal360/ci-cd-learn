package routes

import (
	"praktikum/config"
	mid "praktikum/middleware"
	v1 "praktikum/routes/v1"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	dbConf := config.InitDB()

	controlAPI := config.InitUserAPI(dbConf)

	routes := echo.New()

	// set logger
	mid.LogMiddleware(routes)

	// Test response
	v1.GetHello(routes)

	// unauthenticated route
	v1.UserRouteUnathenticated(routes, controlAPI)
	v1.UserLogin(routes, controlAPI)

	// authenticated route
	v1.UserRouteAuthenticated(routes, controlAPI)

	return routes
}
