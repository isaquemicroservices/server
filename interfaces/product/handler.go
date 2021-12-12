package product

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/servers-microservices-backend/application/product"
)

// getProducts godoc
// @Summary List products
// @Description List all products of the store
// @Tags Products
// @Produce json
// @Success 200 {object} product.ListProducts "List of products"
// @Router /products [get]
func getProducts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, time.Second*20)
	defer cancel()

	data, err := product.GetProducts(ctx)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting products",
			"cause":   err.Error(),
		})

		return
	}

	c.JSON(200, data)
}
