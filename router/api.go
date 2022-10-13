package router

import (
	"hacktiv8-go-final-project/controllers"
	"hacktiv8-go-final-project/repository"
	"hacktiv8-go-final-project/service"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	userRouter := router.Group("/users")
	{
		// userRouter.Use(middleware.Authentication())
		userRouter.POST("/register", userController.UserRegister)
		userRouter.POST("/login", userController.UserLogin)
	}

	return router
}
