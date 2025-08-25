package controllers

import (
	"net/http"
	"strconv"

	"github.com/MarcosVieira71/go-saldo/src/config"
	"github.com/MarcosVieira71/go-saldo/src/models/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		config.JSONResponse(c, http.StatusBadRequest, nil, "", "ID inválido")
		return
	}

	u, err := user.GetUserByID(uc.DB, id)
	if err != nil {
		config.JSONResponse(c, http.StatusNotFound, nil, "", "Usuário não encontrado")
		return
	}

	config.JSONResponse(c, http.StatusOK, u, "", "")
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		config.JSONResponse(c, http.StatusBadRequest, nil, "", "ID inválido")
		return
	}

	deletedUser, err := user.DeleteUser(uc.DB, id)
	if err != nil {
		config.JSONResponse(c, http.StatusNotFound, nil, "", err.Error())
		return
	}

	config.JSONResponse(c, http.StatusOK, deletedUser, "Usuário deletado", "")
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		config.JSONResponse(c, http.StatusBadRequest, nil, "", "ID inválido")
		return
	}

	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		config.JSONResponse(c, http.StatusBadRequest, nil, "", "Dados inválidos")
		return
	}

	_, err = user.GetUserByID(uc.DB, id)
	if err != nil {
		config.JSONResponse(c, http.StatusNotFound, nil, "", "Usuário não encontrado")
		return
	}

	updatedUser, err := user.UpdateUser(uc.DB, id, req.Name, req.Email, req.Password)
	if err != nil {
		config.JSONResponse(c, http.StatusInternalServerError, nil, "", "Erro ao atualizar usuário")
		return
	}

	config.JSONResponse(c, http.StatusOK, updatedUser, "Usuário atualizado com sucesso", "")
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := user.GetAllUsers(uc.DB)
	if err != nil {
		config.JSONResponse(c, http.StatusInternalServerError, nil, "", "Erro ao buscar usuários")
		return
	}

	config.JSONResponse(c, http.StatusOK, users, "", "")
}
