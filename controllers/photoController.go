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


func (c *Controller) CreatePhoto(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	userData := ctx.MustGet("userData").(jwt.MapClaims) // get info from JWT payload 
	id := int(userData["id"].(float64))

	var photo models.Photo = models.Photo{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&photo)
	} else {
		ctx.ShouldBind(&photo)
	}

	photo.UserID = id


	var user models.User

	c.DB.First(&user, id)

	photo.User = user


	if err := c.DB.Debug().Create(&photo).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"message":"gagal crated photo",
			"msg_dev":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":photo.ResponsePhotoCreate(),
	})
}


func (c *Controller)  GetPhotos(ctx *gin.Context) {

	var photos []models.Photo

	var photo models.Photo = models.Photo{}

	c.DB.Model(&models.Photo{}).Preload("User").Find(&photos)

	responsePhotos := photo.ResponseGetPhotos(&photos)

	ctx.JSON(http.StatusOK, gin.H{
		"data": responsePhotos,
	})
}


func (c *Controller) UpdatePhoto(ctx *gin.Context) {
	
	var photo models.Photo
	contentType := helpers.GetContentType(ctx)

	id, _ := ctx.Params.Get("photoId")
	photoID, _ := strconv.Atoi(id)

	if err := c.DB.First(&photo, photoID).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("user with id %s not found", id),
			"msg_dev": err.Error(),
		})
		return 		
	}

	if contentType == appJson {
		ctx.ShouldBindJSON(&photo)
	} else {
		ctx.ShouldBind(&photo)
	}

	if err := c.DB.Save(&photo).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message" : "failure updated data to database", 
			"msg_dev" :err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":"successfully updated data",
		"data":&photo,
	})

}


func (c *Controller) DeletePhoto(ctx *gin.Context) {
	stringId, _ := ctx.Params.Get("commentID")
	id, _ := strconv.Atoi(stringId)

	var Photo models.Photo

	err := c.DB.First(&Photo, id).Error 

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{
			"status":"not found",
			"msg_dev":err.Error(),
		})
		return
	}

	if err := c.DB.Delete(&Photo).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"status":"failed to deleted ",
			"msg_dev":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":"successfully deleted",
	})
}