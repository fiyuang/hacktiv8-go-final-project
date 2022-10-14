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
}

type socialMediaControllerImpl struct {
	SocialMediaService service.SocialMediaService
}

func NewSocialMediaController(newSocialMediaService service.SocialMediaService) SocialMediaController {
	return &socialMediaControllerImpl{
		SocialMediaService: newSocialMediaService,
	}
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
		"status": 201,
		"data":   res,
	})
}

func (controller *socialMediaControllerImpl) DeleteSocialMedia(c *gin.Context) {
	socialMediaId, err := strconv.Atoi(c.Param("id"))

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
