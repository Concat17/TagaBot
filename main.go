package main

import (
	"TagaBot/database"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	//"TagaBot/database"
)

func MainHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Hi there! I'm TagaBot!"))
}

type executor struct {
	update tgbotapi.Update
}

var isAdding bool
var args []string

func main() {
	http.HandleFunc("/", MainHandler)
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	bot := createBot()
	updates := bot.ListenForWebhook("/" + bot.Token)
	database.ConnectDB()
	monitoring(bot, updates)
}

func createBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI("977213939:AAFbg20C4R3Avg9KhtWY2JrTTijncayLhX8")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	return bot
}

func startGetUpd() tgbotapi.UpdateConfig {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return u
}

func monitoring(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	var msg tgbotapi.MessageConfig

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		if isAdding {
			args = append(args, update.Message.Text)
			switch len := len(args); len {
			case 0:
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Enter article's name")
			case 1:
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Enter article's tag")
			case 2:
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Enter article's comment")
			case 3:
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Enter article's url")
			case 4:
				user := getUserName(update)
				database.AddArticle(user, args[0], args[1], args[2], args[3])
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Article added")
				isAdding = false
				args = args[:0]
			}
		} else { //u'\U000026C4
			if update.Message.IsCommand() {
				msg = execCommnd(update)
			} else {
				switch update.Message.Text {
				case "424": //:poland
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "\U0001F37B")
				case "Лиман":
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "\U0001F1F5\U0001F1F1")
				case "close":
				default:
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				}
			}

		}

		bot.Send(msg)
	}
}

func execCommnd(update tgbotapi.Update) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig
	exec := executor{update}
	switch command := strings.Fields(update.Message.Text)[0]; command {
	case "/greetings":
		msg = exec.greetings()
	case "/bye":
		msg = exec.bye()
	case "/articles":
		msg = exec.showAllNames()
	case "/article":
		msg = exec.showConcrArtclByName()
	case "/test":
		msg = exec.showConcrArtclByName()
	case "/add":
		msg = exec.addArticle()
	default:

	}
	return msg
}

func (exec executor) greetings() tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(exec.update.Message.Chat.ID, fmt.Sprintf("Hello, %v!", exec.update.Message.From.UserName))
	return msg
}

func (exec executor) bye() tgbotapi.MessageConfig {
	text := fmt.Sprintf("Bye-bye! %v", exec.update.Message.From.UserName)
	msg := tgbotapi.NewMessage(exec.update.Message.Chat.ID, text)
	return msg
}

func (exec executor) addArticle() tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig
	isAdding = true
	switch len := len(args); len {
	case 0:
		msg = tgbotapi.NewMessage(exec.update.Message.Chat.ID, "Enter article's name")
	case 1:
		msg = tgbotapi.NewMessage(exec.update.Message.Chat.ID, "Enter article's tag")
	case 2:
		msg = tgbotapi.NewMessage(exec.update.Message.Chat.ID, "Enter article's comment")
	case 3:
		msg = tgbotapi.NewMessage(exec.update.Message.Chat.ID, "Enter article's url")
	default:
		isAdding = false
		args = args[:0]
		user := getUserName(exec.update)
		database.AddArticle(user, args[0], args[1], args[2], args[3])
		msg = tgbotapi.NewMessage(exec.update.Message.Chat.ID, "Article added")
	}
	return msg
}

func (exec executor) showAllNames() tgbotapi.MessageConfig {
	names := database.ShowAllNames()
	msg := tgbotapi.NewMessage(exec.update.Message.Chat.ID, names)
	return msg
}

func (exec executor) showConcrArtclByName() tgbotapi.MessageConfig {
	user := getUserName(exec.update)
	args := commndArgs(exec.update)
	inf := database.ShowConcrByName(user, args[0])
	msg := tgbotapi.NewMessage(exec.update.Message.Chat.ID, inf)
	return msg
}

func commndArgs(update tgbotapi.Update) []string {
	args := update.Message.CommandArguments()
	sepArgs := strings.Split(args, " ")
	return sepArgs
}

func getUserName(update tgbotapi.Update) string {
	return update.Message.From.UserName
}

/*
- func see all articles by name +
- func delete article
- func see concrete article +
*/
// 977213939:AAFbg20C4R3Avg9KhtWY2JrTTijncayLhX8
// https://api.telegram.org/bot977213939:AAFbg20C4R3Avg9KhtWY2JrTTijncayLhX8/setWebhook?url=https://taga-bot.herokuapp.com/977213939:AAFbg20C4R3Avg9KhtWY2JrTTijncayLhX8
