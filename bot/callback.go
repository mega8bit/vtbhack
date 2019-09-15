package main

import (
	"encoding/base64"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	tgbotapi "github.com/temamagic/telegram-bot-api"
	"strconv"
	"strings"
)

func handleCallbackQuery(Message *tgbotapi.CallbackQuery) {
	var loc *i18n.Localizer
	var markup tgbotapi.InlineKeyboardMarkup

	loc = i18n.NewLocalizer(bundle, Message.From.LanguageCode)
	user := User{}
	err := db.Find(&user, User{TelegramId: Message.From.ID}).Error
	if err != nil {
		text := fmt.Sprintf(L("no_access", *loc), Message.From.ID)
		go bot.Send(tgbotapi.NewCallback(Message.ID, text))
		return
	}
	params := strings.Split(Message.Data, " ")
	cmd := params[0]
	switch cmd {
	case "open_topic":
		go bot.Send(tgbotapi.NewCallback(Message.ID, ""))
		topic := Topic{}
		topicId, _ := strconv.Atoi(params[1])
		err := db.Find(&topic, Topic{Id: topicId}).Error
		out := ""
		buttons := []tgbotapi.InlineKeyboardButton{}
		if err == nil {
			out = loc.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "topic_view",
				TemplateData: map[string]interface{}{
					"Title":     topic.Title,
					"StartDate": topic.StartDatetime.Format("2006-01-02 15:04:05"),
				},
			})
			questions := getTopicQuestions(topicId)
			if len(questions) > 0 {
				for _, question := range questions {
					out += "\n" + fmt.Sprintf(L("topic_question", *loc), question.Title)
					buttons = []tgbotapi.InlineKeyboardButton{}
					buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s %d", question.Title, question.Id), fmt.Sprintf("open_question %d", question.Id)))
					markup.InlineKeyboard = append(markup.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
						buttons...,
					))
				}
				buttons = []tgbotapi.InlineKeyboardButton{}
			}
		} else {
			out = "No rows"
		}
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(L("back", *loc), "menu"))
		markup.InlineKeyboard = append(markup.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
			buttons...,
		))

		editmsg := tgbotapi.NewEditMessageText(Message.Message.Chat.ID, Message.Message.MessageID, out)
		editmsg.ParseMode = "HTML"
		editmsg.ReplyMarkup = &markup
		editmsg.DisableWebPagePreview = true
		bot.Send(editmsg)
	case "open_question":
		go bot.Send(tgbotapi.NewCallback(Message.ID, ""))
		question := Question{}
		questionId, _ := strconv.Atoi(params[1])
		db.Find(&question, Question{Id: questionId})
		buttons := []tgbotapi.InlineKeyboardButton{}
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("За", fmt.Sprintf("voteup %d", question.TopicId)))
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Против", fmt.Sprintf("votedown %d", question.TopicId)))
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Воздерживаюсь", fmt.Sprintf("voteno %d", question.TopicId)))
		markup.InlineKeyboard = append(markup.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
			buttons...,
		))
		buttons = []tgbotapi.InlineKeyboardButton{}
		messages := getTopicMessages(questionId)
		base64Data := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("q %d", question.Id)))
		hiddenUrl := "<a href=\"http://d.ru/" + base64Data + "\">&#8203;</a>"
		fmt.Println(hiddenUrl)
		out := hiddenUrl + fmt.Sprintf("%s", question.Title) + "\n"
		for _, message := range messages {
			out += message.Name + ": " + message.Body + "\n"
		}
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(L("back", *loc), fmt.Sprintf("open_topic %d", question.TopicId)))
		markup.InlineKeyboard = append(markup.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
			buttons...,
		))
		editmsg := tgbotapi.NewEditMessageText(Message.Message.Chat.ID, Message.Message.MessageID, out)
		editmsg.ParseMode = "HTML"
		editmsg.ReplyMarkup = &markup
		editmsg.DisableWebPagePreview = true
		bot.Send(editmsg)
	case "menu":
		go bot.Send(tgbotapi.NewCallback(Message.ID, ""))
		out, markup := getStartMenu(user, loc)
		editmsg := tgbotapi.NewEditMessageText(Message.Message.Chat.ID, Message.Message.MessageID, out)
		editmsg.ParseMode = "HTML"
		editmsg.ReplyMarkup = &markup
		editmsg.DisableWebPagePreview = true
		bot.Send(editmsg)
	default:
		go bot.Send(tgbotapi.NewCallback(Message.ID, ""))
	}
}
