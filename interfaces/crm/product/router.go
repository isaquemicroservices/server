package product

import (
	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/servers-microservices-backend/middleware"
)

// Router it's a router of product
func Router(r *gin.RouterGroup) {
	r.GET("", getProducts)
	r.POST("", middleware.AdminOnly(), addProduct)
}

// Router it's a router of product with id
func RouterWithID(r *gin.RouterGroup) {
	r.GET(":id", getDetailsProduct)
}
