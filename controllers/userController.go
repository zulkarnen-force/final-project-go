package controllers

import (
	"errors"
	"final-project-go/dto"
	"final-project-go/entity"
	"final-project-go/helpers"
	"final-project-go/mappers"
	"final-project-go/services"
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


// Create User Account godoc
// @Summary      Create user account
// @Description  create user account
// @Tags         users
// @Param        account  body       dto.UserRegisterInput true  "Add account"
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.UserResponseRegister
// @Router       /users/register/ [post]
func (controller *UserController) Register(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)

	user := entity.User{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}
	
	user, err := controller.UserService.Register(user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, dto.ErrorResponse{
			Message: "error register user because " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, mappers.GetResponseRegister(user))
}



// Login godoc
// @Summary      Authentication user
// @Description  authentication and return JWT token 
// @Tags         users
// @Accept       json
// @Param        account  body      dto.UserLoginInput  true  "Auth account"
// @Produce      json
// @Success      200  {object}  dto.TokenResponse
// @Router       /users/login/ [post]
// @Security ApiKeyAuth
func (controller *UserController) Login(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	user := entity.User{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	hp := user.Password 
	
	user, err := controller.UserService.Login(user)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "error login because " + err.Error()})
		return
	}

	ok := helpers.ComparePassword(user.Password, hp)

	if ok == false {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "invalid password",
		})
		return
	}


	token := helpers.GenerateToken(user.ID, user.Email)

	ctx.JSON(http.StatusCreated, dto.TokenResponse{Token : token})
}

// UpdateUser godoc
// @Summary      Update user data
// @Description  Update user data 
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.UserResponseUpdate
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /users/  [put]
// @Security ApiKeyAuth
func (controller *UserController) Update(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	id := int(userData["id"].(float64))
	_ = id
	var user entity.User

	user, err := controller.UserService.GetUserByID(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, dto.ErrorResponse{
			Message: "error update user because " + err.Error(),
		})
		return
	}

	if contentType == appJson {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	user, err = controller.UserService.Update(user, id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "error update user because " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, mappers.GetResponseUpdate(user))

}

// DeleteUser godoc
// @Summary      Delete current User
// @Description   Delete current User
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.SuccessResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /users/  [delete]
// @Security ApiKeyAuth
func (controller *UserController) Delete(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims) // get info from JWT payload 
	id := int(userData["id"].(float64))
	var user entity.User

	user, err := controller.UserService.Delete(user, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, dto.ErrorResponse{Message:  "error update user because " + err.Error()})
			return
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{Message:  "error update user because " + err.Error()})
			return
		}
	}


	ctx.JSON(http.StatusOK, dto.SuccessResponse{Message: "Your user has been successfully deleted"})
}
