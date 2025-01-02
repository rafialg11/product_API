package handler

import (
	"product_api/entities"
	"product_api/services"
	"product_api/utils"

	"github.com/gofiber/fiber/v2"
)

type StockHandler struct {
	stockService services.StockService
}

func NewStockHandler(app fiber.Router, stockService services.StockService) {
	stockHandler := StockHandler{stockService: stockService}
	app.Get("/stocks", stockHandler.FindAll)
	app.Get("/stocks/:id", stockHandler.FindById)
	app.Post("/stocks", stockHandler.Save)
	app.Put("/stocks/:id", stockHandler.Update)
	app.Delete("/stocks/:id", stockHandler.Delete)
}

func (s *StockHandler) FindAll(c *fiber.Ctx) error {
	product, err := s.stockService.FindAll()
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
		Data:    product,
		Error:   nil,
	})
}

func (s *StockHandler) FindById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid product id",
			Data:    nil,
			Error:   err,
		})
	}
	product, err := s.stockService.FindById(uint(id))
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

func (s *StockHandler) Save(c *fiber.Ctx) error {
	stock := new(entities.Stock)
	if err := c.BodyParser(stock); err != nil {
		return c.JSON(utils.ApiResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request",
			Data:    nil,
			Error:   err,
		})
	}
	stock, err := s.stockService.Save(stock)
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
		Message: "Stock saved successfully",
		Data:    stock,
		Error:   nil,
	})
}

func (s *StockHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid product id",
			Data:    nil,
			Error:   err,
		})
	}
	stock := new(entities.Stock)
	stock.ID = uint(id)
	if err := c.BodyParser(stock); err != nil {
		return c.JSON(utils.ApiResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request",
			Data:    nil,
			Error:   err,
		})
	}
	stock, err = s.stockService.Update(stock)
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
		Message: "Stock Updated Successfully",
		Data:    stock,
		Error:   nil,
	})
}

func (s *StockHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid product id",
			Data:    nil,
			Error:   err,
		})
	}
	_, err = s.stockService.Delete(uint(id))
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
		Message: "Stock deleted successfully",
		Data:    nil,
		Error:   nil,
	})
}
