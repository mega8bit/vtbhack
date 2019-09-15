package main

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	tgbotapi "github.com/temamagic/telegram-bot-api"
)

func start(user User, Message *tgbotapi.Message, Loc *i18n.Localizer) {
	out, markup := getStartMenu(user, Loc)
	msg := tgbotapi.NewMessage(Message.Chat.ID, out)
	msg.ReplyMarkup = markup
	msg.ParseMode = "HTML"
	bot.Send(msg)
}
