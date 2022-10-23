package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"praktikum/dto"
	mocks "praktikum/service/mocks"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteUsers struct {
	suite.Suite
	controller *UserController
	mockServ   *mocks.UserMock
}

func (s *suiteUsers) SetupSuite() {
	mocks := &mocks.UserMock{}
	s.mockServ = mocks

	s.controller = &UserController{
		userServ: mocks,
	}
}

func (s *suiteUsers) TestGetAllUsers() {

	s.T().Run("success", func(t *testing.T) {
		s.mockServ.On("GetAllUsers").Return([]dto.UserDTO{
			{
				Email:    "ahmad@gmail.com",
				Password: "aye123",
			},
		}, nil)

		r := httptest.NewRequest(http.MethodGet, "/users", nil)
		w := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(r, w)

		err := s.controller.GetAllUsers(ctx)
		s.NoError(err)

		s.Equal(http.StatusOK, w.Result().StatusCode)
	})

	s.T().Run("failed", func(t *testing.T) {
		s.mockServ.On("GetAllUsers").Return(nil, fmt.Errorf("Cannot connect db"))

		r := httptest.NewRequest(http.MethodGet, "/users", nil)
		w := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(r, w)

		s.controller.GetAllUsers(ctx)

		s.Equal(http.StatusOK, w.Result().StatusCode)
	})
}

func (s *suiteUsers) TestCreateUser() {

	s.T().Run("success create", func(t *testing.T) {
		s.mockServ.On("CreateUser", mock.Anything).Return(dto.UserDTO{
			Email:    "ahmad@gmail.com",
			Password: "aye123",
		}, nil)

		expectBody := dto.UserDTO{
			Email:    "ahmad@gmail.com",
			Password: "aye123",
		}

		res, err := json.Marshal(expectBody)
		s.NoError(err)
		r := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(res))
		w := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.Request().Header.Set("Content-Type", "application/json; charset=UTF-8")
		s.controller.CreateUser(ctx)

		s.Equal(http.StatusOK, w.Result().StatusCode)

		var resp map[string]dto.UserDTO
		json.NewDecoder(w.Result().Body).Decode(&resp)

		s.Equal(expectBody, resp["data"])
	})

	s.T().Run("failed create", func(t *testing.T) {
		s.mockServ.On("CreateUser", dto.UserDTO{}).Return(dto.UserDTO{}, fmt.Errorf("Error"))

		r := httptest.NewRequest(http.MethodPost, "/users", nil)
		w := httptest.NewRecorder()
		w.Result().StatusCode = http.StatusBadRequest

		e := echo.New()
		ctx := e.NewContext(r, w)
		s.controller.CreateUser(ctx)

		s.Equal(http.StatusBadRequest, w.Result().StatusCode)
	})
}

func (s *suiteUsers) TestLoginUser() {

	s.T().Run("succes login", func(t *testing.T) {
		s.mockServ.On("LoginUser", dto.UserDTO{
			"ahmad@gmail.com",
			"aye@123",
		}).Return(dto.UserJWT{
			Email: "ahmad@gmail.com",
		}, nil)

		expectedBody := dto.UserDTO{
			Email:    "ahmad@gmail.com",
			Password: "aye@123",
		}
		res, _ := json.Marshal(expectedBody)
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(res))
		w := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.Request().Header.Set("Content-Type", "application/json; charset=UTF-8")

		err := s.controller.LoginUser(ctx)
		s.NoError(err)
		s.Equal(http.StatusOK, w.Result().StatusCode)

		var resp map[string]dto.UserJWT
		json.NewDecoder(w.Result().Body).Decode(&resp)

		s.Equal(expectedBody.Email, resp["user token"].Email)
	})

	s.T().Run("failed login", func(t *testing.T) {
		s.mockServ.On("LoginUser", dto.UserDTO{}).Return(dto.UserJWT{
			Email: "ahmad@gmail.com",
		}, fmt.Errorf("Error"))

		r := httptest.NewRequest(http.MethodPost, "/", nil)
		w := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(r, w)

		err := s.controller.LoginUser(ctx)
		s.NoError(err)
		s.Equal(http.StatusUnauthorized, w.Result().StatusCode)
	})

}

func TestSuiteUsers(t *testing.T) {
	suite.Run(t, new(suiteUsers))
}
