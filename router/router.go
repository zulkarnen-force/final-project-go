package router

import (
	"final-project-go/controllers"
	"final-project-go/databases"
	docs "final-project-go/docs"
	"final-project-go/middlewares"
	"final-project-go/repository"
	"final-project-go/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func Router() *gin.Engine {

	docs.SwaggerInfo.BasePath = "/"

	db :=  databases.ConnectionDB()
	router := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userServices := services.NewUserService(&userRepo)
	userController := controllers.NewUserController(&userServices)

	controllers := controllers.Controller{
		DB: db,
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	router.POST("users/login", controllers.UserLogin)
	router.POST("users/register", userController.Register)

	userAuthRequired := router.Group("/users")
	userAuthRequired.Use(middlewares.Authentication())
	{
		userAuthRequired.PUT("/", controllers.UserUpdate)
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

	commentRouters := router.Group("/comments")
	commentRouters.Use(middlewares.Authentication())
	{
		commentRouters.POST("/", controllers.CreateComment)
		commentRouters.GET("/", controllers.GetComments)
		commentRouters.PUT("/:commentID", controllers.UpdateComment)
		commentRouters.DELETE("/:commentID", controllers.DeleteComment)
	}


	socialMediaRouters := router.Group("/socialmedias")
	socialMediaRouters.Use(middlewares.Authentication())
	{
		socialMediaRouters.POST("/", controllers.CreateSocialMedia)
		socialMediaRouters.GET("/", controllers.GetSocialMedias)
		socialMediaRouters.PUT("/:id", controllers.UpdateSocialMedia)
		socialMediaRouters.DELETE("/:id", controllers.DeleteSocialMedia)
	}

	return router
}