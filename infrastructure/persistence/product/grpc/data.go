package grpc

import (
	"context"

	domain "github.com/isaqueveras/servers-microservices-backend/domain/product"
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
func (p *Product) GetProducts() (res *domain.ListProducts, err error) {
	res = new(domain.ListProducts)

	response, err := p.client.List(p.ctx, &product.Void{})
	if err != nil {
		return res, err
	}

	res.Products = make([]domain.Product, len(response.Products))
	for i := range response.Products {
		res.Products[i] = domain.Product{
			ID:          &response.Products[i].Id,
			Name:        &response.Products[i].Name,
			Description: &response.Products[i].Description,
			Price:       &response.Products[i].Price,
		}
	}

	return
}
