package controller

import (
	"net/http"
	"praktikum/dto"
	"praktikum/service"

	"github.com/labstack/echo/v4"
)

type UserControllerInterface interface{}

type UserController struct {
	userServ service.UserService
}

func NewUserController(userServ service.UserService) *UserController {
	return &UserController{
		userServ: userServ,
	}
}

func (u *UserController) GetAllUsers(c echo.Context) error {

	dataUsers, err := u.userServ.GetAllUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": dataUsers,
	})
}

func (u *UserController) CreateUser(c echo.Context) error {

	var payloads dto.UserDTO

	if err := c.Bind(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	user, err := u.userServ.CreateUser(payloads)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": user,
	})
}

func (u *UserController) LoginUser(c echo.Context) error {
	var payloads dto.UserDTO

	if err := c.Bind(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	authUser, err := u.userServ.LoginUser(payloads)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user token": authUser,
	})
}
