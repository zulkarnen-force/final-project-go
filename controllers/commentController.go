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

type CommentController struct {
	Service services.CommentService
}

func NewCommentController(s *services.CommentService) CommentController {
	return CommentController{Service: *s}
}


func (c *CommentController) CreateComment(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	id := int(userData["id"].(float64))

	comment := entity.Comment{}
	comment.UserID = id

	if contentType == appJson {
		ctx.ShouldBindJSON(&comment)
	} else {
		ctx.ShouldBind(&comment)
	}

	comment, err := c.Service.Update(comment)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, mappers.GetCommentCreateResponse(comment))
}


func (c *CommentController)  GetComments(ctx *gin.Context) {

	comments, err := c.Service.GetAll()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, dto.ErrorResponse{Message: "comments not found"})
		return
	}

	ctx.JSON(http.StatusOK, mappers.GetResponseComments(&comments))
}


func (c *CommentController) UpdateComment(ctx *gin.Context) {
	
	contentType := helpers.GetContentType(ctx)

	idString, _ := ctx.Params.Get("commentID")
	id, _ := strconv.Atoi(idString)

	comment, err := c.Service.GetByID(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "comment not found",
		})
		return 		
	}

	if contentType == appJson {
		ctx.ShouldBindJSON(&comment)
	} else {
		ctx.ShouldBind(&comment)
	}

	comment, err = c.Service.Update(comment)
	
	 if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: fmt.Sprintf("failure updated data because %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusCreated, comment)

}


func (c *CommentController) DeleteComment(ctx *gin.Context) {
	ParamID, _ := ctx.Params.Get("commentID")
	id, _ := strconv.Atoi(ParamID)

	comment, err := c.Service.GetByID(id) 

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, dto.ErrorResponse{
			Message: "comment not found",
		})
		return
	}

	comment, err = c.Service.Delete(comment)
	
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "failed to delete comment because " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":"Your comment has been successfully deleted",
	})
}