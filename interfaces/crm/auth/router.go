package auth

import "github.com/gin-gonic/gin"

// Router it's a router of auth
func Router(r *gin.RouterGroup) {
	r.POST("/create", create)
	r.POST("/login", login)
	r.POST("/logout", logout)
}
