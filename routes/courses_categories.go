package routes

import (
	"ibj-technical-test/handler"

	"github.com/labstack/echo/v4"
)

func CoursesCategoriesRoutes(e *echo.Echo, coursesCategoriesHandler *handler.CoursesCategoriesHandler) {

	g := e.Group("/api/v1")

	g.GET("/courses-categories", coursesCategoriesHandler.GetListCoursesCategories)
	g.GET("/courses-categories/:id", coursesCategoriesHandler.GetCoursesCategoriesByID)
	g.POST("/courses-categories", coursesCategoriesHandler.CreateCoursesCategories)
	g.PUT("/courses-categories/:id", coursesCategoriesHandler.UpdateCoursesCategoriesByID)
	g.DELETE("/courses-categories/:id", coursesCategoriesHandler.DeleteCoursesCategoriesByID)
}
