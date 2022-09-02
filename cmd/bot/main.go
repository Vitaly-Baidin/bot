package main

import (
	"github.com/Vitaly-Baidin/bot/internal/service/product"
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

	productService := product.NewService()

	for update := range updates {
		if update.Message != nil { // If we got a message
			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			case "list":
				listCommand(bot, update.Message, productService)
			default:
				defaultBehavior(bot, update.Message)
			}
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, input *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(input.Chat.ID,
		"/help - help\n"+
			"/list - list product",
	)

	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, input *tgbotapi.Message, productService *product.Service) {
	products := productService.List()
	var result string

	for _, p := range products {
		result += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(input.Chat.ID, result)

	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, input *tgbotapi.Message) {
	log.Printf("[%s] %s", input.From.UserName, input.Text)
	msg := tgbotapi.NewMessage(input.Chat.ID, "You wrote: "+input.Text)
	bot.Send(msg)
}
