package repository

import (
	"fmt"
	"hacktiv8-go-final-project/database"
	"hacktiv8-go-final-project/models"
	"time"
)

type CommentRepository interface {
	GetAllComments(*[]models.Comment) (*[]models.Comment, error)
	CreateComment(commentReq *models.Comment) (*models.Comment, error)
	DeleteComment(id uint) (*models.Comment, error)
	UpdateComment(commentReq *models.Comment, commentId int) (*models.Comment, error)
}

type commentRepositoryImpl struct{}

func NewCommentRepository() CommentRepository {
	return &commentRepositoryImpl{}
}

func (repository *commentRepositoryImpl) GetAllComments(*[]models.Comment) (*[]models.Comment, error) {
	db := database.GetDB()
	Comments := []models.Comment{}
	err := db.Preload("Photo").Preload("User").Find(&Comments).Error

	if err != nil {
		return nil, err
	}

	return &Comments, err
}

func (repository *commentRepositoryImpl) CreateComment(CommentReq *models.Comment) (*models.Comment, error) {
	var db = database.GetDB()
	fmt.Println(CommentReq)
	Comment := models.Comment{
		Message: CommentReq.Message,
		UserId:  CommentReq.UserId,
		PhotoId: CommentReq.PhotoId,
	}

	err := db.Create(&Comment).Error

	if err != nil {
		return nil, err
	}

	return &Comment, err
}

func (repository *commentRepositoryImpl) DeleteComment(id uint) (*models.Comment, error) {
	db := database.GetDB()
	Comment := models.Comment{}
	err := db.Delete(Comment, uint(id)).Error

	if err != nil {
		return nil, err
	}

	return &Comment, err
}

func (repository *commentRepositoryImpl) UpdateComment(commentReq *models.Comment, commentId int) (*models.Comment, error) {
	var db = database.GetDB()

	Comment := models.Comment{}

	err := db.First(&Comment, "id = ?", commentId).Error
	if err != nil {
		return nil, err
	}

	var updatedInput models.Comment
	updatedInput.Message = commentReq.Message
	updatedInput.UpdatedAt = time.Now()

	err_ := db.Model(&Comment).Updates(updatedInput).Error
	if err_ != nil {
		return nil, err
	}

	return &Comment, err
}
