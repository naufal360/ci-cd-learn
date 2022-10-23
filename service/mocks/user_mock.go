package mocks

import (
	"praktikum/dto"

	"github.com/stretchr/testify/mock"
)

type UserMock struct {
	mock.Mock
}

func (u *UserMock) GetAllUsers() ([]dto.UserDTO, error) {
	args := u.Called()

	return args.Get(0).([]dto.UserDTO), args.Error(1)
}

func (u *UserMock) CreateUser(payloads dto.UserDTO) (dto.UserDTO, error) {
	args := u.Called(payloads)

	return args.Get(0).(dto.UserDTO), args.Error(1)
}

func (u *UserMock) LoginUser(payloads dto.UserDTO) (dto.UserJWT, error) {
	args := u.Called(payloads)

	return args.Get(0).(dto.UserJWT), args.Error(1)
}
