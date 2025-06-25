package main

import (
	"rest_api/internal/infra/http/route"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	route.RegisterUserRoutes(server)
	server.Run(":8080")
}
