package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Porq eu?")
		log.Fatal("Aviso: .env não encontrado")
	}
	// fmt.Println("JWT_KEY from env:", os.Getenv("JWT_KEY"))
}

var JwtKey = []byte(os.Getenv("JWT_KEY"))
