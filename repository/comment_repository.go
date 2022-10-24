package repository

import (
	"final-project-go/entity"

	"gorm.io/gorm"
)


type Comment = entity.Comment

type CommentRepository interface {
	Create(Comment) (Comment, error)
	Save(Comment) (Comment, error)
	Delete(Comment) (Comment, error)
	GetOneByID(int) (Comment, error)
	GetAll() ([]Comment, error)
}


type commentRepositoryImpl struct {
	DB *gorm.DB
}


func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (r *commentRepositoryImpl) Create(comment Comment) (Comment, error){
	err := r.DB.Create(&comment).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *commentRepositoryImpl) Save(comment Comment) (Comment, error){
	err := r.DB.Save(&comment).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *commentRepositoryImpl) Delete(comment Comment) (Comment, error){
	err := r.DB.Delete(&comment).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *commentRepositoryImpl) GetOneByID(id int) (Comment, error){
	comment := Comment{}
	err := r.DB.First(&comment, id).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *commentRepositoryImpl) GetAll() ([]Comment, error){
	comments := []Comment{}

	r.DB.Debug().Model(&entity.Comment{}).Preload("User", func (db *gorm.DB) *gorm.DB {
		return	db.Select("id", "email", "username")
	}).Preload("Photo.User").Find(&comments)

	// if err != nil {
	// 	return comments, err
	// }

	return comments, nil
}