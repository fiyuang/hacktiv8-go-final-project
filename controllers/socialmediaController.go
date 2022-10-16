package controllers

import (
	"errors"
	"fmt"
	"hacktiv8-go-final-project/database"
	"hacktiv8-go-final-project/helpers"
	"hacktiv8-go-final-project/models"
	"hacktiv8-go-final-project/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	var photos []models.SocialMedia
	res, err := controller.SocialMediaService.GetAllSocialMedias(&photos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (controller *socialMediaControllerImpl) UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()

	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter False"})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userEmail := userData["email"].(string)

	contentType := helpers.GetContentType(c)
	_, _, _ = db, contentType, userEmail

	// Photo := models.Photo{}
	// error := db.First(&Photo, "email = ?", userEmail).Error
	// if error != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Photo not found!"})
	// 		return
	// 	}
	// 	return
	// }

	var input models.SocialMedia
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := controller.SocialMediaService.UpdateSocialMedia(&input, socialMediaId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	socialmediaUpdate := models.SocialMediaUpdate{}
	socialmediaUpdate.Id = res.Id
	socialmediaUpdate.Name = res.Name
	socialmediaUpdate.SocialMediaUrl = res.SocialMediaUrl
	socialmediaUpdate.UserId = res.UserId
	socialmediaUpdate.UpdatedAt = res.UpdatedAt

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   socialmediaUpdate,
	})
}

func (controller *socialMediaControllerImpl) CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	// userID := uint(userData["id"].(float64))
	userEmail := userData["email"].(string)

	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	User := models.User{}
	err := db.First(&User, "email = ?", userEmail).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//user not found
			return
		}
		return
	}

	// Validate input
	var input models.SocialMedia
	input.UserId = User.Id

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create socialMedia
	fmt.Println(&input)
	res, err := controller.SocialMediaService.CreateSocialMedia(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusCreated,
		"data":   res,
	})
}

func (controller *socialMediaControllerImpl) DeleteSocialMedia(c *gin.Context) {
	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter False"})
		return
	}
	res, err := controller.SocialMediaService.DeleteSocialMedia(uint(socialMediaId))
	_ = res

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your socialMedia has been successfully deleted",
	})
}
