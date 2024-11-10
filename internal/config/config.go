// internal/config/config.go
package config

import (
	"log"
	"os"
)

// Config структура для хранения конфигурации бота
type Config struct {
	BotToken string
	DBHost   string
	DBPort   string
	DBUser   string
	DBName   string
	DBPass   string
}

// LoadConfig загружает конфигурацию из переменных окружения
func LoadConfig() *Config {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN не найден в переменных окружения")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")

	return &Config{
		BotToken: botToken,
		DBHost:   dbHost,
		DBPort:   dbPort,
		DBUser:   dbUser,
		DBName:   dbName,
		DBPass:   dbPass,
	}
}
