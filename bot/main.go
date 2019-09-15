package main

import tgbotapi "github.com/temamagic/telegram-bot-api"

func main() {
	defer db.Close()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			go handleMessage(update.Message)
		}
		if update.CallbackQuery != nil {
			go handleCallbackQuery(update.CallbackQuery)
		}
	}
}
