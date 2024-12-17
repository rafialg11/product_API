package handler

import (
	"product_api/entities"
	"product_api/services"
	"product_api/utils"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(app *fiber.App, productService services.ProductService) {
	productHandler := ProductHandler{productService: productService}
	app.Get("/products", productHandler.FindAll)
	app.Get("/products/:id", productHandler.FindById)
	app.Post("/products", productHandler.Save)
	app.Put("/products/:id", productHandler.Update)
	app.Delete("/products/:id", productHandler.Delete)
}

func (p *ProductHandler) FindAll(c *fiber.Ctx) error {
	products, err := p.productService.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to fetch products",
			Data:    nil,
			Error:   err,
		})
	}
	return c.JSON(utils.ApiResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    products,
		Error:   nil,
	})
}

func (p *ProductHandler) FindById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid product id",
			Data:    nil,
			Error:   err,
		})
	}
	product, err := p.productService.FindById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to fetch product",
			Data:    nil,
			Error:   err,
		})
	}
	return c.JSON(utils.ApiResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    product,
		Error:   nil,
	})
}

func (p *ProductHandler) Save(c *fiber.Ctx) error {
	product := new(entities.Product)
	if err := c.BodyParser(product); err != nil {
		return c.JSON(utils.ApiResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request",
			Data:    nil,
			Error:   err,
		})
	}
	product, err := p.productService.Save(product)
	if err != nil {
		return c.JSON(utils.ApiResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to save product",
			Data:    nil,
			Error:   err,
		})
	}
	return c.JSON(utils.ApiResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    product,
		Error:   nil,
	})
}

func (p *ProductHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(utils.ApiResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid product id",
			Data:    nil,
			Error:   err,
		})
	}
	product := new(entities.Product)
	product.ID = uint(id)
	if err := c.BodyParser(product); err != nil {
		return c.JSON(utils.ApiResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request",
			Data:    nil,
			Error:   err,
		})
	}
	product, err = p.productService.Update(product)
	if err != nil {
		return c.JSON(utils.ApiResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to update product",
			Data:    nil,
			Error:   err,
		})
	}
	return c.JSON(utils.ApiResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    product,
		Error:   nil,
	})
}

func (p *ProductHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(utils.ApiResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid product id",
			Data:    nil,
			Error:   err,
		})
	}
	product, err := p.productService.Delete(uint(id))
	if err != nil {
		return c.JSON(utils.ApiResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to delete product",
			Data:    nil,
			Error:   err,
		})
	}
	return c.JSON(utils.ApiResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    product,
		Error:   nil,
	})
}
