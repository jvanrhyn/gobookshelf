package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/jvanrhyn/bookfans/internal"
	"github.com/jvanrhyn/bookfans/internal/controller"
	"github.com/jvanrhyn/bookfans/internal/data"
)

func init() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug.Level(),
	}))

	slog.SetDefault(logger)
}

func main() {

	setupEnvironment()

	config := GetConfig()
	app := fiber.New()
	database := data.New(config)

	slog.Debug("Initializing controllers")
	controllers := []controller.ControllerInterface{
		&controller.HomeController{},
	}

	for _, c := range controllers {
		c.RegisterRoutes(app, database)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		slog.Info("Shutting down server...")

		if err := app.Shutdown(); err != nil {
			slog.Error("Server forced to shutdown: " + err.Error())
		}

		slog.Info("Server exiting")
	}()

	err := app.Listen(":8080")
	if err != nil {
		slog.Error(err.Error())
	}
}

func GetConfig() *internal.Config {
	config := internal.Config{
		ConnectionString: os.Getenv("CONNECTION_STRING"),
	}

	return &config
}

func setupEnvironment() {
	err := godotenv.Load()
	if err != nil {
		slog.Warn("Error loading .env file, falling back to OS configured variables")
	}

}
