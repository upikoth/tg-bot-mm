package messagemanager

import (
	"context"
	"strings"

	botmodels "github.com/go-telegram/bot"
)

func (c *MessageManager) AddUsers(
	ctx context.Context,
	tgNames []string,
	chatID int64,
	threadID int,
) error {
	var validTgNames []string

	for _, arg := range tgNames {
		if strings.HasPrefix(arg, "@") {
			validTgNames = append(validTgNames, arg)
		}
	}

	if len(validTgNames) == 0 {
		_, err := c.tgBot.SendMessage(ctx, &botmodels.SendMessageParams{
			ChatID:          chatID,
			MessageThreadID: threadID,
			Text:            "Проверьте корректность введенных данных. Имя пользователя должно начинаться с '@'",
		})

		if err != nil {
			return err
		}

		return nil
	}

	for _, tgName := range validTgNames {
		err := c.users.Create(ctx, tgName)

		if err != nil {
			_, mErr := c.tgBot.SendMessage(ctx, &botmodels.SendMessageParams{
				ChatID:          chatID,
				MessageThreadID: threadID,
				Text:            "Не удалось добавить пользователей. Обратитесь к администратору или попробуйте позже",
			})

			return mErr
		}
	}

	_, err := c.tgBot.SendMessage(ctx, &botmodels.SendMessageParams{
		ChatID:          chatID,
		MessageThreadID: threadID,
		Text:            "Пользователи добавлены",
	})

	return err
}
