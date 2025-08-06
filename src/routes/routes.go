package routes

import (
	"github.com/MarcosVieira71/go-saldo/src/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userController := controllers.NewUserController(db)

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", userController.CreateUser)
		userRoutes.GET("", userController.GetAllUsers)
		userRoutes.GET("/:id", userController.GetByID)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}

	return r
}
