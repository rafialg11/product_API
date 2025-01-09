package handler

import (
	"context"
	"io"
	"product_api/config"
	"product_api/utils"

	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func NewOauthHandler(app fiber.Router) {
	app.Get("/oauth/google", auth)
	app.Get("/oauth/redirect", redirect)
}

func auth(c *fiber.Ctx) error {
	url := config.OAuthConfig().AuthCodeURL("state")
	return c.Redirect(url)
}

func redirect(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to get code",
			Data:    nil,
			Error:   nil,
		})
	}

	token, err := config.OAuthConfig().Exchange(c.Context(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to exchange token",
			Data:    nil,
			Error:   err,
		})
	}

	client := config.OAuthConfig().Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to get user info",
			Data:    nil,
			Error:   err,
		})
	}
	defer response.Body.Close()

	var user utils.User
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to read response",
			Data:    nil,
			Error:   err,
		})
	}
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to unmarshal json",
			Data:    nil,
			Error:   err,
		})
	}
	return c.JSON(utils.ApiResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    user,
		Error:   nil,
	})
}
