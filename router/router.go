package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mochachmar/pbi-btpns-fd-mochachmar/controllers"
	"github.com/mochachmar/pbi-btpns-fd-mochachmar/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/users/register", controllers.RegisterUser)
	r.POST("/users/login", controllers.LoginUser)
	r.GET("/photos", controllers.GetPhotos)
	r.GET("/photos/:photoId", controllers.GetPhotoByID)

	authGroup := r.Group("/auth")
	authGroup.Use(middlewares.AuthMiddleware())
	{
		authGroup.PUT("/users/:userId", controllers.UpdateUser)
		authGroup.DELETE("/users/:userId", controllers.DeleteUser)

		authGroup.POST("/photos", controllers.CreatePhoto)
		authGroup.GET("/photos", controllers.GetPhotos)
		authGroup.GET("/photos/:photoId", controllers.GetPhotoByID)
		authGroup.PUT("/photos/:photoId", controllers.UpdatePhoto)
		authGroup.DELETE("/photos/:photoId", controllers.DeletePhoto)
	}

	r.Use(middlewares.AuthMiddleware())

	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.RegisterUser)
		userGroup.POST("/login", controllers.LoginUser)
		userGroup.PUT("/:userId", controllers.UpdateUser)
		userGroup.DELETE("/:userId", controllers.DeleteUser)
	}

	photoGroup := r.Group("/photos")
	{
		photoGroup.POST("/", controllers.CreatePhoto)
		photoGroup.GET("/", controllers.GetPhotos)
		photoGroup.PUT("/:photoId", controllers.UpdatePhoto)
		photoGroup.DELETE("/:photoId", controllers.DeletePhoto)
	}

	return r
}
