package services

import (
	"final-project-go/entity"
	"final-project-go/repository"
)

type SocialMedia = entity.SocialMedia

type SocialMediaService interface {
	Create(SocialMedia) (SocialMedia, error)
	Update(SocialMedia) (SocialMedia, error)
	GetByID(int) (SocialMedia, error)
	Delete(SocialMedia) (SocialMedia, error)
	GetAll() ([]SocialMedia, error)
}

type socialMediaServiceImpl struct {
	repository repository.SocialMediaRepository
} 

func NewSocialMediaService(repo *repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaServiceImpl{repository: *repo}
}



func (service *socialMediaServiceImpl) Create(sosmed SocialMedia)  (SocialMedia, error) {
	sosmed, err := service.repository.Create(sosmed)
	return sosmed, err
}



func (service *socialMediaServiceImpl) Update(sosmed SocialMedia)  (SocialMedia, error) {
	sosmed, err := service.repository.Save(sosmed)
	return sosmed, err
}

func (service *socialMediaServiceImpl) GetByID(id int)  (SocialMedia, error) {
	sosmed, err := service.repository.GetOneByID(id)
	return sosmed, err
}


func (service *socialMediaServiceImpl) Delete(sosmed SocialMedia)  (SocialMedia, error) {
	sosmed, err := service.repository.Delete(sosmed)
	return sosmed, err
}



func (service *socialMediaServiceImpl) GetAll()  ([]SocialMedia, error) {
	sosmeds, err := service.repository.GetAll()
	return sosmeds, err
}
