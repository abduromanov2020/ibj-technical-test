package repository

import (
	"errors"
	"ibj-technical-test/entity"

	"gorm.io/gorm"
)

type IUserCoursesRepository interface {
	Create(userCourse entity.UserCourses) (entity.UserCourses, error)
	GetAll() ([]entity.UserCourses, error)
	GetByID(id int) (entity.UserCourses, error)
	Update(userCourse entity.UserCourses) (entity.UserCourses, error)
	Delete(id int) error
}

type UserCoursesRepository struct {
	db *gorm.DB
}

func NewUserCoursesRepository(db *gorm.DB) IUserCoursesRepository {
	return &UserCoursesRepository{db: db}
}

func (u *UserCoursesRepository) Create(userCourse entity.UserCourses) (entity.UserCourses, error) {

	errUserID := u.db.First(&userCourse.Users, userCourse.UsersID).Error
	errCoursesID := u.db.First(&userCourse.Courses, userCourse.CoursesID).Error

	if errUserID != nil || errCoursesID != nil {
		return userCourse, errors.New("user or course not found")
	}

	if err := u.db.Create(&userCourse).Error; err != nil {
		return userCourse, err
	}

	return userCourse, nil
}

func (u *UserCoursesRepository) GetAll() ([]entity.UserCourses, error) {
	var userCourses []entity.UserCourses

	if err := u.db.Find(&userCourses).Error; err != nil {
		return userCourses, err
	}

	return userCourses, nil
}

func (u *UserCoursesRepository) GetByID(id int) (entity.UserCourses, error) {
	var userCourse entity.UserCourses

	if err := u.db.First(&userCourse, id).Error; err != nil {
		return userCourse, err
	}

	return userCourse, nil
}

func (u *UserCoursesRepository) Update(userCourse entity.UserCourses) (entity.UserCourses, error) {
	if err := u.db.Save(&userCourse).Error; err != nil {
		return userCourse, err
	}

	return userCourse, nil
}

func (u *UserCoursesRepository) Delete(id int) error {
	if err := u.db.Delete(&entity.UserCourses{}, id).Error; err != nil {
		return err
	}

	return nil
}
