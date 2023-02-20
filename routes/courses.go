package routes

import (
	"ibj-technical-test/handler"

	"github.com/labstack/echo/v4"
)

func CoursesRoutes(e *echo.Echo, coursesHandler *handler.CoursesHandler) {

	g := e.Group("/api/v1")

	g.GET("/courses", coursesHandler.GetListCourses)
	g.GET("/courses/:id", coursesHandler.GetCourseByID)
	g.POST("/courses", coursesHandler.CreateCourses)
	g.PUT("/courses/:id", coursesHandler.UpdateCourseByID)
	g.DELETE("/courses/:id", coursesHandler.DeleteCourseByID)
}
