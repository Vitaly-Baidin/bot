package main

import (
	"github.com/Vitaly-Baidin/bot/internal/app/commands"
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

	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		if update.Message != nil { // If we got a message
			switch update.Message.Command() {
			case "help":
				commander.Help(update.Message)
			case "list":
				commander.List(update.Message)
			default:
				commander.Default(update.Message)
			}
		}
	}
}
