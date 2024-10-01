package main

import (
	"log/slog"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
)

func recoveryMiddleware(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			slog.Error("Recovered from panic", "error", r, "stacktrace", string(debug.Stack()))
		}
	}()
	return c.Next()
}
