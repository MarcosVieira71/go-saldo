package config

import "github.com/gin-gonic/gin"

func JSONResponse(c *gin.Context, status int, data interface{}, message string, err string) {
	c.JSON(status, gin.H{
		"data":    data,
		"message": message,
		"error":   err,
	})
}
