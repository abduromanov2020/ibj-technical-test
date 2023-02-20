package handler

import (
	"ibj-technical-test/entity"
	"ibj-technical-test/entity/response"
	"ibj-technical-test/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CoursesHandler struct {
	CoursesUsecase usecase.ICoursesUsecase
}

func NewCoursesHandler(coursesUsecase usecase.ICoursesUsecase) *CoursesHandler {
	return &CoursesHandler{CoursesUsecase: coursesUsecase}
}

func (h *CoursesHandler) CreateCourses(c echo.Context) error {
	coursesRequest := entity.CreateCourseRequest{}

	c.Bind(&coursesRequest)

	if coursesRequest.Title == "" || coursesRequest.CoursesCategoriesID == 0 {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Title and Courses Categories ID must be filled",
			Data:    []entity.Courses{},
		})
	}

	courses, err := h.CoursesUsecase.CreateCourse(coursesRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Create Courses",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Create Courses",
		Data:    courses,
	})
}

func (h *CoursesHandler) GetListCourses(c echo.Context) error {
	courses, err := h.CoursesUsecase.GetListCourses()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get List Courses",
			Data:    err.Error(),
		})
	}

	if len(courses) == 0 {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Empty List Courses",
			Data:    []entity.Courses{},
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get List Courses",
		Data:    courses,
	})
}

func (h *CoursesHandler) GetCourseByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	courses, err := h.CoursesUsecase.GetCourseByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get Course By ID",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get Course By ID",
		Data:    courses,
	})
}

func (h *CoursesHandler) UpdateCourseByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	coursesRequest := entity.UpdateCourseRequest{}

	c.Bind(&coursesRequest)

	courses, err := h.CoursesUsecase.UpdateCourseByID(id, coursesRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Update Course By ID",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Update Course By ID",
		Data:    courses,
	})
}

func (h *CoursesHandler) DeleteCourseByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.CoursesUsecase.DeleteCourseByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Delete Course By ID",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Delete Course By ID",
		Data:    []entity.Courses{},
	})
}
