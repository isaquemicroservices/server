package product

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/servers-microservices-backend/application/crm/product"
	"github.com/isaqueveras/servers-microservices-backend/configuration"
)

// getProducts godoc
// @Summary List products
// @Description List all products of the store
// @Tags Products
// @Produce json
// @Success 200 {object} product.ListProducts "List of products"
// @Router /v1/crm/products [get]
func getProducts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, configuration.Get().ContextWithTimeout)
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
// @Router /v1/crm/product/{id} [get]
func getDetailsProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, configuration.Get().ContextWithTimeout)
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

// addProduct godoc
// @Summary Create product
// @Description Create product
// @Tags Products
// @Produce json
// @Success 201 nil nil
// @Router /v1/crm/products [post]
func addProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, configuration.Get().ContextWithTimeout)
	defer cancel()

	var (
		req product.Product
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{
			"message": "Error while creating a product",
			"cause":   err.Error(),
		})

		return
	}

	if err = product.CreateProduct(ctx, &req); err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting details of product",
			"cause":   err.Error(),
		})

		return
	}

	c.JSON(201, gin.H{"message": "Product created with success"})
}
