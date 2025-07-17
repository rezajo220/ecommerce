package routes

import (
	"github.com/labstack/echo/v4"
	handlers "github.com/rezajo220/ecommerce/internal/handler"
)

func SetupBrandRoutes(e *echo.Echo, brandHandler *handlers.BrandHandler) {
	api := e.Group("/v1/brands")

	api.POST("/", brandHandler.CreateBrand)
	api.GET("/", brandHandler.GetBrands)
	api.DELETE("/:id", brandHandler.DeleteBrand)
}
