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

// RegisterRoutes sets up the routes for the HomeController.
// It registers the following endpoints:
// - GET /home: Returns a simple string "bookfans.online".
// - GET /home/ping: Pings the database and returns "pong" if successful, otherwise returns a 500 status code.
//
// Parameters:
// - app: The Fiber application instance.
// - db: The database instance to be used by the controller.
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
