package routes

import (
	"ibj-technical-test/handler"

	"github.com/labstack/echo/v4"
)

func UserCoursesRoutes(e *echo.Echo, userCoursesHandler *handler.UserCoursesHandler) {

	g := e.Group("/api/v1")

	g.GET("/user-courses", userCoursesHandler.GetListUserCourses)
	g.GET("/user-courses/:id", userCoursesHandler.GetUserCourseByID)
	g.POST("/user-courses", userCoursesHandler.CreateUserCourse)
	g.PUT("/user-courses/:id", userCoursesHandler.UpdateUserCourseByID)
	g.DELETE("/user-courses/:id", userCoursesHandler.DeleteUserCourseByID)
}
