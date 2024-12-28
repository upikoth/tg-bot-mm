package messagemanager

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/pkg/errors"
	"github.com/upikoth/tg-bot-mm/internal/constants"
	"strings"
)

func (c *MessageManager) SendListOfUsers(
	ctx context.Context,
	chatID int64,
	threadID int,
) error {
	users, err := c.users.GetList(ctx)

	if errors.Is(err, constants.ErrEntityNotFound) {
		_, err = c.tgBot.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:          chatID,
			MessageThreadID: threadID,
			Text:            "Список дежурных пуст",
		})

		return nil
	}

	if err != nil {
		return err
	}

	tgNames := make([]string, 0, len(users))
	for _, user := range users {
		tgNames = append(tgNames, user.TelegramUsername)
	}

	_, err = c.tgBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:          chatID,
		MessageThreadID: threadID,
		Text:            "Полный список дежурных:\n\n" + strings.Join(tgNames, "\n"),
	})

	if err != nil {
		return err
	}

	return err
}
