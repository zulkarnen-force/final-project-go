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

var appJson string = "application/json"

func (c *Controller) UserLogin(ctx *gin.Context) {
	
	contentType := helpers.GetContentType(ctx)

	User := models.User{}
	password := ""

	if contentType == appJson {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	password = User.Password

	err := c.DB.Debug().Where("email = ?", User.Email).Take(&User).Error


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


func (c *Controller)  UserRegister(ctx *gin.Context) {

	contentType := helpers.GetContentType(ctx)

	User := models.User{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}
	
	err := c.DB.Debug().Create(&User).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"bad request",
			"message":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":User.ID,
		"email":User.Email,
	})
}


func (c *Controller) UserUpdate(ctx *gin.Context) {
	
	var UserUpdated models.User
	contentType := helpers.GetContentType(ctx)

	// get ID using jsonData
	id, _ := ctx.Params.Get("id")
	userID, _ := strconv.Atoi(id)

	if err := c.DB.First(&UserUpdated, userID).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("user with id %s not found", id),
			"msg_dev": err.Error(),
		})
		return 		
	}

	if contentType == appJson {
		ctx.ShouldBindJSON(&UserUpdated)
	} else {
		ctx.ShouldBind(&UserUpdated)
	}

	if err := c.DB.Save(&UserUpdated).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message" : "failure updated data to database", 
			"msg_dev" :err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":"successfully updated data",
		"data": &UserUpdated,
	})

}


func (c *Controller) UserDelete(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims) // get info from JWT payload 
	id := int(userData["id"].(float64))
	var user models.User


	// check is exists?
	fmt.Println(id)

	err := c.DB.First(&user, id).Error 

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{
			"status":"not found",
			"msg_dev":err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":"successfully deleted",
	})
}