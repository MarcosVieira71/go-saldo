package tests

import (
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/MarcosVieira71/go-saldo/src/models/user"
	"github.com/MarcosVieira71/go-saldo/tests"
)

func TestCreateUser(t *testing.T) {
	_, err := user.CreateUser("João", "joao@email.com", "123")
	if err != nil {
		t.Errorf("Erro ao criar usuário: %v", err)
	}
}

func TestAddUser(t *testing.T) {
	db := tests.SetupTestDB(t)
	u, _ := user.CreateUser("João", "joao@email.com", "123")
	err := user.AddUser(db, u)
	if err != nil {
		t.Errorf("erro ao adicionar usuário: %v", err)
	}
}

func TestNewPassword(t *testing.T) {
	db := tests.SetupTestDB(t)

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

func TestGetUserByID(t *testing.T) {
	db := tests.SetupTestDB(t)

	u, _ := user.CreateUser("João", "joao@email.com", "123")
	_ = user.AddUser(db, u)

	_, err := user.GetUserByID(db, 1)
	if err != nil {
		t.Errorf("Erro ao encontrar usuário")
	}
}
func TestGetUserByNonExistentID(t *testing.T) {
	db := tests.SetupTestDB(t)

	u, _ := user.CreateUser("João", "joao@email.com", "123")
	_ = user.AddUser(db, u)

	_, err := user.GetUserByID(db, 2)
	if err == nil {
		t.Errorf("Encontrou usuário inexistente")
	}
}

func TestGetUserByEmail(t *testing.T) {
	db := tests.SetupTestDB(t)

	u, _ := user.CreateUser("João", "joao@email.com", "123")
	_ = user.AddUser(db, u)

	_, err := user.GetUserByEmail(db, "joao@email.com")
	if err != nil {
		t.Errorf("Erro ao encontrar usuário")
	}
}

func TestGetUserByNonExistentEmail(t *testing.T) {
	db := tests.SetupTestDB(t)

	_, err := user.GetUserByEmail(db, "joao@email.com")
	if err == nil {
		t.Errorf("Não há usuário encontrado com e-mail")
	}
}

func TestAddUserWithSameEmail(t *testing.T) {
	db := tests.SetupTestDB(t)

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
func TestDeleteUser(t *testing.T) {
	db := tests.SetupTestDB(t)

	u, _ := user.CreateUser("Maria", "maria@email.com", "123")
	_ = user.AddUser(db, u)

	deleted, err := user.DeleteUser(db, int(u.Id))
	if err != nil {
		t.Fatalf("Erro ao deletar usuário: %v", err)
	}

	if deleted.Id != u.Id {
		t.Errorf("Usuário deletado não corresponde ao esperado")
	}

	_, err = user.GetUserByID(db, int(u.Id))
	if err == nil {
		t.Errorf("Usuário deveria ter sido removido")
	}
}

func TestUpdateUser(t *testing.T) {
	db := tests.SetupTestDB(t)

	u, _ := user.CreateUser("Pedro", "Pedro@email.com", "123")
	_ = user.AddUser(db, u)

	updated, err := user.UpdateUser(db, int(u.Id), "Carlos Silva", "carlossilva@email.com", "456")
	if err != nil {
		t.Fatalf("Erro ao atualizar usuário: %v", err)
	}

	if updated.Name != "Carlos Silva" || updated.Email != "carlossilva@email.com" {
		t.Errorf("Nome ou email não foram atualizados corretamente")
	}

	err = bcrypt.CompareHashAndPassword([]byte(updated.Password), []byte("456"))
	if err != nil {
		t.Errorf("Senha não foi atualizada corretamente")
	}
}

func TestUpdateUserWithoutPassword(t *testing.T) {
	db := tests.SetupTestDB(t)

	u, _ := user.CreateUser("Pedro", "Pedro@email.com", "123")
	_ = user.AddUser(db, u)

	updated, err := user.UpdateUser(db, int(u.Id), "Carlos Silva", "carlossilva@email.com", "")
	if err != nil {
		t.Fatalf("Erro ao atualizar usuário: %v", err)
	}

	if updated.Name != "Carlos Silva" || updated.Email != "carlossilva@email.com" {
		t.Errorf("Nome ou email não foram atualizados corretamente")
	}

	err = bcrypt.CompareHashAndPassword([]byte(updated.Password), []byte("123"))
	if err != nil {
		t.Errorf("Senha foi atualizada incorretamente")
	}
}

func TestGetAllUsers(t *testing.T) {
	db := tests.SetupTestDB(t)

	u1, _ := user.CreateUser("Ana", "ana@email.com", "123")
	u2, _ := user.CreateUser("Pedro", "pedro@email.com", "123")
	_ = user.AddUser(db, u1)
	_ = user.AddUser(db, u2)

	usersList, err := user.GetAllUsers(db)
	if err != nil {
		t.Fatalf("Erro ao buscar todos os usuários: %v", err)
	}

	if len(usersList) != 2 {
		t.Errorf("Esperado 2 usuários, mas obteve %d", len(usersList))
	}
}

func TestDeleteUserNonExistent(t *testing.T) {
	db := tests.SetupTestDB(t)

	_, err := user.DeleteUser(db, 999)
	if err == nil {
		t.Errorf("Deveria retornar erro ao tentar deletar usuário inexistente")
	}
}

func TestUpdateUserNonExistent(t *testing.T) {
	db := tests.SetupTestDB(t)

	_, err := user.UpdateUser(db, 999, "Novo Nome", "novo@email.com", "senha")
	if err == nil {
		t.Errorf("Deveria retornar erro ao tentar atualizar usuário inexistente")
	}
}

func TestUpdateUserEmptyPassword(t *testing.T) {
	db := tests.SetupTestDB(t)

	u, _ := user.CreateUser("Lucas", "lucas@email.com", "123")
	_ = user.AddUser(db, u)

	updated, err := user.UpdateUser(db, int(u.Id), "Lucas Atualizado", "lucas@email.com", "")
	if err != nil {
		t.Fatalf("Erro ao atualizar usuário com senha vazia: %v", err)
	}

	if updated.Password != u.Password {
		t.Errorf("Senha não deveria ser alterada quando a senha nova é vazia")
	}
}
