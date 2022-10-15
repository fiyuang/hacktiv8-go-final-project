package repository

import (
	"fmt"
	"hacktiv8-go-final-project/database"
	"hacktiv8-go-final-project/models"
	"time"
)

type PhotoRepository interface {
	GetAllPhotos(*[]models.Photo) (*[]models.Photo, error)
	CreatePhoto(PhotoReq *models.Photo) (*models.Photo, error)
	UpdatePhoto(PhotoReq *models.Photo, photoId int) (*models.Photo, error)
	DeletePhoto(id uint) (*models.Photo, error)
}

type photoRepositoryImpl struct{}

func NewPhotoRepository() PhotoRepository {
	return &photoRepositoryImpl{}
}

func (repository *photoRepositoryImpl) UpdatePhoto(PhotoReq *models.Photo, photoId int) (*models.Photo, error) {
	var db = database.GetDB()

	Photo := models.Photo{}

	err := db.First(&Photo, "id = ?", photoId).Error
	if err != nil {
		return nil, err
	}

	var updatedInput models.Photo
	updatedInput.Title = PhotoReq.Title
	updatedInput.Caption = PhotoReq.Caption
	updatedInput.PhotoUrl = PhotoReq.PhotoUrl
	updatedInput.UpdatedAt = time.Now()

	err_ := db.Model(&Photo).Updates(updatedInput).Error
	if err_ != nil {
		return nil, err
	}

	return &Photo, err
}

func (repository *photoRepositoryImpl) GetAllPhotos(*[]models.Photo) (*[]models.Photo, error) {
	db := database.GetDB()
	Photos := []models.Photo{}
	err := db.Preload("User").Find(&Photos).Error

	if err != nil {
		return nil, err
	}

	return &Photos, err
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
