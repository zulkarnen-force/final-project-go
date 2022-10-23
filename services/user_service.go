package services

import (
	"errors"
	"final-project-go/entity"
	"final-project-go/helpers"
	"final-project-go/mappers"
	"final-project-go/models"
	"final-project-go/repository"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type UserResponseRegister = models.UserResponseRegister
type UserResponseUpdate = models.UserResponseUpdate
type UserLoginResponse = models.UserLoginResponse

type UserService interface {
	Register(entity.User) (interface{}, error)
	Login(entity.User) (interface{}, error)
	Update(entity.User, int) (interface{}, error)
	Delete(entity.User, int) (interface{}, error)
}

type userServiceImpl struct {
	UserRepository repository.UserRepository
}


func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: *userRepository,
	}
}

func (service *userServiceImpl) Register(user entity.User) (interface{}, error) {
	user, err := service.UserRepository.Insert(user)

	if err != nil {
		return models.ErrorResponse{
			Message: "error register",
			MessageDev: err.Error(),
		}, err
	}

	response := mappers.GetResponseRegister(user)
	return response, nil
}

func (service *userServiceImpl) Login(user entity.User) (interface{}, error) {

	type FailLoginResponse struct {
		Error string 
		Message	string
	}
	user, err := service.UserRepository.GetOne(user)

	if err != nil {
		return FailLoginResponse{
			Error: err.Error(),
			Message: "invalid email or password",
		}, err
	}

	token := helpers.GenerateToken(user.ID, user.Email)

	response := mappers.GetUserLoginResponse(token)

	return response, nil
}

func (service *userServiceImpl) Update(user entity.User, id int) (interface{}, error) {
	
	userUpdated, err := service.UserRepository.Update(&user, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.ErrorResponse{
			Message: fmt.Sprintf("error because with %d record not found", id),
			MessageDev: err.Error(),
			Code: http.StatusNotFound,
		}, err
	}

	response := mappers.GetResponseUpdate(userUpdated)
	return response, nil
}

func (service *userServiceImpl) Delete(user entity.User, id int) (interface{}, error) {
	
	_, err := service.UserRepository.Delete(user, id)


	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse{
				Message: fmt.Sprintf("error because with %d record not found", id),
				MessageDev: err.Error(),
				Code: http.StatusNotFound,
			}, err
		} else {
			return models.ErrorResponse{
				Message: fmt.Sprintf("error delete user with %d", id),
				MessageDev: err.Error(),
				Code: http.StatusNotFound,
			}, err
		}
	}

	

	response := models.SuccessResponse{Message: "Your account has been deleted successfully"}
	return response, nil

}

