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
	CreateComment(c *gin.Context)
	DeleteComment(c *gin.Context)
}

type commentControllerImpl struct {
	CommentService service.CommentService
}

func NewCommentController(newCommentService service.CommentService) CommentController {
	return &commentControllerImpl{
		CommentService: newCommentService,
	}
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
			//user not found
			return
		}
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

func (controller *commentControllerImpl) DeleteComment(c *gin.Context) {
	commentId, err := strconv.Atoi(c.Param("id"))

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
