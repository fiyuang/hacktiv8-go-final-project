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
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Validate input
	var input models.Comment
	input.UserId = userID
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create comment
	result, err := controller.CommentService.CreateComment(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusCreated,
		"data":   result,
	})
}

func (controller *commentControllerImpl) DeleteComment(c *gin.Context) { // put authorization
	commentId, err := strconv.Atoi(c.Param("commentId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter False"})
		return
	}
	result, err := controller.CommentService.DeleteComment(uint(commentId))
	_ = result

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}

func (controller *commentControllerImpl) UpdateComment(c *gin.Context) {
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter False"})
		return
	}

	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.CommentService.UpdateComment(&input, commentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	commentUpdate := models.CommentUpdate{}
	commentUpdate.Id = result.Id
	commentUpdate.UserId = result.UserId
	commentUpdate.PhotoId = result.PhotoId
	commentUpdate.Message = result.Message
	commentUpdate.UpdatedAt = result.UpdatedAt

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   commentUpdate,
	})
}
