package tests

import (
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/MarcosVieira71/go-saldo/models/user"
)

func TestCreateUser(t *testing.T) {
	_, err := user.CreateUser("João", "joao@email.com", "123")
	if err != nil {
		t.Errorf("Erro ao criar usuário: %v", err)
	}
}

func TestAddUser(t *testing.T) {
	db := setupTestDB(t)
	u, _ := user.CreateUser("João", "joao@email.com", "123")
	err := user.AddUser(db, u)
	if err != nil {
		t.Errorf("erro ao adicionar usuário: %v", err)
	}
}

func TestNewPassword(t *testing.T) {
	db := setupTestDB(t)

	plainPassword := "123"
	u, _ := user.CreateUser("João", "joao@email.com", plainPassword)
	err := user.AddUser(db, u)
	if err != nil {
		t.Fatalf("Erro ao adicionar usuário: %v", err)
	}

	dbUser, err := user.GetUserByEmail(db, "joao@email.com")
	if err != nil {
		t.Fatalf("Erro ao buscar usuário: %v", err)
	}

	if dbUser.Password == plainPassword {
		t.Errorf("Senha armazenada não deve ser igual à senha em texto puro")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(plainPassword))
	if err != nil {
		t.Errorf("Senha armazenada não corresponde ao hash: %v", err)
	}
}

func TestGetUserByEmail(t *testing.T) {
	db := setupTestDB(t)

	u, _ := user.CreateUser("João", "joao@email.com", "123")
	_ = user.AddUser(db, u)

	_, err := user.GetUserByEmail(db, "joao@email.com")
	if err != nil {
		t.Errorf("Erro ao encontrar usuário")
	}
}

func TestGetUserByNonExistentEmail(t *testing.T) {
	db := setupTestDB(t)

	_, err := user.GetUserByEmail(db, "joao@email.com")
	if err == nil {
		t.Errorf("Não há usuário encontrado com e-mail")
	}
}

func TestAddUserWithSameEmail(t *testing.T) {
	db := setupTestDB(t)

	u, _ := user.CreateUser("João", "joao@email.com", "123")
	err := user.AddUser(db, u)
	if err != nil {
		t.Fatalf("Erro ao adicionar primeiro usuário: %v", err)
	}

	u2, _ := user.CreateUser("JoaoPedro", "joao@email.com", "123")
	err = user.AddUser(db, u2)
	if err == nil {
		t.Errorf("Usuário não deveria poder ser adicionado com email já cadastrado")
	}
}
