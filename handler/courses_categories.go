package handler

import (
	"ibj-technical-test/entity"
	"ibj-technical-test/entity/response"
	"ibj-technical-test/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CoursesCategoriesHandler struct {
	CoursesCategoriesUsecase usecase.ICoursesCategoriesUsecase
}

func NewCoursesCategoriesHandler(coursesCategoriesUsecase usecase.ICoursesCategoriesUsecase) *CoursesCategoriesHandler {
	return &CoursesCategoriesHandler{CoursesCategoriesUsecase: coursesCategoriesUsecase}
}

func (h *CoursesCategoriesHandler) CreateCoursesCategories(c echo.Context) error {
	coursesCategoriesRequest := entity.CreateCoursesCategoriesRequest{}

	c.Bind(&coursesCategoriesRequest)

	if coursesCategoriesRequest.Name == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Name must be filled",
			Data:    []entity.GetCoursesCategoriesResponse{},
		})
	}

	coursesCategories, err := h.CoursesCategoriesUsecase.CreateCoursesCategories(coursesCategoriesRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Create Courses Categories",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Create Courses Categories",
		Data:    coursesCategories,
	})
}

func (h *CoursesCategoriesHandler) GetListCoursesCategories(c echo.Context) error {
	coursesCategories, err := h.CoursesCategoriesUsecase.GetListCoursesCategories()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get List Courses Categories",
			Data:    err.Error(),
		})
	}

	if len(coursesCategories) == 0 {
		return c.JSON(http.StatusOK, response.BaseResponse{
			Code:    http.StatusOK,
			Message: "Empty List Courses Categories",
			Data:    []entity.GetCoursesCategoriesResponse{},
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get List Courses Categories",
		Data:    coursesCategories,
	})
}

func (h *CoursesCategoriesHandler) GetCoursesCategoriesByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	coursesCategories, err := h.CoursesCategoriesUsecase.GetCoursesCategoriesByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get Courses Categories By ID",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get Courses Categories By ID",
		Data:    coursesCategories,
	})
}

func (h *CoursesCategoriesHandler) UpdateCoursesCategoriesByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	coursesCategoriesRequest := entity.UpdateCoursesCategoriesRequest{}

	c.Bind(&coursesCategoriesRequest)

	coursesCategories, err := h.CoursesCategoriesUsecase.UpdateCoursesCategoriesByID(id, coursesCategoriesRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Update Courses Categories By ID",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Update Courses Categories By ID",
		Data:    coursesCategories,
	})

}

func (h *CoursesCategoriesHandler) DeleteCoursesCategoriesByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.CoursesCategoriesUsecase.DeleteCoursesCategoriesByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Delete Courses Categories By ID",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Delete Courses Categories By ID",
		Data:    []entity.GetCoursesCategoriesResponse{},
	})
}
