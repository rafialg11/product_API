package handler

import (
	"product_api/utils"

	"github.com/gofiber/fiber/v2"
)

type StockHandler struct {
}

func NewStockHandler(app *fiber.App) {
	stockHandler := StockHandler{}
	app.Get("/stocks", stockHandler.FindAll)
	app.Get("/stocks/:id", stockHandler.FindById)
	app.Post("/stocks", stockHandler.Save)
	app.Put("/stocks/:id", stockHandler.Update)
	app.Delete("/stocks/:id", stockHandler.Delete)
}

func (s *StockHandler) FindAll(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    nil,
		Error:   nil,
	})
}

func (s *StockHandler) FindById(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    nil,
		Error:   nil,
	})
}

func (s *StockHandler) Save(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    nil,
		Error:   nil,
	})
}

func (s *StockHandler) Update(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    nil,
		Error:   nil,
	})
}

func (s *StockHandler) Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(utils.ApiResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    nil,
		Error:   nil,
	})
}
