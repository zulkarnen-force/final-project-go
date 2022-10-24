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
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "error", MessageDev: err.Error()})
		return 
	}

	ctx.JSON(http.StatusCreated, mappers.GetResponseCreate(sosialMedia))
}


func (c *SocialMediaController)  GetSocialMedias(ctx *gin.Context) {

	sosialMedias, err := c.service.GetAll()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "error", MessageDev: err.Error()})
		return 
	}
	
	ctx.JSON(http.StatusOK, mappers.ToResponseMedias(&sosialMedias))
}


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
		
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message" : "failure updated data to database", 
			"msg_dev" :err.Error(),
		})
		
		return
	}

	ctx.JSON(http.StatusOK, mappers.GetSocialMediaUpdateResponse(socialMedia))

}


func (c *SocialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	paramId, _ := ctx.Params.Get("id")
	id, _ := strconv.Atoi(paramId)

	socialMedia, err := c.service.GetByID(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{
			"status":"not found",
			"msg_dev":err.Error(),
		})
		return
	}

	socialMedia, err = c.service.Delete(socialMedia)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"status":"failed to deleted ",
			"msg_dev":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":"Your social media has been successfully deleted",
	})
}