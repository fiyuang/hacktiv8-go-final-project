package controllers

import (
	"hacktiv8-go-final-project/models"
	"hacktiv8-go-final-project/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type SocialMediaController interface {
	CreateSocialMedia(c *gin.Context)
	DeleteSocialMedia(c *gin.Context)
	UpdateSocialMedia(c *gin.Context)
	GetAllSocialMedias(c *gin.Context)
}

type socialMediaControllerImpl struct {
	SocialMediaService service.SocialMediaService
}

func NewSocialMediaController(newSocialMediaService service.SocialMediaService) SocialMediaController {
	return &socialMediaControllerImpl{
		SocialMediaService: newSocialMediaService,
	}
}

func (controller *socialMediaControllerImpl) GetAllSocialMedias(c *gin.Context) {
	var socialMedia []models.SocialMedia
	result, err := controller.SocialMediaService.GetAllSocialMedias(&socialMedia)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (controller *socialMediaControllerImpl) UpdateSocialMedia(c *gin.Context) {
	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter False"})
		return
	}

	var input models.SocialMedia
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.SocialMediaService.UpdateSocialMedia(&input, socialMediaId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	socialmediaUpdate := models.SocialMediaUpdate{}
	socialmediaUpdate.Id = result.Id
	socialmediaUpdate.Name = result.Name
	socialmediaUpdate.SocialMediaUrl = result.SocialMediaUrl
	socialmediaUpdate.UserId = result.UserId
	socialmediaUpdate.UpdatedAt = result.UpdatedAt

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   socialmediaUpdate,
	})
}

func (controller *socialMediaControllerImpl) CreateSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Validate input
	var input models.SocialMedia
	input.UserId = userID
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create socialMedia
	result, err := controller.SocialMediaService.CreateSocialMedia(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusCreated,
		"data":   result,
	})
}

func (controller *socialMediaControllerImpl) DeleteSocialMedia(c *gin.Context) {
	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter False"})
		return
	}

	result, err := controller.SocialMediaService.DeleteSocialMedia(uint(socialMediaId))
	_ = result
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
