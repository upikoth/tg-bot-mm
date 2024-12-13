package main

import (
	"app/models"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-telegram/bot"
)

func (s *service) handleMessages() error {
	if s.isNotify {
		log.Println("Уведомление")
		return s.notifyOnCallEngineer()
	}

	if !strings.HasPrefix(s.update.Message.Text, "/") {
		log.Printf("Received unexpected message: %s", s.update.Message.Text)
		return nil
	}

	commandWithBotAndArgs := strings.TrimPrefix(s.update.Message.Text, "/")
	parts := strings.Fields(commandWithBotAndArgs)

	commandWithBot := parts[0]
	commandArgs := parts[1:]

	for i := range commandArgs {
		commandArgs[i] = strings.TrimSpace(commandArgs[i])
	}

	commandParts := strings.Split(commandWithBot, "@")
	command := commandParts[0]

	switch command {
	case string(CommandOnCallEngineer):
		log.Println(CommandOnCallEngineer)
		return s.onCallEngineerHandler()
	case string(CommandListOfEngineers):
		log.Println(CommandListOfEngineers)
		return s.listOfEngineersHandler()
	case string(CommandAddToListOfEngineers):
		log.Println(CommandAddToListOfEngineers)
		return s.addToListOfEngineers(commandArgs)
	case string(CommandDeleteFromListOfEngineers):
		log.Println(CommandDeleteFromListOfEngineers)
		return s.deleteFromListOfEngineers(commandArgs)
	default:
		return s.defaultCommandHandler()
	}
}

func (s *service) onCallEngineerHandler() error {
	onCallEngineer, err := s.getOnCallEngineerTelegramName()

	if err != nil {
		return err
	}

	if onCallEngineer == "" {
		_, err = s.bot.SendMessage(s.r.Context(), &bot.SendMessageParams{
			ChatID:          s.update.Message.Chat.ID,
			MessageThreadID: s.update.Message.MessageThreadID,
			Text:            "Список дежурных пуст, сначала нужно его заполнить",
		})
		if err != nil {
			return err
		}

		return nil
	}

	_, err = s.bot.SendMessage(s.r.Context(), &bot.SendMessageParams{
		ChatID:          s.update.Message.Chat.ID,
		MessageThreadID: s.update.Message.MessageThreadID,
		Text:            fmt.Sprintf("Сегодня дежурит %s", onCallEngineer),
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *service) notifyOnCallEngineer() error {
	onCallEngineer, err := s.getOnCallEngineerTelegramName()

	if err != nil {
		return err
	}

	if onCallEngineer == "" {
		_, err = s.bot.SendMessage(s.r.Context(), &bot.SendMessageParams{
			ChatID: s.cfg.NotificationChatID,
			Text:   "Список дежурных пуст, нужно его заполнить",
		})

		if err != nil {
			return err
		}

		return nil
	}

	_, err = s.bot.SendMessage(s.r.Context(), &bot.SendMessageParams{
		ChatID: s.cfg.NotificationChatID,
		Text:   fmt.Sprintf("Сегодня дежурит %s", onCallEngineer),
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *service) getOnCallEngineerTelegramName() (string, error) {
	engineers, err := s.db.Engineers.GetList(s.r.Context())

	if err != nil {
		return "", err
	}

	if len(engineers) == 0 {
		return "", nil
	}

	tgNames := make([]string, 0, len(engineers))
	for _, e := range engineers {
		tgNames = append(tgNames, e.TelegramUsername)
	}

	onCallEngineer := tgNames[time.Now().YearDay()%len(tgNames)]

	return onCallEngineer, nil
}

func (s *service) listOfEngineersHandler() error {
	engineers, err := s.db.Engineers.GetList(s.r.Context())

	if err != nil {
		return err
	}

	tgNames := make([]string, len(engineers))
	for _, e := range engineers {
		tgNames = append(tgNames, e.TelegramUsername)
	}

	_, err = s.bot.SendMessage(s.r.Context(), &bot.SendMessageParams{
		ChatID:          s.update.Message.Chat.ID,
		MessageThreadID: s.update.Message.MessageThreadID,
		Text:            "Полный список дежурных:" + strings.Join(tgNames, "\n"),
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *service) addToListOfEngineers(args []string) error {
	var telegramUsernames []string

	for _, arg := range args {
		if strings.HasPrefix(arg, "@") {
			telegramUsernames = append(telegramUsernames, arg)
		}
	}

	if len(telegramUsernames) == 0 {
		_, err := s.bot.SendMessage(s.r.Context(), &bot.SendMessageParams{
			ChatID:          s.update.Message.Chat.ID,
			MessageThreadID: s.update.Message.MessageThreadID,
			Text:            "Проверьте корректность введенных данных. Имя пользователя должно начинаться с '@'",
		})

		if err != nil {
			return err
		}

		return nil
	}

	for _, telegramUsername := range telegramUsernames {
		err := s.db.Engineers.Create(s.r.Context(), models.Engineer{
			TelegramUsername: telegramUsername,
		})

		if err != nil {
			_, mErr := s.bot.SendMessage(s.r.Context(), &bot.SendMessageParams{
				ChatID:          s.update.Message.Chat.ID,
				MessageThreadID: s.update.Message.MessageThreadID,
				Text:            "Не удалось добавить пользователей. Обратитесь к администратору или попробуйте позже",
			})

			return mErr
		}
	}

	_, err := s.bot.SendMessage(s.r.Context(), &bot.SendMessageParams{
		ChatID:          s.update.Message.Chat.ID,
		MessageThreadID: s.update.Message.MessageThreadID,
		Text:            "Пользователи добавлены",
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *service) deleteFromListOfEngineers(args []string) error {
	var telegramUsernames []string

	for _, arg := range args {
		if strings.HasPrefix(arg, "@") {
			telegramUsernames = append(telegramUsernames, arg)
		}
	}

	if len(telegramUsernames) == 0 {
		_, err := s.bot.SendMessage(s.r.Context(), &bot.SendMessageParams{
			ChatID:          s.update.Message.Chat.ID,
			MessageThreadID: s.update.Message.MessageThreadID,
			Text:            "Проверьте корректность введенных данных. Имя пользователя должно начинаться с '@'",
		})

		if err != nil {
			return err
		}

		return nil
	}

	for _, telegramUsername := range telegramUsernames {
		err := s.db.Engineers.DeleteByTelegramUsername(s.r.Context(), telegramUsername)

		if err != nil {
			_, mErr := s.bot.SendMessage(s.r.Context(), &bot.SendMessageParams{
				ChatID:          s.update.Message.Chat.ID,
				MessageThreadID: s.update.Message.MessageThreadID,
				Text:            "Не удалось удалить пользователей. Обратитесь к администратору или попробуйте позже",
			})

			return mErr
		}
	}

	_, err := s.bot.SendMessage(s.r.Context(), &bot.SendMessageParams{
		ChatID:          s.update.Message.Chat.ID,
		MessageThreadID: s.update.Message.MessageThreadID,
		Text:            "Пользователи удалены",
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *service) defaultCommandHandler() error {
	_, err := s.bot.SendMessage(s.r.Context(), &bot.SendMessageParams{
		ChatID:          s.update.Message.Chat.ID,
		MessageThreadID: s.update.Message.MessageThreadID,
		Text:            "Команда не найдена, используйте одну из доступных",
	})

	if err != nil {
		return err
	}

	return nil
}
