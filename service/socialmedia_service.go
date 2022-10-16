package service

import (
	"hacktiv8-go-final-project/models"
	"hacktiv8-go-final-project/repository"
)

type SocialMediaService interface {
	GetAllSocialMedias(*[]models.SocialMedia) (*[]models.SocialMedia, error)
	CreateSocialMedia(socialMediaReq *models.SocialMedia) (*models.SocialMedia, error)
	UpdateSocialMedia(socialMediaReq *models.SocialMedia, socialmediaId int) (*models.SocialMedia, error)
	DeleteSocialMedia(id uint) (*models.SocialMedia, error)
}

type socialMediaServiceImpl struct {
	SocialMediaRepository repository.SocialMediaRepository
}

func NewSocialMediaService(newSocialMediaRepository repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaServiceImpl{
		SocialMediaRepository: newSocialMediaRepository,
	}
}

func (service *socialMediaServiceImpl) GetAllSocialMedias(socialmedias *[]models.SocialMedia) (*[]models.SocialMedia, error) {
	socialMedia, err := service.SocialMediaRepository.GetAllSocialMedias(socialmedias)
	if err != nil {
		return nil, err
	}

	return socialMedia, err
}

func (service *socialMediaServiceImpl) UpdateSocialMedia(socialMediaReq *models.SocialMedia, socialmediaId int) (*models.SocialMedia, error) {
	socialMedia, err := service.SocialMediaRepository.UpdateSocialMedia(socialMediaReq, socialmediaId)
	if err != nil {
		return nil, err
	}

	return socialMedia, err
}

func (service *socialMediaServiceImpl) CreateSocialMedia(socialMediaReq *models.SocialMedia) (*models.SocialMedia, error) {
	socialMedia, err := service.SocialMediaRepository.CreateSocialMedia(socialMediaReq)
	// fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return socialMedia, err
}

func (service *socialMediaServiceImpl) DeleteSocialMedia(id uint) (*models.SocialMedia, error) {
	socialMedia, err := service.SocialMediaRepository.DeleteSocialMedia(id)
	// fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return socialMedia, err
}
