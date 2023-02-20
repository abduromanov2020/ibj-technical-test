package config

import (
	"fmt"
	"ibj-technical-test/entity"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

var DB *gorm.DB
var err error

func Database() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected")
}

func AutoMigrate() {
	DB.AutoMigrate(&entity.Admin{})
	DB.AutoMigrate(&entity.Users{})
	DB.AutoMigrate(&entity.CoursesCategories{})
	DB.AutoMigrate(&entity.Courses{})
	DB.AutoMigrate(&entity.UserCourses{})
}
