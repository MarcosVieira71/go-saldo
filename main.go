package main

import (
	"github.com/MarcosVieira71/go-saldo/src/config"
	"github.com/MarcosVieira71/go-saldo/src/controllers"
	"github.com/MarcosVieira71/go-saldo/src/routes"
)

func main() {
	DB := config.InitDB()
	r := routes.SetupRoutes(DB, controllers.NewUserController(DB))
	r.Run()
}
