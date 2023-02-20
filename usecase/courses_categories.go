package usecase

import (
	"ibj-technical-test/entity"
	"ibj-technical-test/repository"

	"github.com/jinzhu/copier"
)

type ICoursesCategoriesUsecase interface {
	CreateCoursesCategories(coursesCategories entity.CreateCoursesCategoriesRequest) (entity.CoursesCategories, error)
	GetListCoursesCategories() ([]entity.GetCoursesCategoriesResponse, error)
	GetCoursesCategoriesByID(id int) (entity.GetCoursesCategoriesResponse, error)
	UpdateCoursesCategoriesByID(id int, req entity.UpdateCoursesCategoriesRequest) (entity.CoursesCategories, error)
	DeleteCoursesCategoriesByID(id int) error
}

type CoursesCategoriesUsecase struct {
	coursesCategoriesRepository repository.ICoursesCategoriesRepository
}

func NewCoursesCategoriesUsecase(coursesCategoriesRepository repository.ICoursesCategoriesRepository) ICoursesCategoriesUsecase {
	return &CoursesCategoriesUsecase{coursesCategoriesRepository: coursesCategoriesRepository}
}

func (c *CoursesCategoriesUsecase) CreateCoursesCategories(coursesCategories entity.CreateCoursesCategoriesRequest) (entity.CoursesCategories, error) {
	coursesCategoriesRequest := entity.CoursesCategories{}

	copier.Copy(&coursesCategoriesRequest, &coursesCategories)

	if _, err := c.coursesCategoriesRepository.Create(coursesCategoriesRequest); err != nil {
		return coursesCategoriesRequest, err
	}

	return coursesCategoriesRequest, nil
}

func (c *CoursesCategoriesUsecase) GetListCoursesCategories() ([]entity.GetCoursesCategoriesResponse, error) {
	coursesCategories, err := c.coursesCategoriesRepository.GetAll()

	if err != nil {
		return nil, err
	}

	var coursesCategoriesResponse []entity.GetCoursesCategoriesResponse
	copier.Copy(&coursesCategoriesResponse, &coursesCategories)

	return coursesCategoriesResponse, nil
}

func (c *CoursesCategoriesUsecase) GetCoursesCategoriesByID(id int) (entity.GetCoursesCategoriesResponse, error) {
	coursesCategories, err := c.coursesCategoriesRepository.GetByID(id)

	if err != nil {
		return entity.GetCoursesCategoriesResponse{}, err
	}

	var coursesCategoriesResponse entity.GetCoursesCategoriesResponse
	copier.Copy(&coursesCategoriesResponse, &coursesCategories)

	return coursesCategoriesResponse, nil
}

func (c *CoursesCategoriesUsecase) UpdateCoursesCategoriesByID(id int, req entity.UpdateCoursesCategoriesRequest) (entity.CoursesCategories, error) {
	coursesCategories, err := c.coursesCategoriesRepository.GetByID(id)

	if err != nil {
		return coursesCategories, err
	}

	copier.CopyWithOption(&coursesCategories, &req, copier.Option{IgnoreEmpty: true})

	if _, err := c.coursesCategoriesRepository.Update(coursesCategories); err != nil {
		return coursesCategories, err
	}

	return coursesCategories, nil
}

func (c *CoursesCategoriesUsecase) DeleteCoursesCategoriesByID(id int) error {

	if _, err := c.coursesCategoriesRepository.GetByID(id); err != nil {
		return err
	}

	if err := c.coursesCategoriesRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
