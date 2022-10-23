package config

import (
	"praktikum/controller"
	"praktikum/repository"
	"praktikum/service"

	"gorm.io/gorm"
)

func InitUserAPI(db *gorm.DB) *controller.UserController {
	userRepo := repository.NewUserRepository(db)
	userServ := service.NewUserService(userRepo)
	userAPI := controller.NewUserController(userServ)
	return userAPI
}
