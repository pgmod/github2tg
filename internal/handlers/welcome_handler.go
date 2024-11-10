// internal/handlers/welcome_handler.go
package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"telegram-welcome-bot/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

// HandleNewChatMember обрабатывает событие добавления бота в чат
func HandleNewChatMember(bot *tgbotapi.BotAPI, chatMember *tgbotapi.ChatMemberUpdated, db *gorm.DB) {
	// Проверяем, что бот добавлен в чат
	if chatMember.NewChatMember.Status == "member" {
		chatID := chatMember.Chat.ID

		// Создаем хэш на основе chatID
		hash := sha256.Sum256([]byte(fmt.Sprintf("%d", chatID)))
		endpoint := hex.EncodeToString(hash[:])[:16] // Укорачиваем хэш до 16 символов

		// Проверяем, существует ли уже запись для данного чата
		var chat models.Chat
		result := db.First(&chat, "chat_id = ?", chatID)
		if result.Error == gorm.ErrRecordNotFound {
			// Создаем новую запись, если чата еще нет в базе данных
			chat = models.Chat{ChatID: chatID, Endpoint: endpoint}
			if err := db.Create(&chat).Error; err != nil {
				log.Printf("Ошибка при сохранении чата в базе данных: %v", err)
				return
			}
			log.Printf("Новый эндпоинт создан для чата %d: /%s", chatID, endpoint)
		} else if result.Error != nil {
			log.Printf("Ошибка при поиске чата в базе данных: %v", result.Error)
			return
		}

		// Отправка приветственного сообщения
		msg := tgbotapi.NewMessage(chatID, "Привет! Спасибо, что добавили меня в чат!\nВаш эндпоинт: "+os.Getenv("HOST")+"/"+endpoint)
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Ошибка при отправке сообщения: %v", err)
		}
	}
}
