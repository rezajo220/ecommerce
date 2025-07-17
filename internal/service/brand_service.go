package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rezajo220/ecommerce/internal/domain"
	"github.com/rezajo220/ecommerce/internal/repository"
)

type BrandService interface {
	CreateBrand(ctx context.Context, req *domain.CreateBrandRequest) (*domain.Brand, error)
	DeleteBrand(ctx context.Context, id uuid.UUID) error
	ListBrands(ctx context.Context) ([]domain.Brand, error)
}

type brandService struct {
	brandRepo   repository.BrandRepository
	productRepo repository.ProductRepository
}

func NewBrandService(brandRepo repository.BrandRepository, productRepo repository.ProductRepository) BrandService {
	return &brandService{
		brandRepo:   brandRepo,
		productRepo: productRepo,
	}
}

func (s *brandService) CreateBrand(ctx context.Context, req *domain.CreateBrandRequest) (*domain.Brand, error) {
	return s.brandRepo.Create(ctx, req)
}

func (s *brandService) DeleteBrand(ctx context.Context, id uuid.UUID) error {
	brand, err := s.brandRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if brand == nil {
		return errors.New("brand not found")
	}

	isUsed, err := s.brandRepo.IsUsedByProducts(ctx, id)
	if err != nil {
		return err
	}
	if isUsed {
		return errors.New("cannot delete brand: it is being used by products")
	}

	return s.brandRepo.Delete(ctx, id)
}

func (s *brandService) ListBrands(ctx context.Context) ([]domain.Brand, error) {
	return s.brandRepo.List(ctx)
}
