package routes

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/rezajo220/ecommerce/internal/handler"
)

func SetupBrandRoutes(app *fiber.App, brandHandler *handlers.BrandHandler) {
	api := app.Group("/v1/brands")

	api.Post("/", brandHandler.CreateBrand)
	api.Get("/", brandHandler.GetBrands)
	api.Delete("/:id", brandHandler.DeleteBrand)
}
