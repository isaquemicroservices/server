package product

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/isaqueveras/servers-microservices-backend/infrastructure/persistence/crm/product"
	"github.com/isaqueveras/servers-microservices-backend/services"
	"github.com/isaqueveras/servers-microservices-backend/services/grpc"
	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	// Initializing connections with microservices gRPC
	if err := services.InitializeConnections(); err != nil {
		log.Fatal(err)
	}

	// Initializing context with timeout for calls gRPC
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	// TestGetAllProducts business test to get all products
	t.Run("TestGetAllProducts", func(t *testing.T) {
		// Get connection with microservice of products
		conn := grpc.GetProductConnection()

		// Validator to know connection is not null with gRPC
		assert.NotNil(t, conn)

		// Initializing repository of product
		var repo = product.New(ctx, conn)

		// Get products in call gRPC
		var data, err = repo.GetProducts()
		defer conn.Close()

		// Validations
		assert.Equal(t, err, nil)
		assert.NotEqual(t, data, nil)
	})
}
