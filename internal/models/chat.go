// internal/models/chat.go
package models

import "gorm.io/gorm"

// Chat структура для хранения информации о чате и его уникальном идентификаторе
type Chat struct {
	gorm.Model
	ChatID   int64  `gorm:"uniqueIndex"` // Уникальный идентификатор чата
	Endpoint string `gorm:"uniqueIndex"` // Уникальный эндпоинт для чата
}
