package main

import (
	"log/slog"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type (
	User struct {
		ID           uint           `gorm:"primaryKey;autoIncrement"`
		CreatedAt    time.Time      `gorm:"autoCreateTime"`
		UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
		DeletedAt    gorm.DeletedAt `gorm:"index"`
		Email        string         `gorm:"uniqueIndex;not null"`
		PasswordHash string         `gorm:"not null"`
		FirstName    string         `gorm:"size:100"`
		LastName     string         `gorm:"size:100"`
		// Add other fields as necessary
	}

	UserRepository struct {
		GormDB *gorm.DB
	}
)

func (repo *UserRepository) Create(user *User) error {
	return repo.GormDB.Create(user).Error
}

func (repo *UserRepository) Update(user *User) error {
	return repo.GormDB.Save(user).Error
}

func (repo *UserRepository) Delete(id uint) error {
	return repo.GormDB.Delete(&User{}, id).Error
}

func (repo *UserRepository) FindByID(id uint) (*User, error) {
	var user User
	if err := repo.GormDB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	if err := repo.GormDB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) List() ([]User, error) {
	var users []User
	if err := repo.GormDB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (h *UserController) RegisterRoutes() {
	group := h.App.WebApp.Group("/users")

	group.Get("/", func(c *fiber.Ctx) error {
		repo := UserRepository{GormDB: h.App.GormDB}
		users, err := repo.List()
		if err != nil {
			slog.Error("Failed to retrieve users", "error", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(users)
	})

	group.Get("/:id", func(c *fiber.Ctx) error {
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
		return c.JSON(user)
	})

	group.Post("/", func(c *fiber.Ctx) error {
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
	})

}
