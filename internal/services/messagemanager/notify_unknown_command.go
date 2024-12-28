package messagemanager

import (
	"context"

	"github.com/go-telegram/bot"
)

func (c *MessageManager) NotifyUnknownCommand(
	ctx context.Context,
	chatID int64,
	threadID int,
) error {
	_, err := c.tgBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:          chatID,
		MessageThreadID: threadID,
		Text:            "Команда не найдена, используйте одну из доступных",
	})

	if err != nil {
		return err
	}

	return nil
}
