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

func main() {
	// Configure logging
	configureLogging()

	// Set up the environment
	if err := setupEnvironment(); err != nil {
		slog.Warn("Error encountered during environment setup", "error", err)
		os.Exit(1)
	}

	// Set up the application
	application := &Application{
		WebApp: fiber.New(),
		Config: getConfig(),
	}

	// Configure the database on the application
	if err := application.configureDatabase(); err != nil {
		slog.Error("Failed to configure database", "error", err)
		os.Exit(1)
	}

	mapControllers(application)
	hostApplication(application)
}

// mapControllers initializes and registers routes for the application's controllers.
// It takes a pointer to an Application struct and creates a slice of controllers,
// each of which is initialized with the given application instance. The function
// then iterates over the controllers and calls their RegisterRoutes method to
// set up the necessary routes.
func mapControllers(application *Application) {
	slog.Debug("Initializing controllers")
	controllers := []controller.ControllerInterface{
		&HomeController{App: application},
		&UserController{App: application},
	}

	for _, c := range controllers {
		c.RegisterRoutes()
	}
}

// hostApplication starts the web application and listens for system signals to gracefully shut down the server.
// It takes an Application pointer as an argument, which contains the web application and its configuration.
// The function sets up a channel to listen for SIGINT and SIGTERM signals, and upon receiving one, it logs the shutdown process,
// attempts to gracefully shut down the web application, and logs any errors encountered during the shutdown.
// If the web application fails to start, it logs the error.
func hostApplication(application *Application) {
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

func getConfig() *internal.Config {
	connectionString := os.Getenv("CONNECTION_STRING")
	hostPort := os.Getenv("HOST_PORT")

	// Check for necessary environment variables
	if connectionString == "" || hostPort == "" {
		panic("missing required environment variables")
	}

	return &internal.Config{
		ConnectionString: connectionString,
		HostPort:         hostPort,
	}
}

// configureDatabase establishes a connection to the database using the connection string
// provided in the application's configuration. It initializes the GormDB field of the
// Application struct with the connected database instance. If the connection is successful,
// it logs a message indicating the success and performs auto-migration for the User model.
// If there is an error during the connection process, it returns an error with a descriptive message.
func (a *Application) configureDatabase() error {
	db, err := gorm.Open(postgres.Open(a.Config.ConnectionString), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("database connection error: %w", err)
	}

	a.GormDB = db
	slog.Info("Database connection established successfully")

	a.GormDB.AutoMigrate(&User{})

	return nil
}

// setupEnvironment loads environment variables from a .env file using the godotenv package.
// It logs a debug message indicating the start of the setup process.
// If the .env file cannot be loaded, it returns an error with a descriptive message.
// Returns nil if the environment variables are loaded successfully.
func setupEnvironment() error {
	slog.Debug("Setting up environment")
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

// configureLogging sets up the logging configuration for the application.
// It creates a new logger with a JSON handler that outputs to stdout and sets the log level to debug.
// The created logger is then set as the default logger.
func configureLogging() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug.Level(),
	}))
	slog.SetDefault(logger)
}
