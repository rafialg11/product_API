package main

import (
	"product_api/config"
	handler "product_api/handlers"
	"product_api/repositories"
	"product_api/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.ConnectDB()

	v1 := app.Group("/api/v1")

	// Initializr Product Handler
	productRepo := repositories.NewProductRepository(config.Database)
	productService := services.NewProductService(productRepo)
	handler.NewProductHandler(v1, productService)

	//Initialize Stock Handler
	stockRepo := repositories.NewStockRepository(config.Database)
	stockService := services.NewStockService(stockRepo)
	handler.NewStockHandler(v1, stockService)

	// Start the server
	app.Listen(":3000")
}
