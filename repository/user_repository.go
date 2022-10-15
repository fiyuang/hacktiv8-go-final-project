package repository

import (
	"errors"
	"hacktiv8-go-final-project/database"
	"hacktiv8-go-final-project/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserById(id int) (*models.User, error)
	UserRegister(userReq *models.User) (*models.User, error)
	UserLogin(userReq *models.User) (*models.User, error)
	UpdateUser(userReq *models.User, userId int) (*models.User, error)
	DeleteUser(id uint) (*models.User, error)
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

func (repository *userRepositoryImpl) UpdateUser(UserReq *models.User, userId int) (*models.User, error) {
	var db = database.GetDB()

	User := models.User{}

	err := db.First(&User, "id = ?", userId).Error
	if err != nil {
		return nil, err
	}

	var updatedInput models.User
	updatedInput.Username = UserReq.Username
	updatedInput.Email = UserReq.Email
	updatedInput.UpdatedAt = time.Now()

	err_ := db.Model(&User).Updates(updatedInput).Error
	if err_ != nil {
		return nil, err
	}

	return &User, err
}

func (repository *userRepositoryImpl) DeleteUser(id uint) (*models.User, error) {
	db := database.GetDB()
	User := models.User{}
	err := db.Delete(User, uint(id)).Error

	if err != nil {
		return nil, err
	}

	return &User, err
}
