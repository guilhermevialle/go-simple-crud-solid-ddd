package route

import (
	"rest_api/internal/di/containers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	userContainer := containers.NewUserContainer()
	r.POST("/users", userContainer.UserHandler.CreateUser)
}
