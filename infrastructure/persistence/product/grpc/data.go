package grpc

import (
	"context"

	"github.com/isaqueveras/servers-microservices-backend/infrastructure/persistence/product/grpc/product"
	gogrpc "google.golang.org/grpc"
)

type Product struct {
	client product.ProductsClient
	ctx    context.Context
}

func NewProductDriver(ctx context.Context, conn gogrpc.ClientConnInterface) *Product {
	return &Product{
		client: product.NewProductsClient(conn),
		ctx:    ctx,
	}
}
