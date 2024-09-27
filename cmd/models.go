package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jvanrhyn/bookfans/internal"
	"gorm.io/gorm"
)

type (
	Application struct {
		GormDB *gorm.DB
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
