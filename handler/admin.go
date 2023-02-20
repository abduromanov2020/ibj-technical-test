package handler

import (
	"ibj-technical-test/entity"
	"ibj-technical-test/entity/response"
	"ibj-technical-test/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	AdminUsecase usecase.IAdminUsecase
}

func NewAdminHandler(adminUsecase usecase.IAdminUsecase) *AdminHandler {
	return &AdminHandler{AdminUsecase: adminUsecase}
}

func (a *AdminHandler) LoginAdmin(c echo.Context) error {
	loginRequest := entity.LoginAdminRequest{}

	loginRequest.Email = c.FormValue("email")
	loginRequest.Password = c.FormValue("password")

	if loginRequest.Email == "" || loginRequest.Password == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Email and Password must be filled",
			Data:    []entity.GetAdminResponse{},
		})
	}

	admin, err := a.AdminUsecase.LoginAdmin(loginRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Login Failed",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Login Admin",
		Data:    admin,
	})
}

func (a *AdminHandler) CreateAdmin(c echo.Context) error {
	adminRequest := entity.CreateAdminRequest{}

	c.Bind(&adminRequest)

	if adminRequest.Name == "" || adminRequest.Email == "" || adminRequest.Password == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Name, Email, and Password must be filled",
			Data:    []entity.GetAdminResponse{},
		})
	}

	admin, err := a.AdminUsecase.CreateAdmin(adminRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Req Body",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Create Admin",
		Data:    admin,
	})
}

func (a *AdminHandler) GetListAdmin(c echo.Context) error {
	admins, err := a.AdminUsecase.GetListAdmin()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Req Body",
			Data:    err.Error(),
		})
	}

	if len(admins) == 0 {
		return c.JSON(http.StatusOK, response.BaseResponse{
			Code:    http.StatusOK,
			Message: "Empty List Admin",
			Data:    []entity.GetAdminResponse{},
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get List Admins",
		Data:    admins,
	})
}

func (a *AdminHandler) GetAdminByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	admin, err := a.AdminUsecase.GetAdminByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get Admin By ID",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get Admin",
		Data:    admin,
	})
}

func (a *AdminHandler) UpdateAdminByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	adminRequest := entity.UpdateAdminRequest{}

	c.Bind(&adminRequest)

	admin, err := a.AdminUsecase.UpdateAdminByID(id, adminRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Update Admin",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Update Admin",
		Data:    admin,
	})
}

func (a *AdminHandler) DeleteAdminByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := a.AdminUsecase.DeleteAdminByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Delete Admin",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Delete Admin",
		Data:    []entity.GetAdminResponse{},
	})
}
