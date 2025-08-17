package middlewares

import (
	"github.com/MarcosVieira71/go-saldo/src/config"
	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "token não fornecido"})
			c.Abort()
			return
		}

		claims, err := config.ParseJWT(tokenString)
		if err != nil || claims == nil {
			c.JSON(401, gin.H{"error": "token inválido"})
			c.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok || role != "admin" {
			c.JSON(403, gin.H{"error": "acesso negado"})
			c.Abort()
			return
		}

		c.Next()
	}
}
