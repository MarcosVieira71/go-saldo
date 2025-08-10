package controllers

import (
	"net/http"
	"strconv"

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

func JSONResponse(c *gin.Context, status int, data interface{}, message string, err string) {
	c.JSON(status, gin.H{
		"data":    data,
		"message": message,
		"error":   err,
	})
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		JSONResponse(c, http.StatusBadRequest, nil, "", "Dados inválidos")
		return
	}

	u, err := user.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		JSONResponse(c, http.StatusInternalServerError, nil, "", "Erro ao criar usuário")
		return
	}

	if err := user.AddUser(uc.DB, u); err != nil {
		JSONResponse(c, http.StatusBadRequest, nil, "", "Erro ao salvar no banco: "+err.Error())
		return
	}

	JSONResponse(c, http.StatusCreated, u, "Usuário criado com sucesso", "")
}

func (uc *UserController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		JSONResponse(c, http.StatusBadRequest, nil, "", "ID inválido")
		return
	}

	u, err := user.GetUserByID(uc.DB, id)
	if err != nil {
		JSONResponse(c, http.StatusNotFound, nil, "", "Usuário não encontrado")
		return
	}

	JSONResponse(c, http.StatusOK, u, "", "")
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		JSONResponse(c, http.StatusBadRequest, nil, "", "ID inválido")
		return
	}

	deletedUser, err := user.DeleteUser(uc.DB, id)
	if err != nil {
		JSONResponse(c, http.StatusNotFound, nil, "", err.Error())
		return
	}

	JSONResponse(c, http.StatusOK, deletedUser, "Usuário deletado", "")
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		JSONResponse(c, http.StatusBadRequest, nil, "", "ID inválido")
		return
	}

	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		JSONResponse(c, http.StatusBadRequest, nil, "", "Dados inválidos")
		return
	}

	existingUser, err := user.GetUserByID(uc.DB, id)
	if err != nil {
		JSONResponse(c, http.StatusNotFound, nil, "", "Usuário não encontrado")
		return
	}

	if _, err := user.UpdateUser(uc.DB, id, req.Name, req.Email, req.Password); err != nil {
		JSONResponse(c, http.StatusInternalServerError, nil, "", "Erro ao atualizar usuário")
		return
	}

	JSONResponse(c, http.StatusOK, existingUser, "Usuário atualizado com sucesso", "")
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := user.GetAllUsers(uc.DB)
	if err != nil {
		JSONResponse(c, http.StatusInternalServerError, nil, "", "Erro ao buscar usuários")
		return
	}

	JSONResponse(c, http.StatusOK, users, "", "")
}
