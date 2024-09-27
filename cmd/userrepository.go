package main

import (
	"gorm.io/gorm"
	"time"
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
