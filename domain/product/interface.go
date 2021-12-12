package product

import "github.com/isaqueveras/servers-microservices-backend/infrastructure/persistence/product/grpc/product"

// IProduct interface of methods to product
type IProduct interface {
	GetProducts() (*product.ListProducts, error)
}
