package controllers

import (
	"errors"
	"final-project-go/dto"
	"final-project-go/entity"
	"final-project-go/helpers"
	"final-project-go/mappers"
	"final-project-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type SocialMedia = entity.SocialMedia


type SocialMediaController struct {
	service services.SocialMediaService
}


func NewSocialMediaController(service *services.SocialMediaService) SocialMediaController {
	return SocialMediaController{service: *service}
}


// CreateSocialMedia godoc
// @Summary      Create Social Media User
// @Description   Creating social media user
// @Tags         Social Medias
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.SocialMediaResponseCreate
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /socialmedias/ [post]
// @Security ApiKeyAuth
func (c *SocialMediaController) CreateSocialMedia(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	userData := ctx.MustGet("userData").(jwt.MapClaims) // get info from JWT payload 
	id := int(userData["id"].(float64))

	var socialMedia SocialMedia
	socialMedia.UserID = id

	if contentType == appJson {
		ctx.ShouldBindJSON(&socialMedia)
	} else {
		ctx.ShouldBind(&socialMedia)
	}

	sosialMedia, err := c.service.Create(socialMedia)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "error get social media because " +  err.Error()})
		return 
	}

	ctx.JSON(http.StatusCreated, mappers.GetResponseCreate(sosialMedia))
}


// GetSocialMedia godoc
// @Summary      Get Social Media Users
// @Description   Get Social Media Users
// @Tags         Social Medias
// @Accept       json
// @Produce      json
// @Success      200  {object}  []dto.SocialMediasResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /socialmedias/ [get]
// @Security ApiKeyAuth
func (c *SocialMediaController)  GetSocialMedias(ctx *gin.Context) {

	sosialMedias, err := c.service.GetAll()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "error get social media because " +  err.Error()})
		return 
	}
	
	ctx.JSON(http.StatusOK, mappers.ToResponseMedias(&sosialMedias))
}


// UpdateSocialMedia godoc
// @Summary      Update Social Media  User
// @Description   Update Social with ID
// @Param        id   path      int  true  "Social Media ID"
// @Tags         Social Medias
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.SocialMediaResponseUpdate
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /socialmedias/{id}  [put]
// @Security ApiKeyAuth
func (c *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	
	contentType := helpers.GetContentType(ctx)

	paramId, _ := ctx.Params.Get("id")
	id, _ := strconv.Atoi(paramId)

	socialMedia, err := c.service.GetByID(id)
	
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, dto.ErrorResponse{Message: "social media not found"})
			return 
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
			return
		}
	}

	if contentType == appJson {
		ctx.ShouldBindJSON(&socialMedia)
	} else {
		ctx.ShouldBind(&socialMedia)
	}

	socialMedia, err = c.service.Update(socialMedia)

	if err != nil {
		
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "update social media failure because" + err.Error()})
		
		return
	}

	ctx.JSON(http.StatusOK, mappers.GetSocialMediaUpdateResponse(socialMedia))

}

// DeleteSocialMedia godoc
// @Summary      Delete Social Media User with ID
// @Description   Delete social media user
// @Param        id   path      int  true  "Social Media ID"
// @Tags         Social Medias
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.SuccessResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /socialmedias/{id} [delete]
// @Security ApiKeyAuth
func (c *SocialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	paramId, _ := ctx.Params.Get("id")
	id, _ := strconv.Atoi(paramId)

	socialMedia, err := c.service.GetByID(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, dto.ErrorResponse{Message: "social media not found"})
		return
	}

	socialMedia, err = c.service.Delete(socialMedia)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "delete social media failure"})
		return
	}

	ctx.JSON(http.StatusOK, dto.SuccessResponse{
		Message:"Your social media has been successfully deleted",
	})
}