package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id" db:"id"`
	ProductName string    `json:"product_name" db:"product_name"`
	Price       float64   `json:"price" db:"price"`
	Qty         float64   `json:"qty" db:"qty"`
	BrandID     uuid.UUID `json:"brand_id" db:"brand_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	BrandName   string    `json:"brand_name,omitempty" db:"brand_name"`
}

type CreateProductRequest struct {
	ProductName string    `json:"product_name" validate:"required"`
	Price       float64   `json:"price" validate:"required,gt=0"`
	Qty         float64   `json:"qty" validate:"required,gte=0"`
	BrandID     uuid.UUID `json:"brand_id" validate:"required"`
}

type UpdateProductRequest struct {
	ProductName string    `json:"product_name,omitempty"`
	Price       float64   `json:"price,omitempty"`
	Qty         float64   `json:"qty,omitempty"`
	BrandID     uuid.UUID `json:"brand_id,omitempty"`
}

type ProductListResponse struct {
	Products   []Product `json:"products"`
	Total      int       `json:"total"`
	Page       int       `json:"page"`
	Limit      int       `json:"limit"`
	TotalPages int       `json:"total_pages"`
}
