package grpc

import (
	"context"

	"github.com/isaqueveras/servers-microservices-backend/infrastructure/persistence/product/grpc/product"
	gogrpc "google.golang.org/grpc"
)

// Product is a base struct
type Product struct {
	client product.ProductsClient
	ctx    context.Context
}

// NewProductDriver creates a new driver for querying product data using the backend provided by the connection
func NewProductDriver(ctx context.Context, conn gogrpc.ClientConnInterface) *Product {
	return &Product{
		client: product.NewProductsClient(conn),
		ctx:    ctx,
	}
}

// GetProducts get all products of the database
func (p *Product) GetProducts() (res *product.ListProducts, err error) {
	res = new(product.ListProducts)

	if res, err = p.client.List(p.ctx, &product.Void{}); err != nil {
		return res, err
	}

	return res, nil
}
