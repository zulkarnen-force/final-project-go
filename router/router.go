package router

import (
	"final-project-go/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	router := gin.Default()

	users := router.Group("/users")
	{
		users.POST("/register", controllers.UserRegister)
		users.POST("/login", controllers.UserLogin)
		users.POST("/", controllers.UserUpdate)
		users.DELETE("/", controllers.UserDelete)
	}

	return router
}