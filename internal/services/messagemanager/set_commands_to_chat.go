package messagemanager

import (
	"context"
	"github.com/go-telegram/bot"
	botmodels "github.com/go-telegram/bot/models"
	"github.com/upikoth/tg-bot-mm/internal/models"
)

var commands = []botmodels.BotCommand{
	{
		Command:     string(models.CommandOnCallUser),
		Description: "Выводит информацию о текущем дежурном",
	},
	{
		Command:     string(models.CommandListOfUsers),
		Description: "Выводит полный список дежурных",
	},
	{
		Command:     string(models.CommandAddToListOfUsers),
		Description: "Добавляет в список дежурных",
	},
	{
		Command:     string(models.CommandDeleteFromListOfUsers),
		Description: "Удаляет из списка дежурных",
	},
}

func (c *MessageManager) SetCommandsToChat(
	ctx context.Context,
	chatID int64,
) error {
	_, err := c.tgBot.SetMyCommands(ctx, &bot.SetMyCommandsParams{
		Commands: commands,
		Scope: &botmodels.BotCommandScopeChatAdministrators{
			ChatID: chatID,
		},
	})

	if err != nil {
		return err
	}

	c.logger.Info("SetCommandsToChat success")
	return nil
}
