package tests

import (
	"github.com/MarcosVieira71/go-saldo/src/config"
)

func GenerateAdminToken() string {
	token, _ := config.CreateJWT(1, "admin")
	return token
}

func GenerateUserToken(userID uint) string {
	token, _ := config.CreateJWT(userID, "user")
	return token
}
