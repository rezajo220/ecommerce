package routes

import (
	"github.com/labstack/echo/v4"
	handlers "github.com/rezajo220/ecommerce/internal/handler"
)

func SetupProductRoutes(e *echo.Echo, productHandler *handlers.ProductHandler) {
	api := e.Group("/v1/products")

	api.POST("/", productHandler.CreateProduct)
	api.GET("/", productHandler.GetProducts)
	api.PUT("/:id", productHandler.UpdateProduct)
	api.DELETE("/:id", productHandler.DeleteProduct)
}
