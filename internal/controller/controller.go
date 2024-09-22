package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jvanrhyn/bookfans/internal/data"
)

type (
	ControllerInterface interface {
		RegisterRoutes(app *fiber.App, db *data.Database)
	}
)
