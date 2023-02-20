package usecase

import (
	"ibj-technical-test/entity"
	"ibj-technical-test/repository"

	"github.com/jinzhu/copier"
)

type ICoursesUsecase interface {
	CreateCourse(course entity.CreateCourseRequest) (entity.Courses, error)
	GetListCourses() ([]entity.GetCourseResponse, error)
	GetCourseByID(id int) (entity.GetCourseResponse, error)
	UpdateCourseByID(id int, req entity.UpdateCourseRequest) (entity.Courses, error)
	DeleteCourseByID(id int) error
}

type CoursesUsecase struct {
	CoursesRepository repository.ICoursesRepository
}

func NewCoursesUsecase(coursesRepository repository.ICoursesRepository) ICoursesUsecase {
	return &CoursesUsecase{CoursesRepository: coursesRepository}
}

func (c *CoursesUsecase) CreateCourse(course entity.CreateCourseRequest) (entity.Courses, error) {
	courses := entity.Courses{}

	copier.Copy(&courses, &course)

	if _, err := c.CoursesRepository.Create(courses); err != nil {
		return courses, err
	}

	return courses, nil
}

func (c *CoursesUsecase) GetListCourses() ([]entity.GetCourseResponse, error) {
	courses, err := c.CoursesRepository.GetAll()

	if err != nil {
		return nil, err
	}

	var listCourses []entity.GetCourseResponse

	copier.Copy(&listCourses, &courses)

	return listCourses, nil
}

func (c *CoursesUsecase) GetCourseByID(id int) (entity.GetCourseResponse, error) {
	course, err := c.CoursesRepository.GetByID(id)

	if err != nil {
		return entity.GetCourseResponse{}, err
	}

	var getCourse entity.GetCourseResponse

	copier.Copy(&getCourse, &course)

	return getCourse, nil
}

func (c *CoursesUsecase) UpdateCourseByID(id int, req entity.UpdateCourseRequest) (entity.Courses, error) {
	course, err := c.CoursesRepository.GetByID(id)

	if err != nil {
		return course, err
	}

	copier.CopyWithOption(&course, &req, copier.Option{IgnoreEmpty: true})

	if _, err := c.CoursesRepository.Update(course); err != nil {
		return course, err
	}

	return course, nil
}

func (c *CoursesUsecase) DeleteCourseByID(id int) error {

	if _, err := c.CoursesRepository.GetByID(id); err != nil {
		return err
	}

	if err := c.CoursesRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
