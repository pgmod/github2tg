// internal/database/db.go
package database

import (
	"fmt"
	"log"
	"telegram-welcome-bot/internal/config"
	"telegram-welcome-bot/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB инициализирует подключение к базе данных и проводит миграции
func InitDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPass)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Автоматическая миграция таблицы для модели Chat
	if err := db.AutoMigrate(&models.Chat{}); err != nil {
		log.Fatalf("Ошибка при миграции базы данных: %v", err)
	}

	return db
}
