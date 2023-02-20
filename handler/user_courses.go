package handler

import (
	"ibj-technical-test/entity"
	"ibj-technical-test/entity/response"
	"ibj-technical-test/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserCoursesHandler struct {
	UserCoursesUsecase usecase.IUserCoursesUsecase
}

func NewUserCoursesHandler(userCoursesUsecase usecase.IUserCoursesUsecase) *UserCoursesHandler {
	return &UserCoursesHandler{UserCoursesUsecase: userCoursesUsecase}
}

func (u *UserCoursesHandler) CreateUserCourse(c echo.Context) error {
	userCourseRequest := entity.CreateUserCourseRequest{}

	c.Bind(&userCourseRequest)

	if userCourseRequest.UsersID == 0 || userCourseRequest.CoursesID == 0 {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "ID and CourseID must be filled",
			Data:    []entity.GetUserCourseResponse{},
		})
	}

	userCourse, err := u.UserCoursesUsecase.CreateUserCourses(userCourseRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Req Body",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Create User Course",
		Data:    userCourse,
	})
}

func (u *UserCoursesHandler) GetListUserCourses(c echo.Context) error {
	userCourses, err := u.UserCoursesUsecase.GetListUserCourses()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get List User Courses",
			Data:    err.Error(),
		})
	}

	if len(userCourses) == 0 {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Empty List User Courses",
			Data:    []entity.GetUserCourseResponse{},
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get List User Courses",
		Data:    userCourses,
	})
}

func (u *UserCoursesHandler) GetUserCourseByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	userCourse, err := u.UserCoursesUsecase.GetUserCoursesByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get User Course By ID",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get User Course By ID",
		Data:    userCourse,
	})
}

func (u *UserCoursesHandler) UpdateUserCourseByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	userCourseRequest := entity.UpdateUserCourseRequest{}

	c.Bind(&userCourseRequest)

	userCourse, err := u.UserCoursesUsecase.UpdateUserCoursesByID(id, userCourseRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Update User Course",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Update User Course",
		Data:    userCourse,
	})
}

func (u *UserCoursesHandler) DeleteUserCourseByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := u.UserCoursesUsecase.DeleteUserCoursesByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Delete User Course",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Delete User Course",
		Data:    []entity.GetUserCourseResponse{},
	})
}
