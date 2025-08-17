package routes

import (
	"github.com/MarcosVieira71/go-saldo/src/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, userController *controllers.UserController) *gin.Engine {
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", userController.CreateUser)
		userRoutes.POST("/login", userController.Login)
		userRoutes.GET("", userController.GetAllUsers)
		userRoutes.GET("/:id", userController.GetByID)
		userRoutes.DELETE("/:id", userController.DeleteUser)
		userRoutes.PUT("/:id", userController.UpdateUser)
	}

	return r
}
