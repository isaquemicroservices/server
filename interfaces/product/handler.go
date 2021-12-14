package product

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/servers-microservices-backend/application/product"
	"github.com/isaqueveras/servers-microservices-backend/configuration"
)

// getProducts godoc
// @Summary List products
// @Description List all products of the store
// @Tags Products
// @Produce json
// @Success 200 {object} product.ListProducts "List of products"
// @Router /products [get]
func getProducts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, configuration.ContextWithTimeout)
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

// getDetailsProduct godoc
// @Summary Detials of product
// @Description List all detials of product
// @Tags Products
// @Produce json
// @Success 200 {object} product.Product "Details of product"
// @Router /product/{id} [get]
func getDetailsProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, configuration.ContextWithTimeout)
	defer cancel()

	productID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Parameter format is invalid",
			"cause":   err.Error(),
		})

		return
	}

	data, err := product.GetDetailsProduct(ctx, &productID)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting details of product",
			"cause":   err.Error(),
		})

		return
	}

	c.JSON(200, data)
}
