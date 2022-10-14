package repository

import (
	"fmt"
	"hacktiv8-go-final-project/database"
	"hacktiv8-go-final-project/models"
)

type CommentRepository interface {
	CreateComment(userReq *models.Comment) (*models.Comment, error)
	DeleteComment(id uint) (*models.Comment, error)
}

type commentRepositoryImpl struct{}

func NewCommentRepository() CommentRepository {
	return &commentRepositoryImpl{}
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
