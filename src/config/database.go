package config

import (
	"log"
	"os/user"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	err = DB.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal("Erro na migração do banco:", err)
	}
}
