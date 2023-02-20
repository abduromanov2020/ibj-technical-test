package usecase

import (
	"ibj-technical-test/entity"
	"ibj-technical-test/repository"

	"github.com/jinzhu/copier"
)

type IUserCoursesUsecase interface {
	CreateUserCourses(userCourses entity.CreateUserCourseRequest) (entity.UserCourses, error)
	GetListUserCourses() ([]entity.GetUserCourseResponse, error)
	GetUserCoursesByID(id int) (entity.GetUserCourseResponse, error)
	UpdateUserCoursesByID(id int, req entity.UpdateUserCourseRequest) (entity.UserCourses, error)
	DeleteUserCoursesByID(id int) error
}

type UserCoursesUsecase struct {
	UserCoursesRepository repository.IUserCoursesRepository
}

func NewUserCoursesUsecase(userCoursesRepository repository.IUserCoursesRepository) IUserCoursesUsecase {
	return &UserCoursesUsecase{UserCoursesRepository: userCoursesRepository}
}

func (u *UserCoursesUsecase) CreateUserCourses(userCourses entity.CreateUserCourseRequest) (entity.UserCourses, error) {
	userCoursesReq := entity.UserCourses{}

	copier.Copy(&userCoursesReq, &userCourses)

	if _, err := u.UserCoursesRepository.Create(userCoursesReq); err != nil {
		return userCoursesReq, err
	}

	return userCoursesReq, nil

}

func (u *UserCoursesUsecase) GetListUserCourses() ([]entity.GetUserCourseResponse, error) {
	userCourses, err := u.UserCoursesRepository.GetAll()

	if err != nil {
		return nil, err
	}

	var userCoursesResponse []entity.GetUserCourseResponse
	copier.Copy(&userCoursesResponse, &userCourses)

	return userCoursesResponse, nil
}

func (u *UserCoursesUsecase) GetUserCoursesByID(id int) (entity.GetUserCourseResponse, error) {
	userCourses, err := u.UserCoursesRepository.GetByID(id)

	if err != nil {
		return entity.GetUserCourseResponse{}, err
	}

	var userCoursesResponse entity.GetUserCourseResponse
	copier.Copy(&userCoursesResponse, &userCourses)

	return userCoursesResponse, nil
}

func (u *UserCoursesUsecase) UpdateUserCoursesByID(id int, req entity.UpdateUserCourseRequest) (entity.UserCourses, error) {
	userCourses, err := u.UserCoursesRepository.GetByID(id)

	if err != nil {
		return entity.UserCourses{}, err
	}

	copier.CopyWithOption(&userCourses, &req, copier.Option{IgnoreEmpty: true})

	if _, err := u.UserCoursesRepository.Update(userCourses); err != nil {
		return userCourses, err
	}

	return userCourses, nil
}

func (u *UserCoursesUsecase) DeleteUserCoursesByID(id int) error {

	if _, err := u.UserCoursesRepository.GetByID(id); err != nil {
		return err
	}

	if err := u.UserCoursesRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
