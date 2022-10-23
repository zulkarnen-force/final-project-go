package controllers

import (
	"errors"
	"final-project-go/helpers"
	"final-project-go/mappers"
	"final-project-go/models"
	"final-project-go/services"
	"fmt"
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

var appJson string = "application/json"

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService *services.UserService) UserController {
	return UserController{UserService: *userService}
}


func (controller *UserController) Register(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)

	user := models.User{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}
	
	response, err := controller.UserService.Register(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"bad request",
			"message":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}


func (controller *UserController) Login(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)

	user := models.User{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}
	
	response, err := controller.UserService.Login(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}


func (controller *UserController) Update(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	id := int(userData["id"].(float64))
	var user models.User

	if contentType == appJson {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	response, err := controller.UserService.Update(user, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
	}

	ctx.JSON(http.StatusCreated, response)
}


func (controller *UserController) Delete(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims) // get info from JWT payload 
	id := int(userData["id"].(float64))
	var user models.User

	response, err := controller.UserService.Delete(user, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, response)
			return
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}
	}


	ctx.JSON(http.StatusOK, response)
}

// GetOrders godoc
// @Summary      Show an orders
// @Description  get orders data
// @Tags         orders
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.User
// @Router       /orders [get]
// func (c *UserController) Login(ctx *gin.Context) {
	
// 	contentType := helpers.GetContentType(ctx)

// 	User := models.User{}
	

// 	if contentType == appJson {
// 		ctx.ShouldBindJSON(&User)
// 	} else {
// 		ctx.ShouldBind(&User)
// 	}

// 	password := User.Password
// 	_ = password

	// err := c.DB.Debug().Where("email = ?", User.Email).Take(&User).Error
	
	// ctx.JSON(200, &User)


	// if err != nil {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{
	// 		"error":"unauthorized",
	// 		"message":"invalid email/password",
	// 	})
	// 	return
	// }

	// isValidPassword := helpers.ComparePassword([]byte(User.Password), []byte(password))

	// if !isValidPassword {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{
	// 		"error":"unauthorized",
	// 		"message":"invalid email/password",
	// 	})
	// 	return
	// }

	// token := helpers.GenerateToken(User.ID, User.Email)

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"token": token,

	// })
// }


func (c *Controller)  UserRegister(ctx *gin.Context) {

	contentType := helpers.GetContentType(ctx)

	user := models.User{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}
	
	err := c.DB.Debug().Create(&user).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"bad request",
			"message":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, mappers.GetResponseRegister(user))
}


func (c *Controller) UserUpdate(ctx *gin.Context) {
	
	var UserUpdated models.User
	contentType := helpers.GetContentType(ctx)
	var userData jwt.MapClaims = ctx.MustGet("userData").(jwt.MapClaims)
	var userID int = int(userData["id"].(float64))

	if err := c.DB.First(&UserUpdated, userID).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("user with id %d not found", userID),
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

	ctx.JSON(http.StatusOK,
		mappers.GetResponseUpdate(UserUpdated),
	)

}


func (c *Controller) UserDelete(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims) // get info from JWT payload 
	id := int(userData["id"].(float64))
	var user models.User

	err := c.DB.First(&user, id).Error 

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{
			"status":"not found",
			"msg_dev":err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}