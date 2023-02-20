package routes

import (
	"ibj-technical-test/handler"
	"ibj-technical-test/middleware"

	"github.com/labstack/echo/v4"
)

func AdminRoutes(e *echo.Echo, adminHandler *handler.AdminHandler) {

	g := e.Group("/api/v1")

	g.POST("/login", adminHandler.LoginAdmin)
	g.GET("/admin", adminHandler.GetListAdmin, middleware.IsAuthenticated)
	g.GET("/admin/:id", adminHandler.GetAdminByID, middleware.IsAuthenticated)
	g.POST("/admin", adminHandler.CreateAdmin)
	g.PUT("/admin/:id", adminHandler.UpdateAdminByID, middleware.IsAuthenticated)
	g.DELETE("/admin/:id", adminHandler.DeleteAdminByID, middleware.IsAuthenticated)
}
