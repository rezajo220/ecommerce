package domain

import (
	"time"

	"github.com/google/uuid"
)

type Brand struct {
	ID        uuid.UUID `json:"id" db:"id"`
	BrandName string    `json:"brand_name" db:"brand_name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreateBrandRequest struct {
	BrandName string `json:"brand_name" validate:"required"`
}
