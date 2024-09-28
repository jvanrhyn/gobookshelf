package main

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *UserController) RegisterRoutes() {
	group := h.App.WebApp.Group("/users")

	group.Get("/", h.getAllUsers)
	group.Get("/:id", h.getById)

	group.Post("/", h.createUser)
}

// Routine Handlers defined here
func (h *UserController) getAllUsers(c *fiber.Ctx) error {
	repo := UserRepository{GormDB: h.App.GormDB}
	users, err := repo.List()
	if err != nil {
		slog.Error("Failed to retrieve users", "error", err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(users)
}

func (h *UserController) getById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		slog.Error("Invalid user ID", "error", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}
	repo := UserRepository{GormDB: h.App.GormDB}
	user, err := repo.FindByID(uint(id))
	if err != nil {
		slog.Error("Failed to retrieve user", "error", err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if user == nil {
		slog.Error("User not found", "id", id)
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	return c.JSON(user)
}

func (h *UserController) createUser(c *fiber.Ctx) error {
	slog.Debug("Creating user", "request", c.Body())
	var user User
	if err := c.BodyParser(&user); err != nil {
		slog.Error("Failed to parse request body", "error", err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	repo := UserRepository{GormDB: h.App.GormDB} // Assuming h.DB is your gorm.DB instance
	if err := repo.Create(&user); err != nil {
		slog.Error("Failed to create user", "error", err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
