package main

import "github.com/gofiber/fiber/v2"

func (h *HomeController) RegisterRoutes() {
	group := h.App.WebApp.Group("/home")
	group.Get("/", h.get)

}

func (h *HomeController) get(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
