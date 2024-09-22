package controller

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/jvanrhyn/bookfans/internal/data"
)

type HomeController struct {
	*fiber.App
	Database *data.Database
}

func (h *HomeController) RegisterRoutes(app *fiber.App, db *data.Database) {

	h.Database = db

	group := app.Group("/home")

	group.Get("/", func(c *fiber.Ctx) error {
		slog.Debug("Home controller called")
		return c.SendString("bookfans.online")
	})

	group.Get("/ping", func(c *fiber.Ctx) error {
		slog.Debug("Ping controller called")
		err := h.Database.Ping()
		if err != nil {
			return c.SendStatus(500)
		}

		return c.SendString("pong")
	})
}
