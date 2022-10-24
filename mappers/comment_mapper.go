package mappers

import (
	"final-project-go/entity"
	"final-project-go/models"

	"github.com/jinzhu/copier"
)



func GetResponseComments(commentDatas *[]entity.Comment) *[]models.CommentResponse {

	var comments []models.CommentResponse
	for _, data := range *commentDatas {
		comment := new(models.CommentResponse)
		copier.Copy(comment, &data)
		comments = append(comments, *comment)
	}

	return &comments

}

func GetCommentCreateResponse(comment entity.Comment) *models.CommentResponseCreate {
	commentResponse := new(models.CommentResponseCreate)
	copier.Copy(&commentResponse, &comment)
	return commentResponse
}

func GetCommentUpdateResponse(comment entity.Comment) *models.CommentResponseUpdate {
	commentResponse := new(models.CommentResponseUpdate)
	copier.Copy(&commentResponse, &comment)
	return commentResponse
}