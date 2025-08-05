package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Aviso: .env não encontrado")
	}
	fmt.Println("JWT_KEY from env:", os.Getenv("JWT_KEY"))
}

var JwtKey = []byte(os.Getenv("JWT_KEY"))
