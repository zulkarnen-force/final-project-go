package controllers

import (
	"final-project-go/helpers"
	"final-project-go/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)


func (c *Controller) CreateComment(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	id := int(userData["id"].(float64))

	comment := models.Comment{}
	comment.UserID = id

	if contentType == appJson {
		ctx.ShouldBindJSON(&comment)
	} else {
		ctx.ShouldBind(&comment)
	}

	if err := c.DB.Debug().Create(&comment).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"message":"gagal created comment",
			"msg_dev":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, comment.GetResponseCreate())
}


func (c *Controller)  GetComments(ctx *gin.Context) {

	var comments []models.Comment
	var Comment models.Comment

	c.DB.Debug().Model(&models.Comment{}).Preload("User", func (db *gorm.DB) *gorm.DB {
		return	db.Select("id", "email", "username")
	}).Preload("Photo.User").Find(&comments)

	ctx.JSON(http.StatusOK, Comment.GetResponseComments(&comments))
}


func (c *Controller) UpdateComment(ctx *gin.Context) {
	
	var comment models.Comment
	contentType := helpers.GetContentType(ctx)

	idString, _ := ctx.Params.Get("commentID")
	id, _ := strconv.Atoi(idString)

	if err := c.DB.First(&comment, id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("user with id %d not found", id),
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

	ctx.JSON(http.StatusOK, comment.GetResponseUpdate())

}


func (c *Controller) DeleteComment(ctx *gin.Context) {
	ParamID, _ := ctx.Params.Get("commentID")
	id, _ := strconv.Atoi(ParamID)

	var comment models.Comment

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
		"message":"Your comment has been successfully deleted",
	})
}