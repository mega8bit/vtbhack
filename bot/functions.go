package main

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	tgbotapi "github.com/temamagic/telegram-bot-api"
	"gitlab.com/vtbhack/bot.git/API"
)

func getStartMenu(user User, Loc *i18n.Localizer) (out string, markup tgbotapi.InlineKeyboardMarkup) {
	out = Loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "command_start",
		TemplateData: map[string]interface{}{
			"UserFirstName": user.Name,
		},
	})

	topics, total := getActiveTopics(user.Token)
	if total > 0 {
		out += "\n" + fmt.Sprintf(L("main_current_topics", *Loc), total)
		for _, topic := range topics {
			buttons := []tgbotapi.InlineKeyboardButton{}
			buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(topic.Title, fmt.Sprintf("open_topic %d", topic.Id)))
			markup.InlineKeyboard = append(markup.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
				buttons...,
			))
		}
	} else {
		out += "\n" + L("main_no_topics", *Loc)
		markup.InlineKeyboard = make([][]tgbotapi.InlineKeyboardButton, 0)
	}

	return out, markup
}

func getActiveTopics(token string) (topics []Topic, total int) {

	result := API.GetAllTopics(token)
	for _, topic := range result.Topics {
		topics = append(topics, Topic{
			Id:            int(topic.Id),
			Title:         topic.Title,
			TypeId:        int(topic.TypeId),
			StartDatetime: topic.StartDateTime,
			EndDatetime:   topic.EndDateTime,
			Status:        int(topic.Status),
		})
	}
	total = len(topics)
	return topics, total
}

func getTopicQuestions(id int) (questions []Question) {
	db.Find(&questions, Question{TopicId: id})
	return questions
}

func getTopicMessages(id int) []DbMessageWithUser {
	var messages = []DbMessageWithUser{}
	err := db.Raw(`SELECT
  m.id,
    m.body,
    m.question_id,
    m.quote_id,
    u.name
  FROM message m
  JOIN "user" u ON u.id = m.user_id
  JOIN question q ON q.id = m.question_id
  WHERE q.id = ?
  ORDER BY m.id`, id).Scan(&messages).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return messages
}
