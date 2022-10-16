package repository

import (
	"hacktiv8-go-final-project/database"
	"hacktiv8-go-final-project/models"
	"time"
)

type SocialMediaRepository interface {
	GetAllSocialMedias(*[]models.SocialMedia) (*[]models.SocialMedia, error)
	CreateSocialMedia(socialmediaReq *models.SocialMedia) (*models.SocialMedia, error)
	UpdateSocialMedia(socialmediaReq *models.SocialMedia, socialmediaId int) (*models.SocialMedia, error)
	DeleteSocialMedia(id uint) (*models.SocialMedia, error)
}

type socialMediaRepositoryImpl struct{}

func NewSocialMediaRepository() SocialMediaRepository {
	return &socialMediaRepositoryImpl{}
}

func (repository *socialMediaRepositoryImpl) GetAllSocialMedias(*[]models.SocialMedia) (*[]models.SocialMedia, error) {
	db := database.GetDB()
	SocialMedia := []models.SocialMedia{}
	err := db.Preload("User").Find(&SocialMedia).Error

	if err != nil {
		return nil, err
	}

	return &SocialMedia, err
}

func (repository *socialMediaRepositoryImpl) UpdateSocialMedia(socialMediaReq *models.SocialMedia, socialmediaId int) (*models.SocialMedia, error) {
	var db = database.GetDB()

	SocialMedia := models.SocialMedia{}

	err := db.First(&SocialMedia, "id = ?", socialmediaId).Error
	if err != nil {
		return nil, err
	}

	var updatedInput models.SocialMedia
	updatedInput.Name = socialMediaReq.Name
	updatedInput.SocialMediaUrl = socialMediaReq.SocialMediaUrl
	updatedInput.UpdatedAt = time.Now()

	err_ := db.Model(&SocialMedia).Updates(updatedInput).Error
	if err_ != nil {
		return nil, err
	}

	return &SocialMedia, err
}

func (repository *socialMediaRepositoryImpl) CreateSocialMedia(socialmediaReq *models.SocialMedia) (*models.SocialMedia, error) {
	var db = database.GetDB()

	SocialMedia := models.SocialMedia{
		Name:           socialmediaReq.Name,
		SocialMediaUrl: socialmediaReq.SocialMediaUrl,
		UserId:         socialmediaReq.UserId,
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
