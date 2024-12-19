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

	// Initializr Product Handler
	productRepo := repositories.NewProductRepository(config.Database)
	productService := services.NewProductService(productRepo)
	handler.NewProductHandler(app, productService)

	// Start the server
	app.Listen(":3000")
}
