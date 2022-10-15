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

type PhotoController interface {
	GetAllPhotos(c *gin.Context)
	CreatePhoto(c *gin.Context)
	DeletePhoto(c *gin.Context)
	UpdatePhoto(c *gin.Context)
}

type photoControllerImpl struct {
	PhotoService service.PhotoService
}

func NewPhotoController(newPhotoService service.PhotoService) PhotoController {
	return &photoControllerImpl{
		PhotoService: newPhotoService,
	}
}

func (controller *photoControllerImpl) GetAllPhotos(c *gin.Context) {
	var photos []models.Photo
	res, err := controller.PhotoService.GetAllPhotos(&photos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (controller *photoControllerImpl) CreatePhoto(c *gin.Context) {
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
	var input models.Photo
	input.UserId = User.Id

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create photo
	fmt.Println(&input)
	res, err := controller.PhotoService.CreatePhoto(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusCreated,
		"data":   res,
	})
}

func (controller *photoControllerImpl) DeletePhoto(c *gin.Context) {
	photoId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter False"})
		return
	}
	res, err := controller.PhotoService.DeletePhoto(uint(photoId))
	_ = res

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}

func (controller *photoControllerImpl) UpdatePhoto(c *gin.Context) {
	db := database.GetDB()

	photoId, err := strconv.Atoi(c.Param("photoId"))
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

	var input models.Photo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := controller.PhotoService.UpdatePhoto(&input, photoId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	photoUpdate := models.PhotoUpdate{}
	photoUpdate.Id = res.Id
	photoUpdate.Caption = res.Caption
	photoUpdate.Title = res.Title
	photoUpdate.PhotoUrl = res.PhotoUrl
	photoUpdate.UserId = res.UserId
	photoUpdate.UpdatedAt = res.UpdatedAt

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   photoUpdate,
	})
}
