package routes

import (
	"github.com/MarcosVieira71/go-saldo/src/controllers"
	"github.com/MarcosVieira71/go-saldo/src/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, userController *controllers.UserController) *gin.Engine {
	r := gin.Default()

	public := r.Group("/auth")
	{
		public.POST("/register", userController.CreateUser)
		public.POST("/login", userController.Login)
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
