package main

import (
	"TagaBot/database"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	//"TagaBot/database"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("977213939:AAFbg20C4R3Avg9KhtWY2JrTTijncayLhX8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		if update.Message.IsCommand() {

			switch command := strings.Fields(update.Message.Text)[0]; command {
			case "/greetings":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Hello, %v!", update.Message.From.UserName))
			case "/bye":
				mes := fmt.Sprintf("Bye-bye! %v", update.Message.From.UserName)
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, mes)
			default:
				args := update.Message.CommandArguments()
				sepArgs := strings.Split(args, " ")
				database.MakeQuery(sepArgs[0], sepArgs[1], sepArgs[2], sepArgs[3])
				//editArgs := strings.Join(sepArgs, "~")

				//msg = tgbotapi.NewMessage(update.Message.Chat.ID, args)
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Article added")
			}

		}

		bot.Send(msg)
	}
}
