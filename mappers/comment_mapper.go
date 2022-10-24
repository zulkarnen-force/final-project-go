package mappers

import (
	"final-project-go/dto"
	"final-project-go/entity"

	"github.com/jinzhu/copier"
)



func GetResponseComments(commentDatas *[]entity.Comment) *[]dto.CommentResponse {

	var comments []dto.CommentResponse
	for _, data := range *commentDatas {
		comment := new(dto.CommentResponse)
		copier.Copy(comment, &data)
		comments = append(comments, *comment)
	}

	return &comments

}

func GetCommentCreateResponse(comment entity.Comment) *dto.CommentResponseCreate {
	commentResponse := new(dto.CommentResponseCreate)
	copier.Copy(&commentResponse, &comment)
	return commentResponse
}

func GetCommentUpdateResponse(comment entity.Comment) *dto.CommentResponseUpdate {
	commentResponse := new(dto.CommentResponseUpdate)
	copier.Copy(&commentResponse, &comment)
	return commentResponse
}