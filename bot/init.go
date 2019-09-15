package main

import (
	"flag"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	tgbotapi "github.com/temamagic/telegram-bot-api"
	"log"
)

var Config = struct {
	Admin int
	Bot   struct {
		Token string
		Debug bool
	}
	DB struct {
		Name     string
		User     string `default:"root"`
		Password string
		Host     string `default:"localhost"`
		Port     string `default:"3306"`
		Charset  string `default:"utf8mb4"`
	}
}{}

var TopicStatus = &TopicStatusStruct{
	Created:     0,
	InProcess:   1,
	NeedApprove: 2,
	Closed:      3,
}

var (
	flagLocale = flag.String("i18n", "./i18n", "Загрузка локализаций")
	flagConfig = flag.String("config", "config.yml", "Загрузка локализаций")
	locales    []string
	bundle     *i18n.Bundle

	bot *tgbotapi.BotAPI
	err error
	db  *gorm.DB
)

func init() {
	flag.Parse()

	err := configor.Load(&Config, *flagConfig)
	if err != nil {
		log.Panic(err)
	}
	bot, err = tgbotapi.NewBotAPI(Config.Bot.Token)
	if err != nil {
		log.Panic(err)
	}
	db, err = gorm.Open("postgres", "host="+Config.DB.Host+" port="+Config.DB.Port+" user="+Config.DB.User+" dbname="+Config.DB.Name+" password="+Config.DB.Password+" sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(Config.Bot.Debug)
	bot.Debug = Config.Bot.Debug
	log.Printf("Authorized on account %s", bot.Self.UserName)
	langInit()
}
