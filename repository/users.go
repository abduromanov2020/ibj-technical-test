package repository

import (
	"ibj-technical-test/entity"

	"gorm.io/gorm"
)

type IUsersRepository interface {
	Create(user entity.Users) (entity.Users, error)
	GetAll() ([]entity.Users, error)
	GetByID(id int) (entity.Users, error)
	Update(user entity.Users) (entity.Users, error)
	Delete(id int) error
}

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) IUsersRepository {
	return &UsersRepository{db: db}
}

func (u *UsersRepository) Create(user entity.Users) (entity.Users, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *UsersRepository) GetAll() ([]entity.Users, error) {
	var users []entity.Users

	if err := u.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (u *UsersRepository) GetByID(id int) (entity.Users, error) {
	var user entity.Users

	if err := u.db.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *UsersRepository) Update(user entity.Users) (entity.Users, error) {
	if err := u.db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *UsersRepository) Delete(id int) error {
	if err := u.db.Delete(&entity.Users{}, id).Error; err != nil {
		return err
	}

	return nil
}
