package services

import (
	"final-project-go/entity"
	"final-project-go/repository"
)

type PhotoRepository = repository.PhotoRepository
type Photo = entity.Photo



type PhotoService interface {
	Create(Photo) (Photo, error)
	GetAll() ([]Photo, error)
	GetOne(int) (Photo, error)
	Update(Photo) (Photo, error) 
	Delete(Photo) (Photo, error) 
}

type photoServiceImpl struct {
	PhotoRepository PhotoRepository
}


func NewPhotoService(photoRepo *PhotoRepository) PhotoService {
	return &photoServiceImpl{PhotoRepository: *photoRepo}
}

func (service *photoServiceImpl) Create(p Photo) (Photo, error) {
	photo, err := service.PhotoRepository.CreatePhoto(p)

	if err != nil {
		return photo, err
	}

	return photo, nil
}


func (service *photoServiceImpl) GetAll() ([]Photo, error){
	photos, err := service.PhotoRepository.GetPhotos()

	if err != nil {
		return photos, err
	}

	return photos, nil

}



func (service *photoServiceImpl) GetOne(id int) (Photo, error) { 
	photo, err := service.PhotoRepository.GetPhotoByID(id)

	if err != nil {
		return photo, err
	}

	return photo, nil
}


func (service *photoServiceImpl) Update(p Photo) (Photo, error) { 
	photo, err := service.PhotoRepository.Save(p)

	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (service *photoServiceImpl) Delete(p Photo) (Photo, error) { 
	photo, err := service.PhotoRepository.DeletePhoto(p)

	if err != nil {
		return photo, err
	}

	return photo, nil
}