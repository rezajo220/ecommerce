package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rezajo220/ecommerce/internal/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product *domain.CreateProductRequest) (*domain.Product, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Product, error)
	Update(ctx context.Context, id uuid.UUID, product *domain.UpdateProductRequest) (*domain.Product, error)
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, limit, offset int) ([]domain.Product, int, error)
}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, req *domain.CreateProductRequest) (*domain.Product, error) {
	query := `
		INSERT INTO products (product_name, price, qty, brand_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, product_name, price, qty, brand_id, created_at, updated_at`

	now := time.Now()
	var product domain.Product

	err := r.db.QueryRowxContext(ctx, query, req.ProductName, req.Price, req.Qty, req.BrandID, now, now).StructScan(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	query := `
		SELECT p.id, p.product_name, p.price, p.qty, p.brand_id, p.created_at, p.updated_at, b.brand_name
		FROM products p
		LEFT JOIN brands b ON p.brand_id = b.id
		WHERE p.id = $1`

	var product domain.Product
	err := r.db.GetContext(ctx, &product, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) Update(ctx context.Context, id uuid.UUID, req *domain.UpdateProductRequest) (*domain.Product, error) {
	current, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if current == nil {
		return nil, sql.ErrNoRows
	}
	setParts := []string{}
	args := []interface{}{}
	argIndex := 1

	if req.ProductName != "" {
		setParts = append(setParts, fmt.Sprintf("product_name = $%d", argIndex))
		args = append(args, req.ProductName)
		argIndex++
	}
	if req.Price > 0 {
		setParts = append(setParts, fmt.Sprintf("price = $%d", argIndex))
		args = append(args, req.Price)
		argIndex++
	}
	if req.Qty >= 0 {
		setParts = append(setParts, fmt.Sprintf("qty = $%d", argIndex))
		args = append(args, req.Qty)
		argIndex++
	}
	if req.BrandID != uuid.Nil {
		setParts = append(setParts, fmt.Sprintf("brand_id = $%d", argIndex))
		args = append(args, req.BrandID)
		argIndex++
	}

	if len(setParts) == 0 {
		return current, nil
	}

	setParts = append(setParts, fmt.Sprintf("updated_at = $%d", argIndex))
	args = append(args, time.Now())
	argIndex++
	args = append(args, id)

	query := fmt.Sprintf(`
		UPDATE products 
		SET %s
		WHERE id = $%d
		RETURNING id, product_name, price, qty, brand_id, created_at, updated_at`,
		fmt.Sprintf("%s", setParts[0]), argIndex)

	for i := 1; i < len(setParts); i++ {
		query = fmt.Sprintf(`
		UPDATE products 
		SET %s
		WHERE id = $%d
		RETURNING id, product_name, price, qty, brand_id, created_at, updated_at`,
			fmt.Sprintf("%s, %s", setParts[0], setParts[i]), argIndex)
	}

	var product domain.Product
	err = r.db.QueryRowxContext(ctx, query, args...).StructScan(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM products WHERE id = $1`
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

func (r *productRepository) List(ctx context.Context, limit, offset int) ([]domain.Product, int, error) {
	var total int
	countQuery := `SELECT COUNT(*) FROM products`
	err := r.db.GetContext(ctx, &total, countQuery)
	if err != nil {
		return nil, 0, err
	}
	query := `
		SELECT p.id, p.product_name, p.price, p.qty, p.brand_id, p.created_at, p.updated_at, b.brand_name
		FROM products p
		LEFT JOIN brands b ON p.brand_id = b.id
		ORDER BY p.created_at DESC
		LIMIT $1 OFFSET $2`

	var products []domain.Product
	err = r.db.SelectContext(ctx, &products, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}
