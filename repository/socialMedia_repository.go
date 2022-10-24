package repository

import (
	"final-project-go/entity"

	"gorm.io/gorm"
)

type SocialMedia = entity.SocialMedia

type SocialMediaRepository interface {
	Create(SocialMedia) (SocialMedia, error)
	Save(SocialMedia) (SocialMedia, error)
	Delete(SocialMedia) (SocialMedia, error)
	GetOneByID(int) (SocialMedia, error)
	GetAll() ([]SocialMedia, error)
}

type socialMediaRepositoryImpl struct {
	DB *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepositoryImpl{DB: db}
}

func (r *socialMediaRepositoryImpl) Create(sosmed SocialMedia) (SocialMedia, error) {
	err := r.DB.Create(&sosmed).Error

	if err != nil {
		return sosmed, err
	}

	return sosmed, nil
}


func (r *socialMediaRepositoryImpl) Save(sosmed SocialMedia) (SocialMedia, error) {
	err := r.DB.Save(&sosmed).Error

	if err != nil {
		return sosmed, err
	}

	return sosmed, nil
}


func (r *socialMediaRepositoryImpl) Delete(sosmed SocialMedia) (SocialMedia, error) {

	err := r.DB.Delete(&sosmed).Error

	if err != nil {
		return sosmed, err
	}

	return sosmed, nil
}


func (r *socialMediaRepositoryImpl) GetOneByID(id int) (SocialMedia, error) {
	sosmed := SocialMedia{}
	err := r.DB.First(&sosmed, id).Error

	if err != nil {
		return sosmed, err
	}

	return sosmed, nil
}


func (r *socialMediaRepositoryImpl) GetAll() ([]SocialMedia, error) {
	
	var socialMedias []SocialMedia
	
	err := r.DB.Debug().Model(&SocialMedia{}).Preload("User").Find(&socialMedias).Error

	if err != nil {
		return socialMedias, err
	}
	
	return socialMedias, nil
}