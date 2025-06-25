package containers

import (
	use_cases "rest_api/internal/app/use-cases"
	"rest_api/internal/infra/http/handlers"
	in_memory "rest_api/internal/infra/repositories/in-memory"
)

type Container struct {
	UserHandler *handlers.UserHandler
}

func NewUserContainer() *Container {
	userRepo := in_memory.UserRepository{}
	createUserHandler := use_cases.NewCreateUserHandler(&userRepo)
	userController := handlers.NewUserHandler(createUserHandler)

	return &Container{
		UserHandler: userController,
	}
}
