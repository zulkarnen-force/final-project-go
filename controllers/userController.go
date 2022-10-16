package controllers

import (
	"github.com/gin-gonic/gin"
	"final-project-go/databases"
	"final-project-go/models"
)

db := databases.GetDB()

func UserLogin(ctx *gin.Context) {

	contentType := helpers.GetContentType(ctx)

	User := models.User{}
	password := ""

	if contentType == appJson {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error


	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":"unauthorized",
			"message":"invalid email/password",
		})
		return
	}

	isValidPassword := helpers.ComparePassword([]byte(User.Password), []byte(password))

	if !isValidPassword {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":"unauthorized",
			"message":"invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)
	fmt.Println("Token => ", token)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,

	})
}


func UserRegister(ctx *gin.Context) {

}


func UserUpdate(ctx *gin.Context) {

}


func UserDelete(ctx *gin.Context) {

}