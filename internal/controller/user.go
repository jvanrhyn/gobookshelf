package controller

import (
	"fmt"
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
	}
)

func (u *UserController) RegisterRoutes(app *fiber.App, db *data.Database) {

	slog.Debug("User controller registered")
	u.Database = db
	user := User{}
	group := app.Group("/user")

	group.Post("/register", func(c *fiber.Ctx) error {
		slog.Debug("User controller called")
		user.EmailAddress = c.FormValue("email")

		slog.Info("User registered with email: " + user.EmailAddress)

		return c.JSON("User registered with email: " + user.EmailAddress)
	})

	// Get user by ID

	group.Get("/:id", func(c *fiber.Ctx) error {
		slog.Debug("Getting user from database")

		id := c.Params("id")
		user, err := user.GetUserByID(id)
		if err != nil {
			slog.Error("Error retrieving user: " + err.Error())
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		return c.JSON(fiber.Map{
			"full_name": user.FirstName + " " + user.LastName,
		})
	})

}

func (u *User) GetUserByID(id string) (User, error) {
	// Pretend to get user from database
	if id == "" { // Add check for invalid ID
		return User{}, fmt.Errorf("invalid user ID")
	}
	return User{
		EmailAddress: "johndoe@example.com",
		FirstName:    "John",
		LastName:     "Doe",
		Age:          30,
	}, nil
}
