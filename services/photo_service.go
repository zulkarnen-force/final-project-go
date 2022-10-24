package services

import (
	"final-project-go/dto"
	"final-project-go/entity"
	"final-project-go/mappers"
	"final-project-go/repository"
)

type PhotoRepository = repository.PhotoRepository
type Photo = entity.Photo



type PhotoService interface {
	Create(Photo) (interface{}, error)
	UpdatePhotoByID(Photo, int) (interface{}, error)
	DeletePhotoByID(Photo, int) (interface{}, error)
	GetAll() (interface{}, error)
	GetOne(int)
}

type photoServiceImpl struct {
	PhotoRepository PhotoRepository
}


func NewPhotoService(photoRepo *PhotoRepository) PhotoService {
	return &photoServiceImpl{PhotoRepository: *photoRepo}
}

func (service *photoServiceImpl) Create(p Photo) (interface{}, error) {
	photo, err := service.PhotoRepository.CreatePhoto(p)

	if err != nil {
		return map[string]interface{}{"message":"error", "dev":err.Error()}, err
	}

	return mappers.ResponsePhotoCreate(photo), nil
}

func (service *photoServiceImpl) UpdatePhotoByID(photo Photo, id int)  (interface{}, error) {
	p := photo
	p, err := service.PhotoRepository.GetPhotoByID(id)

	if err != nil {
		return dto.ErrorResponse{
			Message: "error updated photo because " + err.Error(),
			MessageDev: err.Error(),
		}, err
	}

	photo, err = service.PhotoRepository.Update(p)

	if err != nil {
		return dto.ErrorResponse{
			Message: "error updated photo because " + err.Error(),
			MessageDev: err.Error(),
		}, err
	}

	return mappers.GetUpdatePhotoResponse(photo), nil
}

func (service *photoServiceImpl) DeletePhotoByID(p Photo, id int) (interface{}, error) {

	photo, err := service.PhotoRepository.DeletePhotoByID(id)

	if err != nil {
		return dto.ErrorResponse{Message:"tidak bisa deleted", MessageDev:err.Error()}, err
	}

	return photo, nil

}


func (service *photoServiceImpl) GetAll() (interface{}, error){
	photos, err := service.PhotoRepository.GetPhotos()

	if err != nil {
		return map[string]interface{}{"message":"not data photos"}, err
	}

	res := mappers.GetGetPhotoResponse(&photos)
	return res, nil

}

func (service *photoServiceImpl) GetOne(id int) {
	
}

// func (service *photoServiceImpl) UpdatePhoto(p Photo) {
// 	res, photo := service.PhotoRepository.Update(p)
// }