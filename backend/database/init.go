package database

import (
	"errors"
	"time"

	"github.com/alifndaru/test-industrix.git/config"
	"github.com/alifndaru/test-industrix.git/models"
)

func RunMigrations() error {
	// If you prefer SQL-based migrations, run external tool.
	// Here we provide GORM auto-migrate as convenience for dev.
	db := config.DB
	if db == nil {
		return errors.New("db not initialized")
	}

	// auto migrate
	err := db.AutoMigrate(&models.Category{}, &models.Todos{})
	if err != nil {
		return err
	}
	return nil
}

func RunSeed() error {
	db := config.DB
	if db == nil {
		return errors.New("db not initialized")
	}

	// Check if categories exist
	var count int64
	db.Model(&models.Category{}).Count(&count)
	if count > 0 {
		return nil // already seeded
	}

	categories := []models.Category{
		{Name: "Work", Color: "#3B82F6", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Personal", Color: "#F97316", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Shopping", Color: "#10B981", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	if err := db.Create(&categories).Error; err != nil {
		return err
	}

	// sample todos
	todos := []models.Todos{
		{
			Title:       "Complete coding challenge",
			Description: "Build a full-stack todo application for Industrix",
			Completed:   false,
			CategoryID:  categories[0].Id,
			Priority:    "high",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "Buy groceries",
			Description: "Milk, Eggs, Bread",
			Completed:   false,
			CategoryID:  categories[2].Id,
			Priority:    "low",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	if err := db.Create(&todos).Error; err != nil {
		return err
	}

	return nil
}
