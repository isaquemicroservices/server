package product

// IProduct interface of methods to product
type IProduct interface {
	GetProducts() (*ListProducts, error)
	GetDetailsProduct(*int64) (*Product, error)
}
