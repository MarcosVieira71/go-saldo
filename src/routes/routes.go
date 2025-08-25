package routes

import (
	"github.com/MarcosVieira71/go-saldo/src/controllers"
	"github.com/MarcosVieira71/go-saldo/src/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	authController := controllers.NewAuthController(db)
	userController := controllers.NewUserController(db)

	public := r.Group("/auth")
	{
		public.POST("/register", authController.Register)
		public.POST("/login", authController.Login)
	}

	users := r.Group("/users")
	{
		users.GET("", middlewares.AdminOnly(), userController.GetAllUsers)
		users.GET("/:id", middlewares.UserOnly(), userController.GetByID)
		users.PUT("/:id", middlewares.UserOnly(), userController.UpdateUser)
		users.DELETE("/:id", middlewares.UserOnly(), userController.DeleteUser)
	}

	return r
}
