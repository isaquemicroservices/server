package product

import "github.com/gin-gonic/gin"

// Router it's a router of product
func Router(r *gin.RouterGroup) {
	r.GET("", getProducts)
	r.POST("", addProduct)
}

// Router it's a router of product with id
func RouterWithID(r *gin.RouterGroup) {
	r.GET(":id", getDetailsProduct)
}
