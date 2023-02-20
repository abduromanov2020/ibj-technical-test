package main

import (
	"ibj-technical-test/config"
	"ibj-technical-test/handler"
	"ibj-technical-test/repository"
	"ibj-technical-test/routes"
	"ibj-technical-test/usecase"

	"github.com/labstack/echo/v4"
)

func main() {

	config.Database()
	config.AutoMigrate()

	e := echo.New()

	// init user
	userRepository := repository.NewUsersRepository(config.DB)
	userUsecase := usecase.NewUsersUsecase(userRepository)
	userHandler := handler.NewUsersHandler(userUsecase)

	routes.UsersRoutes(e, userHandler)

	// init admin
	adminRepository := repository.NewAdminRepository(config.DB)
	adminUsecase := usecase.NewAdminUsecase(adminRepository)
	adminHandler := handler.NewAdminHandler(adminUsecase)

	routes.AdminRoutes(e, adminHandler)

	// init courses categories
	coursesCategoriesRepository := repository.NewCoursesCategoriesRepository(config.DB)
	coursesCategoriesUsecase := usecase.NewCoursesCategoriesUsecase(coursesCategoriesRepository)
	coursesCategoriesHandler := handler.NewCoursesCategoriesHandler(coursesCategoriesUsecase)

	routes.CoursesCategoriesRoutes(e, coursesCategoriesHandler)

	// init courses
	coursesRepository := repository.NewCoursesRepository(config.DB)
	coursesUsecase := usecase.NewCoursesUsecase(coursesRepository)
	coursesHandler := handler.NewCoursesHandler(coursesUsecase)

	routes.CoursesRoutes(e, coursesHandler)

	// init user courses
	userCoursesRepository := repository.NewUserCoursesRepository(config.DB)
	userCoursesUsecase := usecase.NewUserCoursesUsecase(userCoursesRepository)
	userCoursesHandler := handler.NewUserCoursesHandler(userCoursesUsecase)

	routes.UserCoursesRoutes(e, userCoursesHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
