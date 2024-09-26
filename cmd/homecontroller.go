package main

import "github.com/gofiber/fiber/v2"

func (h *HomeController) RegisterRoutes() {
	group := h.App.WebApp.Group("/home")
	group.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
