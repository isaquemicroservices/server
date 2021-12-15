package product

// Product model of product
type Product struct {
	ID          *int64   `json:"id,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price,omitempty"`
}

// ListProducts model for list products
type ListProducts struct {
	Data []Product `json:"products,omitempty"`
}
