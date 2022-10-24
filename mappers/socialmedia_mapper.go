package mappers

import (
	"final-project-go/entity"
	"final-project-go/models"

	"github.com/jinzhu/copier"
)

type SocialMedia = entity.SocialMedia

func GetResponseCreate(socialMedia SocialMedia) *models.SocialMediaResponseCreate {
	response := new(models.SocialMediaResponseCreate)
	copier.Copy(&response, &socialMedia)
	return response
}

func GetSocialMediaUpdateResponse(socialMedia SocialMedia) *models.SocialMediaResponseUpdate {
	response := new(models.SocialMediaResponseUpdate)
	copier.Copy(&response, &socialMedia)
	return response
}

func ToResponseMedias(socialMediaDatas *[]SocialMedia) *[]models.SocialMediasResponse {
	var socialMedias []models.SocialMediasResponse

	for _, data := range *socialMediaDatas {
		socialMediaResponse := new(models.SocialMediasResponse)
		copier.Copy(&socialMediaResponse, &data)
		socialMedias = append(socialMedias, *socialMediaResponse)
	}

	return &socialMedias
}