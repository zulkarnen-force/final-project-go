package controllers

import (
	"final-project-go/dto"
	"final-project-go/entity"
	"final-project-go/helpers"
	"final-project-go/mappers"
	"final-project-go/services"
	"fmt"
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


// CreatePhoto godoc
// @Summary      Create Photo
// @Description   Create Photo user
// @Tags         Photos
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.ResponsePhoto
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /photos/ [post]
// @Security ApiKeyAuth
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

	photo, err := c.photoService.Create(photo)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "error create photo because " + err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, mappers.ResponsePhotoCreate(photo))

}

// GetPhoto godoc
// @Summary      Get Photos
// @Description   Get Photos user
// @Tags         Photos
// @Accept       json
// @Produce      json
// @Success      200  {object}  []dto.ResponseGetPhotos
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /photos/ [get]
// @Security ApiKeyAuth
func (c *PhotoController)  GetPhotos(ctx *gin.Context) {

	photos, err := c.photoService.GetAll()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: fmt.Sprintf("failed to get photos because %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, mappers.GetGetPhotoResponse(&photos))
}


// UpdatePhoto godoc
// @Summary      Update Photo User with ID
// @Description   Update Photo user
// @Param        id   path      int  true  "Photo ID"
// @Tags         Photos
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.PhotoResponseUpdate
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /photos/{id} [put]
// @Security ApiKeyAuth
func (c *PhotoController) UpdatePhotoByID(ctx *gin.Context) {
	
	var photo entity.Photo
	contentType := helpers.GetContentType(ctx)

	id, _ := ctx.Params.Get("photoId")
	photoID, _ := strconv.Atoi(id)

	photo, err := c.photoService.GetOne(photoID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: fmt.Sprintf("failed to update photos because %s", err.Error()),
		})
		return
	}

	if contentType == appJson {
		ctx.ShouldBindJSON(&photo)
	} else {
		ctx.ShouldBind(&photo)
	}

	photo, err = c.photoService.Update(photo)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: fmt.Sprintf("failed to update photos because %s", err.Error()),
		})
		return
	}
	
	ctx.JSON(http.StatusOK, mappers.GetUpdatePhotoResponse(photo))

}

// DeletePhoto godoc
// @Summary      Delete Photo User with ID
// @Description   Delete Photo user
// @Param        id   path      int  true  "Photo ID"
// @Tags         Photos
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.SuccessResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /photos/{id} [delete]
// @Security ApiKeyAuth
func (c *PhotoController) DeletePhoto(ctx *gin.Context) {
	stringId, _ := ctx.Params.Get("photoId")
	id, _ := strconv.Atoi(stringId)

	photo, err := c.photoService.GetOne(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: fmt.Sprintf("failed to delete photos because %s", err.Error()),
		})
		return
	}

	photo, err = c.photoService.Delete(photo)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: fmt.Sprintf("failed to delete photos because %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Your account has been successfully deleted",
	})
}