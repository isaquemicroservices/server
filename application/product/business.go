package product

import (
	"context"

	"github.com/isaqueveras/servers-microservices-backend/configuration"
	domain "github.com/isaqueveras/servers-microservices-backend/domain/product"
	"github.com/isaqueveras/servers-microservices-backend/infrastructure/persistence/product"
	grpc "github.com/isaqueveras/servers-microservices-backend/infrastructure/persistence/product/grpc/product"
	gogrpc "google.golang.org/grpc"
)

// GetProducts contains the logic to fetch a list of products
func GetProducts(ctx context.Context) (res ListProducts, err error) {
	var productConnection *gogrpc.ClientConn
	var data *grpc.ListProducts

	if productConnection, err = gogrpc.Dial(configuration.ProductURL, gogrpc.WithInsecure()); err != nil {
		return res, err
	}

	var repo domain.IProduct = product.New(ctx, productConnection)
	if data, err = repo.GetProducts(); err != nil {
		return res, err
	}

	for ii := range data.Products {
		res.Data = append(res.Data, Product{
			ID:          &data.Products[ii].Id,
			Name:        &data.Products[ii].Name,
			Description: &data.Products[ii].Description,
			Price:       &data.Products[ii].Price,
		})
	}

	return
}
