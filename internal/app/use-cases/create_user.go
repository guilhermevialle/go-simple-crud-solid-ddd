package use_cases

import (
	"rest_api/internal/domain/entities"
	"rest_api/internal/domain/interfaces/repositories"
)

type Handler struct {
	repo repositories.IUserRepository
}

func NewCreateUserHandler(repo repositories.IUserRepository) *Handler {
	return &Handler{repo: repo}
}

func (handler *Handler) Handle(user *entities.User) error {
	return handler.repo.Save(user)
}
