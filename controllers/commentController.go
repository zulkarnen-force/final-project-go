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


func (c *Controller) CreateComment(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	userData := ctx.MustGet("userData").(jwt.MapClaims) // get info from JWT payload 
	id := int(userData["id"].(float64))

	var comment models.Comment = models.Comment{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&comment)
	} else {
		ctx.ShouldBind(&comment)
	}

	comment.UserID = id
	var photo models.Photo

	c.DB.Model(&models.Photo{}).Preload("User").Find(&photo)
	comment.PhotoID = photo.ID
	comment.UserID = photo.UserID
	comment.Photo = photo
	comment.User = photo.User

	if err := c.DB.Debug().Create(&comment).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"message":"gagal crated comment",
			"msg_dev":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": &comment,
	})
}


func (c *Controller)  GetComments(ctx *gin.Context) {

	var comments []models.Comment

	c.DB.Debug().Model(&models.Comment{}).Preload("User").Preload("Photo.User").Find(&comments)

	ctx.JSON(http.StatusOK, gin.H{
		"data": comments,
	})
}


func (c *Controller) UpdateComment(ctx *gin.Context) {
	
	var comment models.Comment
	contentType := helpers.GetContentType(ctx)

	idString, _ := ctx.Params.Get("commentID")
	id, _ := strconv.Atoi(idString)

	if err := c.DB.First(&comment, id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("user with id %s not found", id),
			"msg_dev": err.Error(),
		})
		return 		
	}

	if contentType == appJson {
		ctx.ShouldBindJSON(&comment)
	} else {
		ctx.ShouldBind(&comment)
	}

	if err := c.DB.Save(&comment).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message" : "failure updated data to database", 
			"msg_dev" :err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":"successfully updated data",
		"data": &comment,
	})

}


func (c *Controller) DeleteComment(ctx *gin.Context) {
	ParamID, _ := ctx.Params.Get("commentID")
	id, _ := strconv.Atoi(ParamID)

	var comment models.Comment
	// check is exists?
	fmt.Println(id)

	err := c.DB.First(&comment, id).Error 

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{
			"status":"not found",
			"msg_dev":err.Error(),
		})
		return
	}

	if err := c.DB.Delete(&comment).Error; err != nil {
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