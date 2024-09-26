package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/jvanrhyn/bookfans/internal"
	"github.com/jvanrhyn/bookfans/internal/controller"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Application struct {
		DB     *gorm.DB
		Config *internal.Config
		WebApp *fiber.App
	}

	HomeController struct {
		App *Application
	}

	UserController struct {
		App *Application
	}
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug.Level(),
	}))
	slog.SetDefault(logger)
}

func main() {
	slog.Debug("Setting up application")

	application := &Application{
		WebApp: fiber.New(),
	}

	if err := setupEnvironment(); err != nil {
		slog.Warn("Error encountered during environment setup", "error", err)
	}

	application.startUp()
	if err := application.configureDatabase(); err != nil {
		slog.Error("Failed to configure database", "error", err)
	}

	slog.Debug("Initializing controllers")
	controllers := []controller.ControllerInterface{
		&HomeController{App: application},
		&UserController{App: application},
	}

	for _, c := range controllers {
		c.RegisterRoutes()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		slog.Info("Shutting down server...")

		if err := application.WebApp.Shutdown(); err != nil {
			slog.Error("Server forced to shutdown", "error", err)
		}

		slog.Info("Server exiting")
	}()

	if err := application.WebApp.Listen(fmt.Sprintf(":%s", application.Config.HostPort)); err != nil {
		slog.Error("Error starting server", "error", err)
	}
}

func (a *Application) startUp() {
	slog.Info("Starting up Bookfans")
	config, err := getConfig() // Capture error from getConfig
	if err != nil {
		slog.Error("Error loading configuration", "error", err)
		return // Early exit if there is an error
	}
	a.Config = config
}

func getConfig() (*internal.Config, error) {
	connectionString := os.Getenv("CONNECTION_STRING")
	hostPort := os.Getenv("HOST_PORT")

	// Check for necessary environment variables
	if connectionString == "" || hostPort == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return &internal.Config{
		ConnectionString: connectionString,
		HostPort:         hostPort,
	}, nil
}

func (a *Application) configureDatabase() error {
	db, err := gorm.Open(postgres.Open(a.Config.ConnectionString), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("database connection error: %w", err)
	}

	a.DB = db
	slog.Info("Database connection established successfully")

	a.DB.AutoMigrate(&User{})

	return nil
}

func setupEnvironment() error {
	slog.Debug("Setting up environment")
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}
