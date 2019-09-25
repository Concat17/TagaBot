package main

import (
	"TagaBot/database"
	"fmt"
	"log"
	"strings"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
	//"TagaBot/database"
)

type executor struct {
	update tgbot.Update
}

func main() {
	bot := createBot()
	u := startGetUpd()
	updates, _ := bot.GetUpdatesChan(u)
	database.ConnectDB()
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
	exec := executor{update}
	switch command := strings.Fields(update.Message.Text)[0]; command {
	case "/greetings":
		msg = exec.greetings()
	case "/bye":
		msg = exec.bye()
	case "/test":
		msg = exec.showConcrArtclByName()
	default:

	}
	return msg
}

func (exec executor) greetings() tgbot.MessageConfig {
	msg := tgbot.NewMessage(exec.update.Message.Chat.ID, fmt.Sprintf("Hello, %v!", exec.update.Message.From.UserName))
	return msg
}

func (exec executor) bye() tgbot.MessageConfig {
	text := fmt.Sprintf("Bye-bye! %v", exec.update.Message.From.UserName)
	msg := tgbot.NewMessage(exec.update.Message.Chat.ID, text)
	return msg
}

func (exec executor) addArticle() tgbot.MessageConfig {
	args := commndArgs(exec.update)
	if len(args) < 4 { // this is needs rewriting. poor error handling
		msg := tgbot.NewMessage(exec.update.Message.Chat.ID, "Not enough args for adding articles.")
		return msg
	}
	database.AddArticle(args[0], args[1], args[2], args[3])
	msg := tgbot.NewMessage(exec.update.Message.Chat.ID, "Article added")
	return msg
}

func (exec executor) showAllNames() tgbot.MessageConfig {
	names := database.ShowAllNames()
	msg := tgbot.NewMessage(exec.update.Message.Chat.ID, names)
	return msg
}

func (exec executor) showConcrArtclByName() tgbot.MessageConfig {
	args := commndArgs(exec.update)
	inf := database.ShowConcrByName(args[0])
	msg := tgbot.NewMessage(exec.update.Message.Chat.ID, inf)
	return msg
}

func commndArgs(update tgbot.Update) []string {
	args := update.Message.CommandArguments()
	sepArgs := strings.Split(args, " ")
	return sepArgs
}

/*
- func see all articles by name +
- func delete article
- func see concrete article +
*/
