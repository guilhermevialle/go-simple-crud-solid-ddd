package in_memory

import (
	"rest_api/internal/domain/entities"
	"rest_api/internal/domain/interfaces/repositories"
)

type UserRepository struct {
	repositories.IUserRepository
}

var users = []entities.User{}

func (repo *UserRepository) Save(user *entities.User) error {
	users = append(users, *user)
	return nil
}

func (repo *UserRepository) FindByID(id string) (*entities.User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, nil
}
