package services

import (
	"context"
	"errors"
	"math"

	"github.com/google/uuid"
	"github.com/rezajo220/ecommerce/internal/domain"
	"github.com/rezajo220/ecommerce/internal/repository"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req *domain.CreateProductRequest) (*domain.Product, error)
	GetProduct(ctx context.Context, id uuid.UUID) (*domain.Product, error)
	UpdateProduct(ctx context.Context, id uuid.UUID, req *domain.UpdateProductRequest) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id uuid.UUID) error
	ListProducts(ctx context.Context, page, limit int) (*domain.ProductListResponse, error)
}

type productService struct {
	productRepo repository.ProductRepository
	brandRepo   repository.BrandRepository
}

func NewProductService(productRepo repository.ProductRepository, brandRepo repository.BrandRepository) ProductService {
	return &productService{
		productRepo: productRepo,
		brandRepo:   brandRepo,
	}
}

func (s *productService) CreateProduct(ctx context.Context, req *domain.CreateProductRequest) (*domain.Product, error) {
	brand, err := s.brandRepo.GetByID(ctx, req.BrandID)
	if err != nil {
		return nil, err
	}
	if brand == nil {
		return nil, errors.New("brand not found")
	}

	return s.productRepo.Create(ctx, req)
}

func (s *productService) GetProduct(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	product, err := s.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (s *productService) UpdateProduct(ctx context.Context, id uuid.UUID, req *domain.UpdateProductRequest) (*domain.Product, error) {
	if req.BrandID != uuid.Nil {
		brand, err := s.brandRepo.GetByID(ctx, req.BrandID)
		if err != nil {
			return nil, err
		}
		if brand == nil {
			return nil, errors.New("brand not found")
		}
	}

	return s.productRepo.Update(ctx, id, req)
}

func (s *productService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	product, err := s.productRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("product not found")
	}

	return s.productRepo.Delete(ctx, id)
}

func (s *productService) ListProducts(ctx context.Context, page, limit int) (*domain.ProductListResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit
	products, total, err := s.productRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &domain.ProductListResponse{
		Products:   products,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}
