package tests

import (
	"testing"

	"github.com/MarcosVieira71/go-saldo/models/user"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("Falha ao abrir banco de teste: %v", err)
	}

	err = db.AutoMigrate(&user.User{})
	if err != nil {
		t.Fatalf("Falha na migração: %v", err)
	}

	return db
}
