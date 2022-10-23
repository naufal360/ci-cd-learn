package service

import (
	"praktikum/dto"
	m "praktikum/middleware"
	"praktikum/model"
	"praktikum/repository"
)

type UserService interface {
	GetAllUsers() ([]dto.UserDTO, error)
	CreateUser(payloads dto.UserDTO) (dto.UserDTO, error)
	LoginUser(payloads dto.UserDTO) (dto.UserJWT, error)
}

type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{UserRepository: userRepo}
}

func (s *userService) GetAllUsers() ([]dto.UserDTO, error) {
	allUsers, err := s.UserRepository.GetAllUsers()

	dtos := make([]dto.UserDTO, len(allUsers))

	for i, usr := range allUsers {
		dtos[i] = dto.UserDTO{
			Email:    usr.Email,
			Password: usr.Password,
		}
	}

	if err != nil {
		return dtos, err
	}

	return dtos, nil
}

func (s *userService) CreateUser(payloads dto.UserDTO) (dto.UserDTO, error) {
	userData := model.User{
		Email:    payloads.Email,
		Password: payloads.Password,
	}

	err := s.UserRepository.CreateUser(userData)

	if err != nil {
		return payloads, err
	}
	return payloads, nil
}

func (s *userService) LoginUser(payloads dto.UserDTO) (dto.UserJWT, error) {
	var returnResult dto.UserJWT

	user, errs := s.UserRepository.LoginUser(payloads)

	token, errt := m.CreateToken(int(user.ID), user.Email)

	returnResult = dto.UserJWT{
		Email: user.Email,
		Token: token,
	}

	if errs != nil {
		return returnResult, errs
	}

	if errt != nil {
		return returnResult, errs
	}

	return returnResult, nil
}
