package repository

import (
	"final-project-go/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(entity.User) (entity.User, error)

	Update(entity.User)  entity.User

	Delete(entity.User)
}


type userRepositoryImpl struct {
	DB *gorm.DB
}


func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: database,
	}
}

func (repository *userRepositoryImpl) Insert(user entity.User) (entity.User, error) {
	u := user
	err := repository.DB.Create(&u).Error

	if err != nil {
		return entity.User{}, err
	}

	return u, nil
}


func (repository *userRepositoryImpl) Update(user entity.User) entity.User {
	return entity.User{}
}

func (repository *userRepositoryImpl) Delete(user entity.User) {
	
}