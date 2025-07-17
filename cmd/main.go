package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/rezajo220/ecommerce/docs"
	handlers "github.com/rezajo220/ecommerce/internal/handler"
	"github.com/rezajo220/ecommerce/internal/handler/routes"
	"github.com/rezajo220/ecommerce/internal/repository"
	services "github.com/rezajo220/ecommerce/internal/service"
	fiberSwagger "github.com/swaggo/fiber-swagger"
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

	app := fiber.New(fiber.Config{
		AppName:      "Product Service",
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
			}
			return nil
		},
	})

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: false,
	}))
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

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

	routes.SetupProductRoutes(app, productHandler)
	routes.SetupBrandRoutes(app, brandHandler)

	log.Printf("Server starting on port %s", cfg.Server.Port)
	log.Printf("Swagger documentation available at: http://localhost:%s/swagger/", cfg.Server.Port)
	log.Fatal(app.Listen(":" + cfg.Server.Port))
}
