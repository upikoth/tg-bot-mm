package main

import (
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Command string

const CommandOnCallEngineer Command = "oncallengineer"
const CommandListOfEngineers Command = "list"
const CommandAddToListOfEngineers Command = "add"
const CommandDeleteFromListOfEngineers Command = "delete"

func (s *service) setCommands() error {
	commands := []models.BotCommand{
		{
			Command:     string(CommandOnCallEngineer),
			Description: "Выводит информацию о текущем дежурном",
		},
		{
			Command:     string(CommandListOfEngineers),
			Description: "Выводит полный список дежурных",
		},
		{
			Command:     string(CommandAddToListOfEngineers),
			Description: "Добавляет в список дежурных",
		},
		{
			Command:     string(CommandDeleteFromListOfEngineers),
			Description: "Удаляет из списка дежурных",
		},
	}

	_, err := s.bot.SetMyCommands(s.r.Context(), &bot.SetMyCommandsParams{
		Commands: commands,
		Scope: &models.BotCommandScopeChatAdministrators{
			ChatID: s.update.Message.Chat.ID,
		},
	})

	if err != nil {
		return err
	}

	log.Println("setCommands success")
	return nil
}
