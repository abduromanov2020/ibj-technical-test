package repository

import (
	"errors"
	"ibj-technical-test/entity"

	"gorm.io/gorm"
)

type ICoursesRepository interface {
	Create(course entity.Courses) (entity.Courses, error)
	GetAll() ([]entity.Courses, error)
	GetByID(id int) (entity.Courses, error)
	Update(course entity.Courses) (entity.Courses, error)
	Delete(id int) error
}

type CoursesRepository struct {
	db *gorm.DB
}

func NewCoursesRepository(db *gorm.DB) ICoursesRepository {
	return &CoursesRepository{db: db}
}

func (c *CoursesRepository) Create(course entity.Courses) (entity.Courses, error) {

	if err := c.db.First(&entity.CoursesCategories{}, course.CoursesCategoriesID).Error; err != nil {
		return course, errors.New("category not found, please create category first")
	}

	if err := c.db.Create(&course).Error; err != nil {
		return course, err
	}

	return course, nil
}

func (c *CoursesRepository) GetAll() ([]entity.Courses, error) {
	var courses []entity.Courses

	if err := c.db.Find(&courses).Error; err != nil {
		return courses, err
	}

	return courses, nil
}

func (c *CoursesRepository) GetByID(id int) (entity.Courses, error) {
	var course entity.Courses

	if err := c.db.First(&course, id).Error; err != nil {
		return course, err
	}

	return course, nil
}

func (c *CoursesRepository) Update(course entity.Courses) (entity.Courses, error) {
	if err := c.db.Save(&course).Error; err != nil {
		return course, err
	}

	return course, nil
}

func (c *CoursesRepository) Delete(id int) error {
	if err := c.db.Delete(&entity.Courses{}, id).Error; err != nil {
		return err
	}

	return nil
}
