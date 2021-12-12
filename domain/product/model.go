package product

// Product model of product
type Product struct {
	ID          *int64
	Name        *string
	Description *string
	Price       *float64
}

// ListProducts model a list of products
type ListProducts struct {
	Products []Product
}
