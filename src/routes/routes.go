package routes

import (
	"github.com/MarcosVieira71/go-saldo/src/controllers"
	"github.com/MarcosVieira71/go-saldo/src/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, userController *controllers.UserController) *gin.Engine {
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", userController.CreateUser)
		userRoutes.POST("/login", userController.Login)

		userRoutes.GET("", middlewares.AdminOnly(), userController.GetAllUsers)

		userRoutes.GET("/:id", middlewares.UserOnly(), userController.GetByID)
		userRoutes.PUT("/:id", middlewares.UserOnly(), userController.UpdateUser)
		userRoutes.DELETE("/:id", middlewares.UserOnly(), userController.DeleteUser)
	}

	return r
}
