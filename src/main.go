package main

import (
	"github.com/MarcosVieira71/go-saldo/config"
	"github.com/MarcosVieira71/go-saldo/routes"
)

func main() {
	DB := config.InitDB()
	r := routes.SetupRoutes(DB)
	r.Run()
}
