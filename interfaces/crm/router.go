package crm

import (
	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/servers-microservices-backend/interfaces/crm/product"
)

// Router it's a router of products
func Router(r *gin.RouterGroup) {
	product.Router(r.Group("products"))
	product.RouterWithID(r.Group("product"))
}
