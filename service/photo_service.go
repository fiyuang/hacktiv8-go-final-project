package service

import (
	"hacktiv8-go-final-project/models"
	"hacktiv8-go-final-project/repository"
)

type PhotoService interface {
	GetAllPhotos(photos *[]models.Photo) (*[]models.Photo, error)
	CreatePhoto(photoReq *models.Photo) (*models.Photo, error)
	UpdatePhoto(photoReq *models.Photo, photoId int) (*models.Photo, error)
	DeletePhoto(id uint) (*models.Photo, error)
}

type photoServiceImpl struct {
	PhotoRepository repository.PhotoRepository
}

func NewPhotoService(newPhotoRepository repository.PhotoRepository) PhotoService {
	return &photoServiceImpl{
		PhotoRepository: newPhotoRepository,
	}
}

func (service *photoServiceImpl) GetAllPhotos(photos *[]models.Photo) (*[]models.Photo, error) {
	photo, err := service.PhotoRepository.GetAllPhotos(photos)
	if err != nil {
		return nil, err
	}

	return photo, err
}

func (service *photoServiceImpl) CreatePhoto(photoReq *models.Photo) (*models.Photo, error) {
	photo, err := service.PhotoRepository.CreatePhoto(photoReq)
	// fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return photo, err
}

func (service *photoServiceImpl) DeletePhoto(id uint) (*models.Photo, error) {
	photo, err := service.PhotoRepository.DeletePhoto(id)
	// fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return photo, err
}

func (service *photoServiceImpl) UpdatePhoto(photoReq *models.Photo, photoId int) (*models.Photo, error) {
	photo, err := service.PhotoRepository.UpdatePhoto(photoReq, photoId)
	if err != nil {
		return nil, err
	}

	return photo, err
}
