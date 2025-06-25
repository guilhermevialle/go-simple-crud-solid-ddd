package repositories

import "rest_api/internal/domain/entities"

type IUserRepository interface {
	Save(user *entities.User) error
	FindByID(id string) (*entities.User, error)
}
