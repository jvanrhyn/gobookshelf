package controller

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/jvanrhyn/bookfans/internal/data"
)

type (
	User struct {
		EmailAddress string `validate:"required,email"`
		FirstName    string `validate:"required"`
		LastName     string `validate:"required"`
		Age          int    `validate:"required,gte=13,lte=100"`
	}

	UserController struct {
		*fiber.App
		Database *data.Database

		User User
	}
)

func (u *UserController) RegisterRoutes(app *fiber.App, db *data.Database) {

	slog.Debug("Home controller registered")
	u.Database = db
	u.User = User{}
	group := app.Group("/user")

	group.Post("/register", func(c *fiber.Ctx) error {
		slog.Debug("User controller called")
		u.User.EmailAddress = c.FormValue("email")

		slog.Info("User registered with email: " + u.User.EmailAddress)

		return c.SendString("User registered with email: " + u.User.EmailAddress)
	})

	// Get user by ID

	group.Get("/:id", func(c *fiber.Ctx) error {
		slog.Debug("Getting user from database")
		return c.SendString(u.User.FirstName + " " + u.User.LastName)
	})
}
