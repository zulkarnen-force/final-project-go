package controllers

import (
	"final-project-go/helpers"
	"final-project-go/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *Controller) CreatePhoto(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	var Photo models.Photo = models.Photo{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&Photo)
	} else {
		ctx.ShouldBind(&Photo)
	}


	if err := c.DB.Debug().Create(&Photo).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"message":"gagal crated photo",
			"msg_dev":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":&Photo,
	})
}


func (c *Controller)  GetPhotos(ctx *gin.Context) {

	// contentType := helpers.GetContentType(ctx)

	var photos []models.Photo

	// if contentType == appJson {
	// 	ctx.ShouldBindJSON(&User)
	// } else {
	// 	ctx.ShouldBind(&User)
	// }
	

	c.DB.Debug().Find(&photos)

	fmt.Println(photos)
	// err := c.DB.Debug().Create(&User).Error

	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error":"bad request",
	// 		"message":err.Error(),
	// 	})
	// 	return
	// }

	ctx.JSON(http.StatusOK, gin.H{
		"data": photos,
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
		"data": &photo,
	})

}


func (c *Controller) DeletePhoto(ctx *gin.Context) {
	// get ID using jsonData
	id, _ := ctx.Params.Get("photoId")
	photoID, _ := strconv.Atoi(id)

	var Photo models.Photo
	// check is exists?
	fmt.Println(id)

	err := c.DB.First(&Photo, photoID).Error 

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