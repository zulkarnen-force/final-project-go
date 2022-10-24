package services

import (
	"final-project-go/entity"
	"final-project-go/repository"
)

type Comment = entity.Comment


type CommentService interface {
	Create(Comment)  (Comment, error)
	Update(Comment)  (Comment, error)
	GetByID(int)  (Comment, error)
	Delete(Comment)  (Comment, error)
	GetAll()  ([]Comment, error)
}

type commentServiceImpl struct {
	repository repository.CommentRepository
}

func NewCommentService(repository *repository.CommentRepository) CommentService {
	return &commentServiceImpl{repository: *repository}
}

func (service *commentServiceImpl) Create(c Comment)  (Comment, error) {
	comment, err := service.repository.Create(c)
	return comment, err
}



func (service *commentServiceImpl) Update(c Comment)  (Comment, error) {
	comment, err := service.repository.Save(c)
	return comment, err
}

func (service *commentServiceImpl) GetByID(id int)  (Comment, error) {
	comment, err := service.repository.GetOneByID(id)
	return comment, err
}


func (service *commentServiceImpl) Delete(c Comment)  (Comment, error) {
	comment, err := service.repository.Delete(c)
	return comment, err
}



func (service *commentServiceImpl) GetAll()  ([]Comment, error) {
	comments, err := service.repository.GetAll()
	return comments, err
}


