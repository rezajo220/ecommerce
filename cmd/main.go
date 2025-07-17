package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/rezajo220/ecommerce/docs"
	handlers "github.com/rezajo220/ecommerce/internal/handler"
	"github.com/rezajo220/ecommerce/internal/handler/routes"
	"github.com/rezajo220/ecommerce/internal/repository"
	services "github.com/rezajo220/ecommerce/internal/service"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title E-commerce API
// @version 1.0

// @host localhost:8000
// @BasePath /v1
// @schemes http https

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: false,
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	pDB, err := NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pDB.Close()

	productRepository := repository.NewProductRepository(pDB)
	brandRepository := repository.NewBrandRepository(pDB)

	productService := services.NewProductService(productRepository, brandRepository)
	brandService := services.NewBrandService(brandRepository, productRepository)

	productHandler := handlers.NewProductHandler(productService)
	brandHandler := handlers.NewBrandHandler(brandService)

	routes.SetupProductRoutes(e, productHandler)
	routes.SetupBrandRoutes(e, brandHandler)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "ok",
			"service": "e-commerce-api",
		})
	})

	log.Printf("Server starting on port %s", cfg.Server.Port)
	log.Printf("Swagger documentation available at: http://localhost:%s/swagger/", cfg.Server.Port)
	e.Logger.Fatal(e.Start(":" + cfg.Server.Port))
}
