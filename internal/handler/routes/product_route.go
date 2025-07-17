package routes

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/rezajo220/ecommerce/internal/handler"
)

func SetupProductRoutes(app *fiber.App, productHandler *handlers.ProductHandler) {
	api := app.Group("/v1/products")

	api.Post("/", productHandler.CreateProduct)
	api.Get("/", productHandler.GetProducts)
	api.Put("/:id", productHandler.UpdateProduct)
	api.Delete("/:id", productHandler.DeleteProduct)
}
