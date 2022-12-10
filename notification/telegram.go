package notification

import (
	"log"
	"net/http"
	"time"

	"github.com/williamneok/urlcheckertoolcli/config"
)

var app *config.AppConfig

func NewNotificationConfig(a *config.AppConfig) {
	app = a
}

//TelegramBotMsg takes url status and push it to telegram bot server

func TelegramBotMsg(msg string) {

	botToken := app.BotToken
	telegramChannel := app.TelegramChannel
	downTime := time.Now().Format(time.Stamp)

	_, err := http.PostForm("https://api.telegram.org/bot"+botToken+"/sendMessage?chat_id="+telegramChannel+"&text=Status: Down! - "+msg+" Timestamp: "+downTime, nil)
	if err != nil {
		log.Println(err)
	}
}
