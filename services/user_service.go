package services

import (
	"final-project-go/entity"
	"final-project-go/mappers"
	"final-project-go/models"
	"final-project-go/repository"
)

type UserResponseRegister = models.UserResponseRegister

type UserService interface {
	Create(entity.User) (UserResponseRegister, error)
}

type userServiceImpl struct {
	UserRepository repository.UserRepository
}


func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: *userRepository,
	}
}

func (service *userServiceImpl) Create(user entity.User) (models.UserResponseRegister, error) {
	user, err := service.UserRepository.Insert(user)

	if err != nil {
		return models.UserResponseRegister{}, err
	}

	response := mappers.GetResponseRegister(user)
	return response, nil
}
