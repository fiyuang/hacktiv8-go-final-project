package service

import (
	"hacktiv8-go-final-project/models"
	"hacktiv8-go-final-project/repository"
)

type CommentService interface {
	GetAllComments(comments *[]models.Comment) (*[]models.Comment, error)
	CreateComment(commentReq *models.Comment) (*models.Comment, error)
	DeleteComment(id uint) (*models.Comment, error)
	UpdateComment(commentReq *models.Comment, commentId int) (*models.Comment, error)
}

type commentServiceImpl struct {
	CommentRepository repository.CommentRepository
}

func NewCommentService(newCommentRepository repository.CommentRepository) CommentService {
	return &commentServiceImpl{
		CommentRepository: newCommentRepository,
	}
}

func (service *commentServiceImpl) GetAllComments(comments *[]models.Comment) (*[]models.Comment, error) {
	photo, err := service.CommentRepository.GetAllComments(comments)
	if err != nil {
		return nil, err
	}

	return photo, err
}

func (service *commentServiceImpl) CreateComment(commentReq *models.Comment) (*models.Comment, error) {
	comment, err := service.CommentRepository.CreateComment(commentReq)
	// fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return comment, err
}

func (service *commentServiceImpl) DeleteComment(id uint) (*models.Comment, error) {
	comment, err := service.CommentRepository.DeleteComment(id)
	// fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return comment, err
}

func (service *commentServiceImpl) UpdateComment(commentReq *models.Comment, photoId int) (*models.Comment, error) {
	photo, err := service.CommentRepository.UpdateComment(commentReq, photoId)
	if err != nil {
		return nil, err
	}

	return photo, err
}
