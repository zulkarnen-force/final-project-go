package mappers

import (
	"final-project-go/dto"
	"final-project-go/entity"

	"github.com/jinzhu/copier"
)

type SocialMedia = entity.SocialMedia

func GetResponseCreate(socialMedia SocialMedia) *dto.SocialMediaResponseCreate {
	response := new(dto.SocialMediaResponseCreate)
	copier.Copy(&response, &socialMedia)
	return response
}

func GetSocialMediaUpdateResponse(socialMedia SocialMedia) *dto.SocialMediaResponseUpdate {
	response := new(dto.SocialMediaResponseUpdate)
	copier.Copy(&response, &socialMedia)
	return response
}

func ToResponseMedias(socialMediaDatas *[]SocialMedia) *[]dto.SocialMediasResponse {
	var socialMedias []dto.SocialMediasResponse

	for _, data := range *socialMediaDatas {
		socialMediaResponse := new(dto.SocialMediasResponse)
		copier.Copy(&socialMediaResponse, &data)
		socialMedias = append(socialMedias, *socialMediaResponse)
	}

	return &socialMedias
}