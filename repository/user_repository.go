package repository

import (
	"errors"
	"hacktiv8-go-final-project/database"
	"hacktiv8-go-final-project/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserById(id int) (*models.User, error)
	UserRegister(userReq *models.User) (*models.User, error)
	UserLogin(userReq *models.User) (*models.User, error)
}

type userRepositoryImpl struct{}

//Inisiasi struct dengan kontrak interface
func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (repository *userRepositoryImpl) GetUserById(id int) (*models.User, error) {
	var db = database.GetDB()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//user not found
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}

func (repository *userRepositoryImpl) UserRegister(UserReq *models.User) (*models.User, error) {
	var db = database.GetDB()

	User := models.User{
		Username: UserReq.Username,
		Email:    UserReq.Email,
		Password: UserReq.Password,
		Age:      UserReq.Age,
	}

	err := db.Create(&User).Error

	if err != nil {
		return nil, err
	}

	return &User, err

}

func (repository *userRepositoryImpl) UserLogin(UserReq *models.User) (*models.User, error) {
	var db = database.GetDB()

	user := models.User{}

	err := db.First(&user, "email = ?", UserReq.Email).Take(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, err

}
