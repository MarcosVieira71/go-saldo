package controllers

import (
	"net/http"

	"github.com/MarcosVieira71/go-saldo/src/config"
	"github.com/MarcosVieira71/go-saldo/src/models/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

func (ac *AuthController) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		config.JSONResponse(c, http.StatusBadRequest, nil, "", "Dados inválidos")
		return
	}

	u, err := user.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		config.JSONResponse(c, http.StatusInternalServerError, nil, "", "Erro ao criar usuário")
		return
	}

	if err := user.AddUser(ac.DB, u); err != nil {
		config.JSONResponse(c, http.StatusBadRequest, nil, "", "Erro ao salvar no banco: "+err.Error())
		return
	}

	config.JSONResponse(c, http.StatusCreated, u, "Usuário criado com sucesso", "")
}

func (ac *AuthController) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		config.JSONResponse(c, http.StatusBadRequest, nil, "", "Dados inválidos")
		return
	}

	u, err := user.AuthenticateUser(ac.DB, req.Email, req.Password)
	if err != nil {
		config.JSONResponse(c, http.StatusUnauthorized, nil, "", "Credenciais inválidas")
		return
	}

	token, err := config.CreateJWT(u.Id, u.Role)
	if err != nil {
		config.JSONResponse(c, http.StatusInternalServerError, nil, "", "Erro ao gerar token")
		return
	}

	config.JSONResponse(c, http.StatusOK, gin.H{"access_token": token}, "Login realizado com sucesso", "")
}
