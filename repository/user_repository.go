package repository

import (
	"errors"
	"final-project-go/entity"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(entity.User) (entity.User, error)
	Update(entity.User, int)  (entity.User, error)
	Delete(entity.User, int)  (entity.User, error)
	GetUser(entity.User) (entity.User, error)
	GetByID(int) (entity.User, error) 
}


type userRepositoryImpl struct {
	DB *gorm.DB
}


func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: database,
	}
}

var DuplicateError = errors.New("email atau username telah digunakan")

func (repository *userRepositoryImpl) Create(user entity.User) (entity.User, error) {
	u := user
 	err := repository.DB.Create(&u).Error

	if err != nil {
		if e, ok := err.(*pgconn.PgError); ok {
			if e.Code == pgerrcode.UniqueViolation {
				return u, DuplicateError
			}
		}
		return u, err
	}

	
	return u, nil
}

func (repository *userRepositoryImpl) GetUser(user entity.User) (entity.User, error) {
	u := entity.User{}
	err := repository.DB.Where("email = ?", user.Email).Take(&u).Error
	
	if err != nil {
		return u, err
	}
	
	return u, nil
}


func (repository *userRepositoryImpl) Update(user entity.User, id int) (entity.User, error) {

	err := repository.DB.First(&entity.User{}, id).Updates(user).Error

	if err != nil {
		return user, err
	}

	return user, nil

}

func (repository *userRepositoryImpl) Delete(user entity.User, id int) (entity.User, error) {
	err := repository.DB.First(&user, id).Error
	
	if err != nil {
		return user, err
	}

	repository.DB.Debug().Delete(&user)

	return user, nil

}


func (repository *userRepositoryImpl) GetByID(id int) (entity.User, error) {
	usr := entity.User{}

	err := repository.DB.First(&usr, id).Error
	
	if err != nil {
		return usr, err
	}

	return usr, nil

}