package repository

import (
	"fmt"
	"hacktiv8-go-final-project/database"
	"hacktiv8-go-final-project/models"
)

type PhotoRepository interface {
	// GetAllPhotos(*[]models.Photo) (*[]models.Photo, error)
	CreatePhoto(userReq *models.Photo) (*models.Photo, error)
	// UpdatePhoto(userReq *models.Photo) (*models.Photo, error)
	DeletePhoto(id uint) (*models.Photo, error)
}

type photoRepositoryImpl struct{}

func NewPhotoRepository() PhotoRepository {
	return &photoRepositoryImpl{}
}

func (repository *photoRepositoryImpl) CreatePhoto(PhotoReq *models.Photo) (*models.Photo, error) {
	var db = database.GetDB()
	fmt.Println(PhotoReq)
	Photo := models.Photo{
		Title:    PhotoReq.Title,
		Caption:  PhotoReq.Caption,
		PhotoUrl: PhotoReq.PhotoUrl,
		UserId:   PhotoReq.UserId,
	}

	err := db.Create(&Photo).Error

	if err != nil {
		return nil, err
	}

	return &Photo, err
}

func (repository *photoRepositoryImpl) DeletePhoto(id uint) (*models.Photo, error) {
	db := database.GetDB()
	Photo := models.Photo{}
	err := db.Delete(Photo, uint(id)).Error

	if err != nil {
		return nil, err
	}

	return &Photo, err
}
