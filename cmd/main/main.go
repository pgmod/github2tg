// cmd/bot/main.go
package main

import (
	"log"
	"sync"
	"telegram-welcome-bot/internal/config"
	"telegram-welcome-bot/internal/database"
	"telegram-welcome-bot/internal/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Инициализация базы данных
	db := database.InitDB(cfg)

	// Создание нового экземпляра бота с использованием токена
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Fatalf("Ошибка при создании бота: %v", err)
	}

	bot.Debug = true // Включаем режим отладки

	log.Printf("Бот %s успешно запущен", bot.Self.UserName)

	// Настройка HTTP-сервера и передача экземпляра бота
	router := handlers.SetupRouter(bot, db)

	// Основной цикл обработки обновлений
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		// Запуск HTTP-сервера
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Ошибка при запуске HTTP-сервера: %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		// Создаем канал для получения обновлений
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60
		updates := bot.GetUpdatesChan(u)

		// Обработка обновлений
		for update := range updates {
			if update.MyChatMember != nil {
				handlers.HandleNewChatMember(bot, update.MyChatMember, db)
			}
		}
	}()

	wg.Wait()
}
