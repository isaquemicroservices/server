package product

import (
	"context"

	"github.com/isaqueveras/servers-microservices-backend/domain/product"
	"github.com/isaqueveras/servers-microservices-backend/infrastructure/persistence/product/grpc"
	gogrpc "google.golang.org/grpc"
)

// repository is the base struct for repository
type repository struct {
	grpcData *grpc.Product
}

// New initializes a repository
func New(ctx context.Context, conn gogrpc.ClientConnInterface) product.IProduct {
	return &repository{
		grpcData: grpc.NewProductDriver(ctx, conn),
	}
}
