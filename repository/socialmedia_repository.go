package repository

import (
	"fmt"
	"hacktiv8-go-final-project/database"
	"hacktiv8-go-final-project/models"
)

type SocialMediaRepository interface {
	// GetAllSocialMedias(*[]models.SocialMedia) (*[]models.SocialMedia, error)
	CreateSocialMedia(userReq *models.SocialMedia) (*models.SocialMedia, error)
	// UpdateSocialMedia(userReq *models.SocialMedia) (*models.SocialMedia, error)
	DeleteSocialMedia(id uint) (*models.SocialMedia, error)
}

type socialMediaRepositoryImpl struct{}

func NewSocialMediaRepository() SocialMediaRepository {
	return &socialMediaRepositoryImpl{}
}

func (repository *socialMediaRepositoryImpl) CreateSocialMedia(SocialMediaReq *models.SocialMedia) (*models.SocialMedia, error) {
	var db = database.GetDB()
	fmt.Println(SocialMediaReq)
	SocialMedia := models.SocialMedia{
		Name:           SocialMediaReq.Name,
		SocialMediaUrl: SocialMediaReq.SocialMediaUrl,
		UserId:         SocialMediaReq.UserId,
	}

	err := db.Create(&SocialMedia).Error

	if err != nil {
		return nil, err
	}

	return &SocialMedia, err
}

func (repository *socialMediaRepositoryImpl) DeleteSocialMedia(id uint) (*models.SocialMedia, error) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}
	err := db.Delete(SocialMedia, uint(id)).Error

	if err != nil {
		return nil, err
	}

	return &SocialMedia, err
}
