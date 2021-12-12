package product

import (
	"context"

	domain "github.com/isaqueveras/servers-microservices-backend/domain/product"
	"github.com/isaqueveras/servers-microservices-backend/infrastructure/persistence/product"
	"github.com/isaqueveras/servers-microservices-backend/services/grpc"
)

// GetProducts contains the logic to fetch a list of products
func GetProducts(ctx context.Context) (res *ListProducts, err error) {
	res = new(ListProducts)

	var (
		data *domain.ListProducts
		conn = grpc.GetProductConnection()
		repo = product.New(ctx, conn)
	)

	if data, err = repo.GetProducts(); err != nil {
		return res, err
	}

	defer conn.Close()

	for ii := range data.Products {
		res.Data = append(res.Data, Product{
			ID:          data.Products[ii].ID,
			Name:        data.Products[ii].Name,
			Description: data.Products[ii].Description,
			Price:       data.Products[ii].Price,
		})
	}

	return
}
