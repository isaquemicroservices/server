package product

import (
	"context"

	domain "github.com/isaqueveras/servers-microservices-backend/domain/product"
	"github.com/isaqueveras/servers-microservices-backend/infrastructure/persistence/product/grpc"
	gogrpc "google.golang.org/grpc"
)

// repository is the base struct for repository
type repository struct {
	grpcData *grpc.Product
}

// New initializes a repository
func New(ctx context.Context, conn gogrpc.ClientConnInterface) domain.IProduct {
	return &repository{
		grpcData: grpc.NewProductDriver(ctx, conn),
	}
}

// GetProducts is a data flow manager to get all products
func (r *repository) GetProducts() (*domain.ListProducts, error) {
	return r.grpcData.GetProducts()
}

// GetDetailsProduct is a data flow manager to get the details of product
func (r *repository) GetDetailsProduct(id *int64) (*domain.Product, error) {
	return r.grpcData.GetDetailsProduct(id)
}
