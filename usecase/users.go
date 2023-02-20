package usecase

import (
	"ibj-technical-test/entity"
	"ibj-technical-test/repository"

	"github.com/jinzhu/copier"
)

type IUsersUsecase interface {
	CreateUser(user entity.CreateUserRequest) (entity.Users, error)
	GetListUsers() ([]entity.GetUserResponse, error)
	GetUserByID(id int) (entity.GetUserResponse, error)
	UpdateUserByID(id int, req entity.UpdateUserRequest) (entity.Users, error)
	DeleteUserByID(id int) error
}

type UsersUsecase struct {
	UsersRepository repository.IUsersRepository
}

func NewUsersUsecase(usersRepository repository.IUsersRepository) IUsersUsecase {
	return &UsersUsecase{UsersRepository: usersRepository}
}

func (u *UsersUsecase) CreateUser(user entity.CreateUserRequest) (entity.Users, error) {
	users := entity.Users{}

	copier.Copy(&users, &user)

	if _, err := u.UsersRepository.Create(users); err != nil {
		return users, err
	}

	return users, nil
}

func (u *UsersUsecase) GetListUsers() ([]entity.GetUserResponse, error) {
	users, err := u.UsersRepository.GetAll()

	if err != nil {
		return nil, err
	}

	var usersResponse []entity.GetUserResponse
	copier.Copy(&usersResponse, &users)

	return usersResponse, nil
}

func (u *UsersUsecase) GetUserByID(id int) (entity.GetUserResponse, error) {
	user, err := u.UsersRepository.GetByID(id)

	if err != nil {
		return entity.GetUserResponse{}, err
	}

	var userResponse entity.GetUserResponse
	copier.Copy(&userResponse, &user)

	return userResponse, nil
}

func (u *UsersUsecase) UpdateUserByID(id int, req entity.UpdateUserRequest) (entity.Users, error) {
	user, err := u.UsersRepository.GetByID(id)

	if err != nil {
		return entity.Users{}, err
	}

	copier.CopyWithOption(&user, &req, copier.Option{IgnoreEmpty: true})

	if _, err := u.UsersRepository.Update(user); err != nil {
		return user, err
	}

	return user, nil
}

func (u *UsersUsecase) DeleteUserByID(id int) error {
	if _, err := u.UsersRepository.GetByID(id); err != nil {
		return err
	}

	if err := u.UsersRepository.Delete(id); err != nil {
		return err
	}

	return nil

}
