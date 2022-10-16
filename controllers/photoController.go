package controllers

import (
	"fmt"
	"hacktiv8-go-final-project/models"
	"hacktiv8-go-final-project/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
	result, err := controller.PhotoService.GetAllPhotos(&photos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (controller *photoControllerImpl) CreatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Validate input
	var input models.Photo
	input.UserId = userID
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create photo
	fmt.Println(&input)
	result, err := controller.PhotoService.CreatePhoto(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusCreated,
		"data":   result,
	})
}

func (controller *photoControllerImpl) DeletePhoto(c *gin.Context) {
	photoId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter False"})
		return
	}
	result, err := controller.PhotoService.DeletePhoto(uint(photoId))
	_ = result

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}

func (controller *photoControllerImpl) UpdatePhoto(c *gin.Context) {
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter False"})
		return
	}

	var input models.Photo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.PhotoService.UpdatePhoto(&input, photoId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	photoUpdate := models.PhotoUpdate{}
	photoUpdate.Id = result.Id
	photoUpdate.Caption = result.Caption
	photoUpdate.Title = result.Title
	photoUpdate.PhotoUrl = result.PhotoUrl
	photoUpdate.UserId = result.UserId
	photoUpdate.UpdatedAt = result.UpdatedAt

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   photoUpdate,
	})
}
