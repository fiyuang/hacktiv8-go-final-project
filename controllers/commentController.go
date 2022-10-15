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

type CommentController interface {
	GetAllComments(c *gin.Context)
	CreateComment(c *gin.Context)
	DeleteComment(c *gin.Context)
	UpdateComment(c *gin.Context)
}

type commentControllerImpl struct {
	CommentService service.CommentService
}

func NewCommentController(newCommentService service.CommentService) CommentController {
	return &commentControllerImpl{
		CommentService: newCommentService,
	}
}

func (controller *commentControllerImpl) GetAllComments(c *gin.Context) {
	var comments []models.Comment
	res, err := controller.CommentService.GetAllComments(&comments)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (controller *commentControllerImpl) CreateComment(c *gin.Context) {
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	// Validate input
	var input models.Comment
	input.UserId = User.Id

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create comment
	fmt.Println(&input)
	res, err := controller.CommentService.CreateComment(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 201,
		"data":   res,
	})
}

func (controller *commentControllerImpl) DeleteComment(c *gin.Context) { // put authorization
	commentId, err := strconv.Atoi(c.Param("commentId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter False"})
		return
	}
	res, err := controller.CommentService.DeleteComment(uint(commentId))
	_ = res

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}

func (controller *commentControllerImpl) UpdateComment(c *gin.Context) {
	db := database.GetDB()

	commentId, err := strconv.Atoi(c.Param("commentId"))
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

	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := controller.CommentService.UpdateComment(&input, commentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	commentUpdate := models.CommentUpdate{}
	commentUpdate.Id = res.Id
	commentUpdate.UserId = res.UserId
	commentUpdate.PhotoId = res.PhotoId
	commentUpdate.Message = res.Message
	commentUpdate.UpdatedAt = res.UpdatedAt

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   commentUpdate,
	})
}
