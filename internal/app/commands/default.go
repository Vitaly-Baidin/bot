package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (r *Commander) Default(input *tgbotapi.Message) {
	log.Printf("[%s] %s", input.From.UserName, input.Text)
	msg := tgbotapi.NewMessage(input.Chat.ID, "You wrote: "+input.Text)
	r.bot.Send(msg)
}
