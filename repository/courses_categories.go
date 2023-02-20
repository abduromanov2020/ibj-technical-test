package repository

import (
	"ibj-technical-test/entity"

	"gorm.io/gorm"
)

type ICoursesCategoriesRepository interface {
	Create(coursesCategories entity.CoursesCategories) (entity.CoursesCategories, error)
	GetAll() ([]entity.CoursesCategories, error)
	GetByID(id int) (entity.CoursesCategories, error)
	Update(coursesCategories entity.CoursesCategories) (entity.CoursesCategories, error)
	Delete(id int) error
}

type CoursesCategoriesRepository struct {
	db *gorm.DB
}

func NewCoursesCategoriesRepository(db *gorm.DB) ICoursesCategoriesRepository {
	return &CoursesCategoriesRepository{db: db}
}

func (c *CoursesCategoriesRepository) Create(coursesCategories entity.CoursesCategories) (entity.CoursesCategories, error) {
	if err := c.db.Create(&coursesCategories).Error; err != nil {
		return coursesCategories, err
	}

	return coursesCategories, nil
}

func (c *CoursesCategoriesRepository) GetAll() ([]entity.CoursesCategories, error) {
	var coursesCategories []entity.CoursesCategories

	if err := c.db.Find(&coursesCategories).Error; err != nil {
		return coursesCategories, err
	}

	return coursesCategories, nil
}

func (c *CoursesCategoriesRepository) GetByID(id int) (entity.CoursesCategories, error) {
	var coursesCategories entity.CoursesCategories

	if err := c.db.First(&coursesCategories, id).Error; err != nil {
		return coursesCategories, err
	}

	return coursesCategories, nil
}

func (c *CoursesCategoriesRepository) Update(coursesCategories entity.CoursesCategories) (entity.CoursesCategories, error) {
	if err := c.db.Save(&coursesCategories).Error; err != nil {
		return coursesCategories, err
	}

	return coursesCategories, nil
}

func (c *CoursesCategoriesRepository) Delete(id int) error {
	if err := c.db.Delete(&entity.CoursesCategories{}, id).Error; err != nil {
		return err
	}

	return nil
}
