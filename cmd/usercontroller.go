package main

import (
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
		DB *gorm.DB
	}
)

func (repo *UserRepository) Create(user *User) error {
	return repo.DB.Create(user).Error
}

func (repo *UserRepository) Update(user *User) error {
	return repo.DB.Save(user).Error
}

func (repo *UserRepository) Delete(id uint) error {
	return repo.DB.Delete(&User{}, id).Error
}

func (repo *UserRepository) FindByID(id uint) (*User, error) {
	var user User
	if err := repo.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	if err := repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) List() ([]User, error) {
	var users []User
	if err := repo.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (h *UserController) RegisterRoutes() {
	group := h.App.WebApp.Group("/users")
	group.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Example routes for CRUD operations
	group.Post("/", func(c *fiber.Ctx) error {
		var user User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		repo := UserRepository{DB: h.App.DB} // Assuming h.DB is your gorm.DB instance
		if err := repo.Create(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.Status(fiber.StatusCreated).JSON(user)
	})

	// Implement other CRUD routes (PUT, DELETE, GET by ID, etc.) similarly
}
