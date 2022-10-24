package controllers

import (
	"final-project-go/entity"
	"final-project-go/helpers"
	"final-project-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

type PhotoService = services.PhotoService

type PhotoController struct {
	photoService services.PhotoService
}


func NewPhotoController(photoService *PhotoService) PhotoController {
	return PhotoController{photoService: *photoService}
}

func (c *PhotoController) CreatePhoto(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	userData := ctx.MustGet("userData").(jwt.MapClaims) // get info from JWT payload 
	id := int(userData["id"].(float64))

	var photo entity.Photo = entity.Photo{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&photo)
	} else {
		ctx.ShouldBind(&photo)
	}

	photo.UserID = id

	response, err := c.photoService.Create(photo)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusCreated, response)

}


func (c *PhotoController)  GetPhotos(ctx *gin.Context) {

	response, err := c.photoService.GetAll()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}


func (c *PhotoController) UpdatePhotoByID(ctx *gin.Context) {
	
	var photo entity.Photo
	contentType := helpers.GetContentType(ctx)

	id, _ := ctx.Params.Get("photoId")
	photoID, _ := strconv.Atoi(id)

	if contentType == appJson {
		ctx.ShouldBindJSON(&photo)
	} else {
		ctx.ShouldBind(&photo)
	}

	response, err := c.photoService.UpdatePhotoByID(photo, photoID)

	_ = err

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	
	ctx.JSON(http.StatusOK, response)

}


func (c *PhotoController) DeletePhoto(ctx *gin.Context) {
	stringId, _ := ctx.Params.Get("photoId")
	id, _ := strconv.Atoi(stringId)
	var photo entity.Photo

	response, err := c.photoService.DeletePhotoByID(photo, id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}