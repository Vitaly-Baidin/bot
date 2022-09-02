package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (r *Commander) List(input *tgbotapi.Message) {
	products := r.productService.List()
	var result string

	for _, p := range products {
		result += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(input.Chat.ID, result)

	r.bot.Send(msg)
}
