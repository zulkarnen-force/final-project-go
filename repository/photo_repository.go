package repository

import (
	"final-project-go/entity"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	CreatePhoto(entity.Photo) (entity.Photo, error)
	DeletePhoto(entity.Photo)  (entity.Photo, error)
	GetPhotos() ([]entity.Photo, error)
	GetPhotoByID(int) (entity.Photo, error)
	Update(entity.Photo) (entity.Photo, error)
	Save(entity.Photo) (entity.Photo, error)
}


type photoRepositoryImpl struct {
	DB *gorm.DB
}


func NewPhotoRepository(database *gorm.DB) PhotoRepository {
	return &photoRepositoryImpl{
		DB: database,
	}
}

func (repository *photoRepositoryImpl) CreatePhoto(photo entity.Photo) (entity.Photo, error) {
	
	err := repository.DB.Create(&photo).Error

	if err != nil {
		return entity.Photo{}, err
	}

	return photo, nil
}



func (repository *photoRepositoryImpl) DeletePhoto(photo entity.Photo) (entity.Photo, error) {
	err := repository.DB.Delete(&photo).Error
	
	if err != nil {
		return photo, err
	}

	return photo, nil

}

func (repository *photoRepositoryImpl) GetPhotos() ([]entity.Photo, error) {
	var photos []entity.Photo
	
	err := repository.DB.Model(&entity.Photo{}).Preload("User").Find(&photos).Error
	
	if err != nil {
		return photos, err
	}

	return photos, nil
}


func (repository *photoRepositoryImpl) GetPhotoByID(id int) (entity.Photo, error) {
	var photo entity.Photo
	
	err := repository.DB.First(&photo, id).Error
	
	if err != nil {
		return photo, err
	}

	return photo, nil

}


func (repository *photoRepositoryImpl) Update(photo entity.Photo)   (entity.Photo, error)  {

	err := repository.DB.Save(&photo).Error

	if err != nil {
		return photo, err
	}

	return photo, nil
}


func (repository *photoRepositoryImpl) Save(photo entity.Photo)   (entity.Photo, error)  {

	err := repository.DB.Save(&photo).Error

	if err != nil {
		return photo, err
	}

	return photo, nil
}