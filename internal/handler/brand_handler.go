package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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
func (h *BrandHandler) CreateBrand(c echo.Context) error {
	var req domain.CreateBrandRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	brand, err := h.brandService.CreateBrand(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
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
func (h *BrandHandler) GetBrands(c echo.Context) error {
	brands, err := h.brandService.ListBrands(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
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
func (h *BrandHandler) DeleteBrand(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid brand ID",
		})
	}

	if err := h.brandService.DeleteBrand(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Brand deleted successfully",
	})
}
