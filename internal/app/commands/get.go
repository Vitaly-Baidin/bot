package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (r *Commander) Get(input *tgbotapi.Message) {
	var msg tgbotapi.MessageConfig
	args := input.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		msg = tgbotapi.NewMessage(input.Chat.ID, fmt.Sprintf("wrong args %v", args))
		return
	}

	product, err := r.productService.Get(id)
	if err != nil {
		return
	}

	msg = tgbotapi.NewMessage(input.Chat.ID, product.Title)

	r.bot.Send(msg)
}
