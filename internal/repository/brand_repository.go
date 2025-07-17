package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rezajo220/ecommerce/internal/domain"
)

type BrandRepository interface {
	Create(ctx context.Context, brand *domain.CreateBrandRequest) (*domain.Brand, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Brand, error)
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context) ([]domain.Brand, error)
	IsUsedByProducts(ctx context.Context, id uuid.UUID) (bool, error)
}

type brandRepository struct {
	db *sqlx.DB
}

func NewBrandRepository(db *sqlx.DB) BrandRepository {
	return &brandRepository{db: db}
}

func (r *brandRepository) Create(ctx context.Context, req *domain.CreateBrandRequest) (*domain.Brand, error) {
	query := `
		INSERT INTO brands (brand_name, created_at, updated_at)
		VALUES ($1, $2, $3)
		RETURNING id, brand_name, created_at, updated_at`

	now := time.Now()
	var brand domain.Brand

	err := r.db.QueryRowxContext(ctx, query, req.BrandName, now, now).StructScan(&brand)
	if err != nil {
		return nil, err
	}

	return &brand, nil
}

func (r *brandRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Brand, error) {
	query := `
		SELECT id, brand_name, created_at, updated_at
		FROM brands
		WHERE id = $1`

	var brand domain.Brand
	err := r.db.GetContext(ctx, &brand, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &brand, nil
}

func (r *brandRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM brands WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *brandRepository) List(ctx context.Context) ([]domain.Brand, error) {
	query := `
		SELECT id, brand_name, created_at, updated_at
		FROM brands
		ORDER BY brand_name ASC`

	var brands []domain.Brand
	err := r.db.SelectContext(ctx, &brands, query)
	return brands, err
}

func (r *brandRepository) IsUsedByProducts(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM products WHERE brand_id = $1`
	err := r.db.GetContext(ctx, &count, query, id)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
