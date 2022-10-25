package services

import (
	"final-project-go/dto"
	"final-project-go/entity"
	"final-project-go/repository"
)

type UserResponseRegister = dto.UserResponseRegister
type UserResponseUpdate = dto.UserResponseUpdate
type User = entity.User


type UserService interface {
	Register(User) (User, error)
	Login(User) (User, error)
	Update(User, int) (User, error)
	Delete(User, int) (User, error)
	GetUserByID(int) (User, error)
}

type userServiceImpl struct {
	UserRepository repository.UserRepository
}


func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: *userRepository,
	}
}

func (service *userServiceImpl) Register(user entity.User) (User, error) {
	
	user, err := service.UserRepository.Create(user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (service *userServiceImpl) Login(user User) (User, error) {

	user, err := service.UserRepository.GetUser(user)

	if err != nil {
		return User{}, err
	}

	return user, nil
}


func (service *userServiceImpl) Update(user User, id int) (User, error) {
	
	usr, err := service.UserRepository.Update(user, id)

	if err != nil {
		return usr, err
	}

	return usr, nil
	
}


func (service *userServiceImpl) Delete(user User, id int) (User, error) {
	
	user, err := service.UserRepository.Delete(user, id)


	if err != nil {
		return user, err
	}

	return user, err

}




func (service *userServiceImpl) GetUserByID(id int) (User, error) {
	
	user, err := service.UserRepository.GetByID(id)

	if err != nil {
		return User{}, err
	}

	return user, nil

}
