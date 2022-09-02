package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (r *Commander) Help(input *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(input.Chat.ID,
		"/help - help\n"+
			"/list - list product",
	)

	r.bot.Send(msg)
}
