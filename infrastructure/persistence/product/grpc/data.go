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
	var prod domain.Product

	response, err := p.client.List(p.ctx, &product.Void{})
	if err != nil {
		return res, err
	}

	res.Products = make([]domain.Product, len(response.Products))
	for ii := range res.Products {
		prod.ID = &response.Products[ii].Id
		prod.Name = &response.Products[ii].Name
		prod.Description = &response.Products[ii].Description
		prod.Price = &response.Products[ii].Price

		res.Products = append(res.Products, prod)
	}

	return res, nil
}
