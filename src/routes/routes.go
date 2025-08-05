package routes

import (
	"github.com/MarcosVieira71/go-saldo/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userController := controllers.NewUserController(db)

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", userController.CreateUser)
		userRoutes.GET("/:id", userController.GetByID)
	}

	return r
}
