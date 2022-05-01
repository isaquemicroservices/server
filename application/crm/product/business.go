package product

import (
	"context"

	domain "github.com/isaqueveras/servers-microservices-backend/domain/crm/product"
	"github.com/isaqueveras/servers-microservices-backend/infrastructure/persistence/crm/product"
	"github.com/isaqueveras/servers-microservices-backend/oops"
	"github.com/isaqueveras/servers-microservices-backend/services/grpc"
)

// GetProducts contains the logic to fetch a list of products
func GetProducts(ctx context.Context) (res *ListProducts, err error) {
	const errorMessage string = "Error getting products"
	res = new(ListProducts)

	var (
		data *domain.ListProducts
		conn = grpc.GetProductConnection()
	)

	repo := product.New(ctx, conn)
	if data, err = repo.GetProducts(); err != nil {
		return nil, oops.Wrap(err, errorMessage)
	}

	defer conn.Close()

	res.Data = make([]Product, len(data.Products))
	for i := range data.Products {
		res.Data[i] = Product{
			ID:          data.Products[i].ID,
			Name:        data.Products[i].Name,
			Description: data.Products[i].Description,
			Price:       data.Products[i].Price,
		}
	}

	return
}

// GetDetailsProduct contains the logic to fetch the details of product
func GetDetailsProduct(ctx context.Context, id *int64) (*Product, error) {
	const errorMessage string = "Error getting product details"
	var (
		data *domain.Product
		conn = grpc.GetProductConnection()
	)

	repo := product.New(ctx, conn)
	data, err := repo.GetDetailsProduct(id)
	if err != nil {
		return nil, oops.Wrap(err, errorMessage)
	}

	defer conn.Close()

	return (*Product)(data), nil
}

// CreateProduct contains the logic to create a product
func CreateProduct(ctx context.Context, in *Product) error {
	const errorMessage string = "Error when registering product"
	var (
		conn = grpc.GetProductConnection()
		repo = product.New(ctx, conn)
		prod = &domain.Product{
			ID:          in.ID,
			Name:        in.Name,
			Description: in.Description,
			Price:       in.Price,
		}
	)

	if err := repo.CreateProduct(prod); err != nil {
		return oops.Wrap(err, errorMessage)
	}

	defer conn.Close()

	return nil
}

// ListAllProductsWithMinimumQuantity contains the logic to fetch a list of products with minimum quantity
func ListAllProductsWithMinimumQuantity(ctx context.Context) (res *ListProducts, err error) {
	const errorMessage string = "Error listing all products with minimum quantity"
	res = new(ListProducts)

	var (
		data *domain.ListProducts
		conn = grpc.GetProductConnection()
	)

	repo := product.New(ctx, conn)
	if data, err = repo.ListAllProductsWithMinimumQuantity(); err != nil {
		return nil, oops.Wrap(err, errorMessage)
	}

	defer conn.Close()

	res.Data = make([]Product, len(data.Products))
	for i := range data.Products {
		res.Data[i] = Product{
			ID:     data.Products[i].ID,
			Name:   data.Products[i].Name,
			Amount: data.Products[i].Amount,
		}
	}

	return
}
