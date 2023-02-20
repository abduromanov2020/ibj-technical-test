package usecase

import (
	"ibj-technical-test/entity"
	"ibj-technical-test/helpers"
	"ibj-technical-test/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/copier"
)

type IAdminUsecase interface {
	LoginAdmin(admin entity.LoginAdminRequest) (entity.LoginAdminResponse, error)
	CreateAdmin(admin entity.CreateAdminRequest) (entity.Admin, error)
	GetListAdmin() ([]entity.GetAdminResponse, error)
	GetAdminByID(id int) (entity.GetAdminResponse, error)
	UpdateAdminByID(id int, req entity.UpdateAdminRequest) (entity.Admin, error)
	DeleteAdminByID(id int) error
}

type AdminUsecase struct {
	adminRepository repository.IAdminRepository
}

func NewAdminUsecase(adminRepository repository.IAdminRepository) IAdminUsecase {
	return &AdminUsecase{adminRepository: adminRepository}
}

func (a *AdminUsecase) LoginAdmin(admin entity.LoginAdminRequest) (entity.LoginAdminResponse, error) {
	adminRequest := entity.Admin{}
	adminResponse := entity.LoginAdminResponse{}

	copier.Copy(&adminRequest, &admin)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = adminRequest.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		return adminResponse, err
	}

	adminResponse.Token = tokenString

	if _, err := a.adminRepository.Login(adminRequest.Email, adminRequest.Password); err != nil {
		return adminResponse, err
	}

	return adminResponse, nil
}

func (a *AdminUsecase) CreateAdmin(admin entity.CreateAdminRequest) (entity.Admin, error) {
	adminRequest := entity.Admin{}

	password, err := helpers.HashPassword(admin.Password)

	if err != nil {
		return entity.Admin{}, err
	}

	adminRequest.Name = admin.Name
	adminRequest.Password = password
	adminRequest.Email = admin.Email

	if _, err := a.adminRepository.Create(adminRequest); err != nil {
		return adminRequest, err
	}

	return adminRequest, nil
}

func (a *AdminUsecase) GetListAdmin() ([]entity.GetAdminResponse, error) {
	admins, err := a.adminRepository.GetAll()

	if err != nil {
		return nil, err
	}

	var adminsResponse []entity.GetAdminResponse
	copier.Copy(&adminsResponse, &admins)

	return adminsResponse, nil
}

func (a *AdminUsecase) GetAdminByID(id int) (entity.GetAdminResponse, error) {
	admin, err := a.adminRepository.GetByID(id)

	if err != nil {
		return entity.GetAdminResponse{}, err
	}

	var adminResponse entity.GetAdminResponse
	copier.Copy(&adminResponse, &admin)

	return adminResponse, nil
}

func (a *AdminUsecase) UpdateAdminByID(id int, req entity.UpdateAdminRequest) (entity.Admin, error) {
	admin, err := a.adminRepository.GetByID(id)

	if err != nil {
		return entity.Admin{}, err
	}

	copier.CopyWithOption(&admin, &req, copier.Option{IgnoreEmpty: true})

	if _, err := a.adminRepository.Update(admin); err != nil {
		return entity.Admin{}, err
	}

	return admin, nil
}

func (a *AdminUsecase) DeleteAdminByID(id int) error {
	if _, err := a.adminRepository.GetByID(id); err != nil {
		return err
	}

	if err := a.adminRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
