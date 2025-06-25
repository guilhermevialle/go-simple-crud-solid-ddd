package handlers

import (
	"net/http"
	use_cases "rest_api/internal/app/use-cases"
	"rest_api/internal/domain/entities"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	createUserHandler *use_cases.Handler
}

func NewUserHandler(handler *use_cases.Handler) *UserHandler {
	return &UserHandler{createUserHandler: handler}
}

func (c *UserHandler) CreateUser(ctx *gin.Context) {
	var input struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entities.NewUser(input.Name, input.Email)

	if err := c.createUserHandler.Handle(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save user"})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
