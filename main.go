package main

import (
	"TagaBot/database"
	"fmt"
	"log"
	"strings"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
	//"TagaBot/database"
)

type messager struct {
	bot    *tgbot.BotAPI
	chatID int64
}

func main() {
	bot := createBot()
	u := startGetUpd()
	updates, _ := bot.GetUpdatesChan(u)
	monitoring(bot, updates)
}

func createBot() *tgbot.BotAPI {
	bot, err := tgbot.NewBotAPI("977213939:AAFbg20C4R3Avg9KhtWY2JrTTijncayLhX8")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	return bot
}

func startGetUpd() tgbot.UpdateConfig {
	u := tgbot.NewUpdate(0)
	u.Timeout = 60
	return u
}

func monitoring(bot *tgbot.BotAPI, updates tgbot.UpdatesChannel) {
	var msg tgbot.MessageConfig
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if update.Message.IsCommand() {
			msg = execCommnd(update)
		} else {
			msg = tgbot.NewMessage(update.Message.Chat.ID, update.Message.Text)
		}

		bot.Send(msg)
	}
}

func execCommnd(update tgbot.Update) tgbot.MessageConfig {
	var msg tgbot.MessageConfig

	switch command := strings.Fields(update.Message.Text)[0]; command {
	case "/greetings":
		msg = greetings(update)
	case "/bye":
		msg = bye(update)
	case "/test":
		msg = addArticle(update)
	default:

	}

	return msg
}

func greetings(update tgbot.Update) tgbot.MessageConfig {
	msg := tgbot.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Hello, %v!", update.Message.From.UserName))
	return msg
}

func bye(update tgbot.Update) tgbot.MessageConfig {
	text := fmt.Sprintf("Bye-bye! %v", update.Message.From.UserName)
	msg := tgbot.NewMessage(update.Message.Chat.ID, text)
	return msg
}

func addArticle(update tgbot.Update) tgbot.MessageConfig {
	args := update.Message.CommandArguments()
	sepArgs := strings.Split(args, " ")
	database.MakeQuery(sepArgs[0], sepArgs[1], sepArgs[2], sepArgs[3])
	msg := tgbot.NewMessage(update.Message.Chat.ID, "Article added")
	return msg
}

/*
- func see all articles
- func delete article
- func see concrete article
*/
