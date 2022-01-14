package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/servers-microservices-backend/middleware"
)

// Router it's a router of auth
func RouterWithAuth(r *gin.RouterGroup) {
	r.Use(middleware.AuthorizationGin())

	r.POST("/create", create)
}

// Router it's a router
func RouterWithoutAuth(r *gin.RouterGroup) {
	r.POST("/login", login)
}
