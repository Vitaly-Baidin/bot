package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func main() {
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			default:
				defaultBehavior(bot, update.Message)
			}
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, input *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(input.Chat.ID, "/help - help")

	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, input *tgbotapi.Message) {
	log.Printf("[%s] %s", input.From.UserName, input.Text)
	msg := tgbotapi.NewMessage(input.Chat.ID, "You wrote: "+input.Text)
	bot.Send(msg)
}
