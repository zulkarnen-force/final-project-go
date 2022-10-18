package router

import (
	"final-project-go/controllers"
	"final-project-go/databases"
	"final-project-go/middlewares"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	db :=  databases.ConnectionDB()
	router := gin.Default()

	controllers := controllers.Controller{
		DB: db,
	}

	router.POST("users/login", controllers.UserLogin)
	router.POST("users/register", controllers.UserRegister)

	userAuthRequired := router.Group("/users")
	userAuthRequired.Use(middlewares.Authentication())
	{
		userAuthRequired.PUT("/:id", controllers.UserUpdate)
		userAuthRequired.DELETE("/", controllers.UserDelete)
	}


	photoRouters := router.Group("/photos")
	photoRouters.Use(middlewares.Authentication())
	{
		photoRouters.POST("/", controllers.CreatePhoto)
		photoRouters.GET("/", controllers.GetPhotos)
		photoRouters.PUT("/:photoId", controllers.UpdatePhoto)
		photoRouters.DELETE("/:photoId", controllers.DeletePhoto)
	}

	return router
}