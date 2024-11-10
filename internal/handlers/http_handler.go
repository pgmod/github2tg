// internal/handlers/http_handler.go
package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"telegram-welcome-bot/internal/models"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

// SetupRouter настраивает маршруты HTTP-сервера
func SetupRouter(bot *tgbotapi.BotAPI, db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Маршрут для обработки уникальных эндпоинтов
	router.POST("/:endpoint", func(c *gin.Context) {
		endpoint := c.Param("endpoint")

		// Поиск чата в базе данных по уникальному эндпоинту
		var chat models.Chat
		if err := db.First(&chat, "endpoint = ?", endpoint).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Эндпоинт не найден"})
				return
			}
			log.Printf("Ошибка при поиске эндпоинта в базе данных: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
			return
		}

		// Чтение тела запроса от GitHub
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Printf("Ошибка при чтении тела запроса: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный запрос"})
			return
		}

		// Парсинг тела запроса в структуру GitHubWebhookPayload
		var payload GitHubWebhookPayload
		if err := json.Unmarshal(body, &payload); err != nil {
			log.Printf("Ошибка при разборе JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный JSON формат"})
			return
		}
		var messageText string
		if payload.Action == "published" {
			messageText = fmt.Sprintf("%s опубликовал релиз [%s](%s):\n", payload.Sender.Login, payload.Release.TagName, payload.Release.HTMLURL)
		}

		if payload.Issue.URL != "" {
			switch payload.Action {
			case "opened":
				messageText = fmt.Sprintf("%s создал [ишью](%s):\n", payload.Sender.Login, payload.Issue.HTMLURL)
			case "reopened":
				messageText = fmt.Sprintf("%s переоткрыл [ишью](%s):\n", payload.Sender.Login, payload.Issue.HTMLURL)
			case "closed":
				messageText = fmt.Sprintf("%s закрыл [ишью](%s):\n", payload.Sender.Login, payload.Issue.HTMLURL)
			case "deleted":
				messageText = fmt.Sprintf("%s удалил [ишью](%s):\n", payload.Sender.Login, payload.Issue.HTMLURL)
			}

		} else if payload.HeadCommit.URL != "" {
			if payload.Before == "0000000000000000000000000000000000000000" {
				messageText = fmt.Sprintf("%s создал [ветку](%s) *%s*\n", payload.Sender.Login, payload.Repository.HTMLURL, payload.Ref)
			} else {

				messageText = fmt.Sprintf("%s внес [изменения](%s) в ветку *%s*:\n", payload.Sender.Login, payload.HeadCommit.URL, payload.Ref)
				for _, add := range payload.HeadCommit.Added {
					messageText += fmt.Sprintf("+ %s\n", add)
				}
				messageText += "\n"
				for _, rem := range payload.HeadCommit.Removed {
					messageText += fmt.Sprintf("- %s\n", rem)
				}
				messageText += "\n"
				for _, mod := range payload.HeadCommit.Modified {
					messageText += fmt.Sprintf("* %s\n", mod)
				}
			}
		}
		// Формирование сообщения для отправки в чат
		// messageText := fmt.Sprintf("Новые изменения в репозитории %s от %s:\n", payload.Repository.FullName, payload.Sender.Login)
		// for _, commit := range payload.Repository {
		// 	messageText += fmt.Sprintf("\n- %s: %s\n  Автор: %s\n  URL: %s\n", commit.ID[:7], commit.Message, commit.Author.Name, commit.)

		// }

		// Отправка сообщения в Telegram-чат
		msg := tgbotapi.NewMessage(chat.ChatID, messageText)
		msg.ParseMode = tgbotapi.ModeMarkdown
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Ошибка при отправке сообщения: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось отправить сообщение"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Уведомление отправлено"})
	})

	return router
}
