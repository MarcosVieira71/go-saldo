package middlewares

import (
	"strconv"

	"github.com/MarcosVieira71/go-saldo/src/config"
	"github.com/gin-gonic/gin"
)

func UserOnly() gin.HandlerFunc {
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

		tokenUserID, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(401, gin.H{"error": "token inválido"})
			c.Abort()
			return
		}

		paramID := c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.JSON(400, gin.H{"error": "ID inválido"})
			c.Abort()
			return
		}

		if uint(id) != uint(tokenUserID) {
			c.JSON(403, gin.H{"error": "acesso negado"})
			c.Abort()
			return
		}

		c.Next()
	}
}
