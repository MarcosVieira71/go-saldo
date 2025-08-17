package tests

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/MarcosVieira71/go-saldo/src/controllers"
	"github.com/MarcosVieira71/go-saldo/src/models/user"
	"github.com/MarcosVieira71/go-saldo/src/routes"
	"github.com/MarcosVieira71/go-saldo/tests"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouterWithDB(t *testing.T) (*gin.Engine, *controllers.UserController) {
	db := tests.SetupTestDB(t)
	uc := controllers.NewUserController(db)

	router := routes.SetupRoutes(db, uc)
	return router, uc
}

func TestCreateUser_Success(t *testing.T) {
	router, _ := setupRouterWithDB(t)

	body := `{"name":"Test User","email":"test@example.com","password":"123456"}`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Test User")
	assert.Contains(t, w.Body.String(), "test@example.com")
}

func TestCreateUser_InvalidJSON(t *testing.T) {
	router, _ := setupRouterWithDB(t)

	body := `{"name":"Test User","email":"test@example.com"`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Dados inválidos")
}

func TestGetAllUsers_EmptyAndWithUsers(t *testing.T) {
	router, uc := setupRouterWithDB(t)

	req := httptest.NewRequest("GET", "/users", nil)
	req.Header.Set("Authorization", "Bearer "+tests.GenerateAdminToken())
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "[]")

	u, _ := user.CreateUser("Alice", "alice@email.com", "123")
	err := user.AddUser(uc.DB, u)
	assert.NoError(t, err)

	req = httptest.NewRequest("GET", "/users", nil)
	req.Header.Set("Authorization", "Bearer "+tests.GenerateAdminToken())
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Alice")
	assert.Contains(t, w.Body.String(), "alice@email.com")
}

func TestUpdateUser_Success(t *testing.T) {
	router, uc := setupRouterWithDB(t)

	u, _ := user.CreateUser("Bob", "bob@email.com", "123")
	err := user.AddUser(uc.DB, u)
	assert.NoError(t, err)

	body := `{"name":"Bob Updated","email":"bobupdated@email.com","password":"456"}`
	idStr := strconv.FormatUint(uint64(u.Id), 10)
	req := httptest.NewRequest("PUT", "/users/"+idStr, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+tests.GenerateUserToken(u.Id))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Bob Updated")
	assert.Contains(t, w.Body.String(), "bobupdated@email.com")
}

func TestUpdateUser_InvalidID(t *testing.T) {
	router, _ := setupRouterWithDB(t)

	body := `{"name":"NoOne","email":"noone@email.com","password":"123"}`
	req := httptest.NewRequest("PUT", "/users/invalid-id", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+tests.GenerateAdminToken())
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "ID inválido")
}

func TestDeleteUser_Success(t *testing.T) {
	router, uc := setupRouterWithDB(t)

	u, _ := user.CreateUser("Charlie", "charlie@email.com", "123")
	err := user.AddUser(uc.DB, u)
	assert.NoError(t, err)
	idStr := strconv.FormatUint(uint64(u.Id), 10)

	req := httptest.NewRequest("DELETE", "/users/"+idStr, nil)
	req.Header.Set("Authorization", "Bearer "+tests.GenerateUserToken(u.Id))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Usuário deletado")
}

func TestDeleteUser_InvalidID(t *testing.T) {
	router, _ := setupRouterWithDB(t)

	req := httptest.NewRequest("DELETE", "/users/abc", nil)
	req.Header.Set("Authorization", "Bearer "+tests.GenerateAdminToken())
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "ID inválido")
}

func TestLogin_Success(t *testing.T) {
	router, uc := setupRouterWithDB(t)

	u, _ := user.CreateUser("TestLogin", "login@example.com", "123456")
	err := user.AddUser(uc.DB, u)
	assert.NoError(t, err)

	body := `{"email":"login@example.com","password":"123456"}`
	req := httptest.NewRequest("POST", "/users/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "access_token")
	assert.Contains(t, w.Body.String(), "Login realizado com sucesso")
}

func TestLogin_Failure_InvalidCredentials(t *testing.T) {
	router, _ := setupRouterWithDB(t)

	body := `{"email":"wrong@example.com","password":"123456"}`
	req := httptest.NewRequest("POST", "/users/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Credenciais inválidas")
}
