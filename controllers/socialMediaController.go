package controllers

import (
	"final-project-go/helpers"
	"final-project-go/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)


func (c *Controller) CreateSocialMedia(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	userData := ctx.MustGet("userData").(jwt.MapClaims) // get info from JWT payload 
	id := int(userData["id"].(float64))

	var socialMedia models.SocialMedia
	socialMedia.UserID = id

	if contentType == appJson {
		ctx.ShouldBindJSON(&socialMedia)
	} else {
		ctx.ShouldBind(&socialMedia)
	}

	c.DB.Create(&socialMedia)

	ctx.JSON(http.StatusCreated, socialMedia.GetResponseCreate())
}


func (c *Controller)  GetSocialMedias(ctx *gin.Context) {

	var socialMedias *[]models.SocialMedia
	
	c.DB.Debug().Model(&models.SocialMedia{}).Preload("User").Find(&socialMedias)
	
	ctx.JSON(http.StatusOK, models.SocialMedia{}.ToResponseMedias(socialMedias))
}


func (c *Controller) UpdateSocialMedia(ctx *gin.Context) {
	
	var socialMedia models.SocialMedia = models.SocialMedia{}
	contentType := helpers.GetContentType(ctx)

	paramId, _ := ctx.Params.Get("id")
	id, _ := strconv.Atoi(paramId)

	if err := c.DB.First(&socialMedia, id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("user with id %d not found", id),
			"msg_dev": err.Error(),
		})
		return 		
	}

	if contentType == appJson {
		ctx.ShouldBindJSON(&socialMedia)
	} else {
		ctx.ShouldBind(&socialMedia)
	}


	if err := c.DB.Save(&socialMedia).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message" : "failure updated data to database", 
			"msg_dev" :err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMedia.GetResponseUpdate())

}


func (c *Controller) DeleteSocialMedia(ctx *gin.Context) {
	paramId, _ := ctx.Params.Get("id")
	id, _ := strconv.Atoi(paramId)

	var socialMedia models.SocialMedia

	err := c.DB.First(&socialMedia, id).Error 

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{
			"status":"not found",
			"msg_dev":err.Error(),
		})
		return
	}

	if err := c.DB.Delete(&socialMedia).Error; err != nil {
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