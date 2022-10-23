package repository

import (
	"final-project-go/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(entity.User) (entity.User, error)

	Update(*entity.User, int)  (entity.User, error)

	Delete(entity.User, int)  (entity.User, error)

	GetOne(entity.User) (entity.User, error)
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

func (repository *userRepositoryImpl) GetOne(user entity.User) (entity.User, error) {
	u := entity.User{}
	err := repository.DB.Where("email = ?", user.Email).Take(&u).Error
	
	if err != nil {
		return u, err
	}
	
	return u, nil
}


func (repository *userRepositoryImpl) Update(user *entity.User, id int) (entity.User, error) {
	
	if err := repository.DB.First(&user, id).Error; err != nil {
			return *user, err
	}

	if err := repository.DB.Save(&user).Error; err != nil {
		return *user, err
	}
	
	return *user, nil
}

func (repository *userRepositoryImpl) Delete(user entity.User, id int) (entity.User, error) {
	err := repository.DB.First(&user, id).Error
	
	if err != nil {
		return user, err
	}

	repository.DB.Debug().Delete(&user)

	return user, nil

}