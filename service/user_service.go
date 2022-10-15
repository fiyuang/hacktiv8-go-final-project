package service

import (
	"hacktiv8-go-final-project/models"
	"hacktiv8-go-final-project/repository"
)

type UserService interface {
	GetUserById(id int) (*models.User, error)
	UserRegister(userReq *models.User) (*models.User, error)
	UserLogin(userReq *models.User) (*models.User, error)
	UpdateUser(userReq *models.User, userId int) (*models.User, error)
	DeleteUser(id uint) (*models.User, error)
}

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

//Inisiasi struct dengan kontrak interface
func NewUserService(newUserRepository repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: newUserRepository,
	}
}

func (service *userServiceImpl) GetUserById(id int) (*models.User, error) {
	user, err := service.UserRepository.GetUserById(id)
	// fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (service *userServiceImpl) UserRegister(userReq *models.User) (*models.User, error) {
	user, err := service.UserRepository.UserRegister(userReq)
	// fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (service *userServiceImpl) UserLogin(userReq *models.User) (*models.User, error) {
	user, err := service.UserRepository.UserLogin(userReq)
	// fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (service *userServiceImpl) UpdateUser(userReq *models.User, userId int) (*models.User, error) {
	user, err := service.UserRepository.UpdateUser(userReq, userId)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (service *userServiceImpl) DeleteUser(id uint) (*models.User, error) {
	user, err := service.UserRepository.DeleteUser(id)
	if err != nil {
		return nil, err
	}

	return user, err
}
