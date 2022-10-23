package mappers

import "final-project-go/models"

type ResponseGetPhotos = models.ResponseGetPhotos
type ResponsePhoto = models.ResponsePhoto
type PhotoResponseUpdate = models.PhotoResponseUpdate
type Photo = models.Photo

func ResponsePhotoCreate(p models.Photo) ResponsePhoto {
	return ResponsePhoto{
		ID:        p.ID,
		Title:     p.Title,
		Caption:   p.Caption,
		PhotoURL:  p.PhotoURL,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt,
	}
}

func GetGetPhotoResponse(photosData *[]Photo) *[]ResponseGetPhotos {

	var photos []ResponseGetPhotos = []ResponseGetPhotos{}

	photo := ResponseGetPhotos{}

	for _, data := range *photosData {
		photo.ID = data.ID
		photo.Title = data.Title
		photo.Caption = data.Caption
		photo.PhotoURL = data.PhotoURL
		photo.UserID = data.UserID
		photo.CreatedAt = data.CreatedAt
		photo.UpdatedAt = data.UpdatedAt

		photo.User.Email = data.User.Email
		photo.User.Username = data.User.Username

		photos = append(photos, photo)
	}

	return &photos

}

func GetUpdatePhotoResponse(p models.Photo) *PhotoResponseUpdate {

	return &PhotoResponseUpdate{
		ID:        p.ID,
		Title:     p.Title,
		Caption:   p.Caption,
		PhotoURL:  p.PhotoURL,
		UserID:    p.UserID,
		UpdatedAt: p.UpdatedAt,
	}

}