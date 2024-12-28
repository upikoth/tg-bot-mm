package messagemanager

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/pkg/errors"
	"github.com/upikoth/tg-bot-mm/internal/constants"
)

func (c *MessageManager) NotifyOnCallUser(
	ctx context.Context,
	chatID int64,
	threadID int,
) error {
	onCallUser, err := c.users.GetOnCallUser(ctx)

	if errors.Is(err, constants.ErrEntityNotFound) {
		_, err = c.tgBot.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:          chatID,
			MessageThreadID: threadID,
			Text:            "Список дежурных пуст, сначала нужно его заполнить",
		})

		return err
	}

	if err != nil {
		return err
	}

	_, err = c.tgBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:          chatID,
		MessageThreadID: threadID,
		Text:            fmt.Sprintf("Сегодня дежурит %s", onCallUser.TelegramUsername),
	})

	return err
}
