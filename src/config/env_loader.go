package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Print("Aviso: .env não encontrado")
	}
	if len(JwtKey) == 0 {
		log.Fatal("JWT_KEY não pode estar vazio")
	}
}

var JwtKey = []byte(os.Getenv("JWT_KEY"))
