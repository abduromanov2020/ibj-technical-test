package handler

import (
	"ibj-technical-test/entity"
	"ibj-technical-test/entity/response"
	"ibj-technical-test/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	UsersUsecase usecase.IUsersUsecase
}

func NewUsersHandler(usersUsecase usecase.IUsersUsecase) *UsersHandler {
	return &UsersHandler{UsersUsecase: usersUsecase}
}

func (u *UsersHandler) CreateUser(c echo.Context) error {
	userRequest := entity.CreateUserRequest{}

	c.Bind(&userRequest)

	if userRequest.Name == "" || userRequest.Email == "" || userRequest.Password == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Name, Email, and Password must be filled",
			Data:    []entity.GetUserResponse{},
		})
	}

	user, err := u.UsersUsecase.CreateUser(userRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Req Body",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Create User",
		Data:    user,
	})
}

func (u *UsersHandler) GetListUsers(c echo.Context) error {
	users, err := u.UsersUsecase.GetListUsers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get List Users",
			Data:    err.Error(),
		})
	}

	if len(users) == 0 {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Empty List User",
			Data:    []entity.GetUserResponse{},
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get List Users",
		Data:    users,
	})
}

func (u *UsersHandler) GetUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := u.UsersUsecase.GetUserByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get User By ID",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get User By ID",
		Data:    user,
	})
}

func (u *UsersHandler) UpdateUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	userRequest := entity.UpdateUserRequest{}

	c.Bind(&userRequest)

	user, err := u.UsersUsecase.UpdateUserByID(id, userRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Update User By ID",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Update User By ID",
		Data:    user,
	})
}

func (u *UsersHandler) DeleteUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := u.UsersUsecase.DeleteUserByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Delete User By ID",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Delete User By ID",
		Data:    []entity.GetUserResponse{},
	})
}
