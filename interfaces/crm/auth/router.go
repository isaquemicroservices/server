package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/servers-microservices-backend/middleware"
)

// Router it's a router of auth
func RouterAuth(r *gin.RouterGroup) {
	r.POST("/create", middleware.AuthorizationGin(), create)
	r.POST("/login", login)
}
