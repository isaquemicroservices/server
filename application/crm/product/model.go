package product

import "time"

// Product model of product
type Product struct {
	ID             *int64     `json:"id,omitempty"`
	Name           *string    `json:"name,omitempty"`
	Description    *string    `json:"description,omitempty"`
	Price          *float64   `json:"price,omitempty"`
	Amount         *int64     `json:"amount,omitempty"`
	ExpirationDate *time.Time `json:"expiration_date,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}

// ListProducts model for list products
type ListProducts struct {
	Data []Product `json:"products,omitempty"`
}
