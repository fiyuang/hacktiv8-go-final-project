package router

import (
	"hacktiv8-go-final-project/controllers"
	"hacktiv8-go-final-project/middleware"
	"hacktiv8-go-final-project/repository"
	"hacktiv8-go-final-project/service"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	photoRepository := repository.NewPhotoRepository()
	photoService := service.NewPhotoService(photoRepository)
	photoController := controllers.NewPhotoController(photoService)

	commentRepository := repository.NewCommentRepository()
	commentService := service.NewCommentService(commentRepository)
	commentController := controllers.NewCommentController(commentService)

	socialMediaRepository := repository.NewSocialMediaRepository()
	socialMediaService := service.NewSocialMediaService(socialMediaRepository)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.UserRegister)
		userRouter.POST("/login", userController.UserLogin)

		userRouter.Use(middleware.Authentication())
		userRouter.PUT("/:userId", userController.UpdateUser)
		userRouter.DELETE("/:userId", userController.DeleteUser)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.GET("/", photoController.GetAllPhotos)
		photoRouter.POST("/", photoController.CreatePhoto)
		photoRouter.DELETE("/:id", photoController.DeletePhoto)
		photoRouter.PUT("/:photoId", photoController.UpdatePhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.GET("/", commentController.GetAllComments)
		commentRouter.POST("/", commentController.CreateComment)
		commentRouter.DELETE("/:commentId", commentController.DeleteComment)
		commentRouter.PUT("/:commentId", commentController.UpdateComment)
	}

	socialMediaRouter := router.Group("/socialMedias")
	{
		socialMediaRouter.Use(middleware.Authentication())
		socialMediaRouter.POST("/create", socialMediaController.CreateSocialMedia)
		socialMediaRouter.DELETE("/delete/:id", socialMediaController.DeleteSocialMedia)
	}

	return router
}
