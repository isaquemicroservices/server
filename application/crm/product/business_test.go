package product

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/isaqueveras/servers-microservices-backend/services"
	"github.com/isaqueveras/servers-microservices-backend/utils"
	"github.com/stretchr/testify/assert"
)

var (
	ctx            context.Context = context.Background()
	ctxWithSuccess                 = time.Second * 2
	ctxWithError                   = time.Second * 0
)

// TestProduct function of tests for methods the of product domain on application
func TestProduct(t *testing.T) {
	// Initializing connections with microservices gRPC
	if err := services.InitializeConnections(); err != nil {
		log.Fatal(err)
	}

	// Initializing context with timeout for calls gRPC for context with success
	var (
		// Context duration for success
		ctxWithSuccess, cancelWithSuccess = context.WithTimeout(ctx, ctxWithSuccess)
		// Context duration for errors
		ctxWithError, cancelWithError = context.WithTimeout(ctx, ctxWithError)
	)

	// Defers for contexts
	defer cancelWithSuccess()
	defer cancelWithError()

	// TestGetAllProduct business test to gel all products with success and error
	t.Run("TestGetAllProduct", func(t *testing.T) {
		// WithSuccess business test to get all products with success
		t.Run("WithSuccess", func(t *testing.T) {
			data, err := GetProducts(ctxWithSuccess)
			assert.NoError(t, err)
			assert.NotNil(t, data.Data)
			assert.NotEmpty(t, data.Data)
		})

		// WithError business test to get all products with error
		t.Run("WithError", func(t *testing.T) {
			data, err := GetProducts(ctxWithError)

			// Validations
			assert.Error(t, err)
			assert.Nil(t, data.Data)
		})
	})

	// TestGetDetailsProduct business test to get details of products with status of success and errors
	t.Run("TestGetDetailsProduct", func(t *testing.T) {
		// WithSuccess business test to get details of product with success
		t.Run("WithSuccess", func(t *testing.T) {
			var productID int64 = 1
			var data, err = GetDetailsProduct(ctxWithSuccess, &productID)

			// Validations
			assert.Equal(t, err, nil)
			assert.NotEqual(t, data, nil)
		})

		// WithSuccess business test to get details of product with errors
		t.Run("WithError", func(t *testing.T) {
			var productID int64 = 12453586
			var data, err = GetDetailsProduct(ctxWithError, &productID)

			// Validations
			assert.Error(t, err)
			assert.Nil(t, data)
		})
	})

	// TestCreateProduct business test to get details of products with status of success and errors
	t.Run("TestCreateProduct", func(t *testing.T) {
		// Variable to errors
		var err error

		// WithSuccess business test to create product with success
		t.Run("WithSuccess", func(t *testing.T) {
			var product = &Product{
				Name:        utils.GetPointerString("Coffee"),
				Description: utils.GetPointerString("Coffee is darkly colored, bitter, slightly acidic and has a stimulating effect in humans, primarily due to its caffeine content."),
				Price:       utils.GetPointerFloat64(24.69),
			}

			// Create product
			err = CreateProduct(ctxWithSuccess, product)

			// Validations
			assert.Equal(t, *product.Name, "Coffee")
			assert.Nil(t, err)
		})

		// WithError business test to create product with errors
		t.Run("WithError", func(t *testing.T) {
			var product = &Product{
				Name:        utils.GetPointerString("Chocolate"),
				Description: utils.GetPointerString("chocolate, food product made from cocoa beans, consumed as candy and used to make beverages and to flavour or coat various confections and bakery products."),
				Price:       utils.GetPointerFloat64(24.69),
			}

			// Create product
			err = CreateProduct(ctxWithError, product)

			// Validations
			assert.Equal(t, *product.Name, "Chocolate")
			assert.NotNil(t, err)
		})
	})
}
