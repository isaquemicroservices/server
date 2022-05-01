package auth

import (
	"github.com/gin-gonic/gin"
)

// Router it's a router of auth
func RouterAuth(r *gin.RouterGroup) {
	r.POST("/create", create)
	r.POST("/login", login)
}
