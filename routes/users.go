package routes

import (
	"ibj-technical-test/handler"

	"github.com/labstack/echo/v4"
)

func UsersRoutes(e *echo.Echo, userHandler *handler.UsersHandler) {

	g := e.Group("/api/v1")

	g.GET("/users", userHandler.GetListUsers)
	g.GET("/users/:id", userHandler.GetUserByID)
	g.POST("/users", userHandler.CreateUser)
	g.PUT("/users/:id", userHandler.UpdateUserByID)
	g.DELETE("/users/:id", userHandler.DeleteUserByID)
}
