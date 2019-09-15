package main

import (
	"encoding/base64"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	tgbotapi "github.com/temamagic/telegram-bot-api"
	"strconv"
	"strings"
)

func handleMessage(Message *tgbotapi.Message) {
	var loc *i18n.Localizer
	loc = i18n.NewLocalizer(bundle, Message.From.LanguageCode)
	user := User{}
	err := db.Find(&user, User{TelegramId: Message.From.ID}).Error
	if err != nil {
		text := fmt.Sprintf(L("no_access", *loc), Message.From.ID)
		bot.Send(tgbotapi.NewMessage(Message.Chat.ID, text))
		return
	}

	switch {
	case Message.IsCommand():
		switch strings.ToLower(Message.Command()) {
		case "start":
			start(user, Message, loc)
		default:
			start(user, Message, loc)
		}
	case Message.ReplyToMessage.Text != "":
		if len(Message.ReplyToMessage.Entities) > 0 {
			hiddenUrl := Message.ReplyToMessage.Entities[0].URL
			fmt.Println(hiddenUrl)
			if strings.Contains(hiddenUrl, "http://d.ru/") {
				// наша скрытая ссылка
				hiddenContent := strings.TrimPrefix(hiddenUrl, "http://d.ru/")
				data, _ := base64.StdEncoding.DecodeString(hiddenContent)
				params := strings.Split(fmt.Sprintf("%s", data), " ")
				switch params[0] {
				case "q":
					var markup tgbotapi.InlineKeyboardMarkup
					txtOut := ""
					question := Question{}
					questionId, _ := strconv.Atoi(params[1])
					db.Find(&question, Question{Id: questionId})
					if Message.Text != "" {
						message := DbMessage{
							Body:       Message.Text,
							QuestionId: questionId,
							UserId:     user.Id,
						}
						err = db.Create(&message).Error
						if err == nil {
							buttons := []tgbotapi.InlineKeyboardButton{}
							messages := getTopicMessages(questionId)
							base64Data := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("q %d", question.Id)))
							hiddenUrl := "<a href=\"http://d.ru/" + base64Data + "\">&#8203;</a>"
							fmt.Println(hiddenUrl)
							txtOut = hiddenUrl + fmt.Sprintf("%s %d", question.Title, question.Id) + "\n"
							for _, message := range messages {
								txtOut += message.Name + ": " + message.Body + "\n"
							}
							buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(L("back", *loc), fmt.Sprintf("open_topic %d", question.TopicId)))
							markup.InlineKeyboard = append(markup.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
								buttons...,
							))
						} else {
							markup.InlineKeyboard = make([][]tgbotapi.InlineKeyboardButton, 0)
							txtOut = err.Error()
						}
					}
					msg := tgbotapi.NewMessage(Message.Chat.ID, txtOut)
					msg.ReplyMarkup = markup
					msg.ParseMode = "HTML"
					bot.Send(msg)
				default:
					msg := tgbotapi.NewMessage(Message.Chat.ID, params[0])
					bot.Send(msg)
				}
			}
		}
	case Message.Text != "":
		go bot.Send(tgbotapi.NewChatAction(Message.Chat.ID, "typing"))
		start(user, Message, loc)
	default:
		start(user, Message, loc)
	}

}
