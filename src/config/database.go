package config

import (
	"log"

	"github.com/MarcosVieira71/go-saldo/models/user"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	var err error
	DB, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	err = DB.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal("Erro na migração do banco:", err)
	}

	return DB
}
