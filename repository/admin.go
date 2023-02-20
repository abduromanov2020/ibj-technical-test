package repository

import (
	"ibj-technical-test/entity"
	"ibj-technical-test/helpers"

	"gorm.io/gorm"
)

type IAdminRepository interface {
	Login(email string, password string) (entity.Admin, error)
	Create(admin entity.Admin) (entity.Admin, error)
	GetAll() ([]entity.Admin, error)
	GetByID(id int) (entity.Admin, error)
	Update(admin entity.Admin) (entity.Admin, error)
	Delete(id int) error
}

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) IAdminRepository {
	return &AdminRepository{db: db}
}

func (a *AdminRepository) Login(email string, password string) (entity.Admin, error) {
	var admin entity.Admin

	if err := a.db.Where("email = ?", email).First(&admin).Error; err != nil {
		return admin, err
	}

	if match, err := helpers.CheckPasswordHash(password, admin.Password); err != nil || !match {
		return admin, err
	}

	return admin, nil
}

func (a *AdminRepository) Create(admin entity.Admin) (entity.Admin, error) {

	if err := a.db.Create(&admin).Error; err != nil {
		return admin, err
	}

	return admin, nil
}

func (a *AdminRepository) GetAll() ([]entity.Admin, error) {
	var admins []entity.Admin

	if err := a.db.Find(&admins).Error; err != nil {
		return admins, err
	}

	return admins, nil
}

func (a *AdminRepository) GetByID(id int) (entity.Admin, error) {
	var admin entity.Admin

	if err := a.db.First(&admin, id).Error; err != nil {
		return admin, err
	}

	return admin, nil
}

func (a *AdminRepository) Update(admin entity.Admin) (entity.Admin, error) {
	if err := a.db.Save(&admin).Error; err != nil {
		return admin, err
	}

	return admin, nil
}

func (a *AdminRepository) Delete(id int) error {
	if err := a.db.Delete(&entity.Admin{}, id).Error; err != nil {
		return err
	}

	return nil
}
