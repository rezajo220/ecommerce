package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rezajo220/ecommerce/internal/domain"
	services "github.com/rezajo220/ecommerce/internal/service"
)

type BrandHandler struct {
	brandService services.BrandService
}

func NewBrandHandler(brandService services.BrandService) *BrandHandler {
	return &BrandHandler{brandService: brandService}
}

// CreateBrand godoc
// @Summary Create a new brand
// @Description Create a new brand with the provided information
// @Tags brands
// @Accept json
// @Produce json
// @Param brand body domain.CreateBrandRequest true "Brand information"
// @Success 201 {object} domain.BrandResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /brands [post]
func (h *BrandHandler) CreateBrand(c *fiber.Ctx) error {
	var req domain.CreateBrandRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	brand, err := h.brandService.CreateBrand(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Brand created successfully",
		"data":    brand,
	})
}

// GetBrands godoc
// @Summary Get all brands
// @Description Get a list of all brands
// @Tags brands
// @Accept json
// @Produce json
// @Success 200 {object} domain.BrandListResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /brands [get]
func (h *BrandHandler) GetBrands(c *fiber.Ctx) error {
	brands, err := h.brandService.ListBrands(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Brands retrieved successfully",
		"data":    brands,
	})
}

// DeleteBrand godoc
// @Summary Delete a brand
// @Description Delete an existing brand by ID (only if not used by products)
// @Tags brands
// @Accept json
// @Produce json
// @Param id path string true "Brand ID (UUID)"
// @Success 200 {object} domain.MessageResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 409 {object} domain.ErrorResponse "Brand is being used by products"
// @Failure 500 {object} domain.ErrorResponse
// @Router /brands/{id} [delete]
func (h *BrandHandler) DeleteBrand(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid brand ID",
		})
	}

	if err := h.brandService.DeleteBrand(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Brand deleted successfully",
	})
}
