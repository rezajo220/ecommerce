package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rezajo220/ecommerce/internal/domain"
	services "github.com/rezajo220/ecommerce/internal/service"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the provided information
// @Tags products
// @Accept json
// @Produce json
// @Param product body domain.CreateProductRequest true "Product information"
// @Success 201 {object} domain.ProductResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var req domain.CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	product, err := h.productService.CreateProduct(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
		"data":    product,
	})
}

// GetProducts godoc
// @Summary Get products with pagination
// @Description Get a list of products with pagination support
// @Tags products
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} domain.ProductListResponseWrapper
// @Failure 500 {object} domain.ErrorResponse
// @Router /products [get]
func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	response, err := h.productService.ListProducts(c.Context(), page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Products retrieved successfully",
		"data":    response,
	})
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update an existing product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID (UUID)"
// @Param product body domain.UpdateProductRequest true "Updated product information"
// @Success 200 {object} domain.ProductResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	var req domain.UpdateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	product, err := h.productService.UpdateProduct(c.Context(), id, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product updated successfully",
		"data":    product,
	})
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete an existing product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID (UUID)"
// @Success 200 {object} domain.MessageResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	if err := h.productService.DeleteProduct(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}
