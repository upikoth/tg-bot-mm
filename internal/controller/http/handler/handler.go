package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	botmodels "github.com/go-telegram/bot/models"
	"github.com/upikoth/tg-bot-mm/internal/models"
)

func (h *Handler) MainHandler(_ http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error(err.Error())
		return
	}

	h.logger.Info(fmt.Sprintf("Request body: %s", string(body)))

	if len(body) == 0 {
		// Когда функция вызывается с помощью триггера у нее нельзя передать body
		// поэтому это будет признаком, что нужно уведомить пользователя
		// в других случаях функция должна всегда вызываться с body
		notifyErr := h.services.MessageManager.NotifyOnCallUser(r.Context(), h.config.NotificationChatID, 0)

		if notifyErr != nil {
			h.logger.Error(notifyErr.Error())
		}
		return
	}

	var update botmodels.Update

	err = json.Unmarshal(body, &update)
	if err != nil {
		h.logger.Error(err.Error())
		return
	}

	err = h.services.MessageManager.SetCommandsToChat(
		r.Context(),
		update.Message.Chat.ID,
	)

	if err != nil {
		h.logger.Error(err.Error())
		return
	}

	err = h.handleMessage(r.Context(), update.Message)

	if err != nil {
		h.logger.Error(fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	h.logger.Info("Handler успешно выполнен")
}

func (h *Handler) handleMessage(ctx context.Context, message *botmodels.Message) error {
	if !strings.HasPrefix(message.Text, "/") {
		return nil
	}

	commandWithBotAndArgs := strings.TrimPrefix(message.Text, "/")
	parts := strings.Fields(commandWithBotAndArgs)

	commandWithBot := parts[0]
	commandArgs := parts[1:]

	for i := range commandArgs {
		commandArgs[i] = strings.TrimSpace(commandArgs[i])
	}

	commandParts := strings.Split(commandWithBot, "@")
	command := commandParts[0]

	switch command {
	case string(models.CommandOnCallUser):
		return h.services.MessageManager.NotifyOnCallUser(ctx, message.Chat.ID, message.MessageThreadID)
	case string(models.CommandListOfUsers):
		return h.services.MessageManager.SendListOfUsers(ctx, message.Chat.ID, message.MessageThreadID)
	case string(models.CommandAddToListOfUsers):
		return h.services.MessageManager.AddUsers(ctx, commandArgs, message.Chat.ID, message.MessageThreadID)
	case string(models.CommandDeleteFromListOfUsers):
		return h.services.MessageManager.DeleteUsers(ctx, commandArgs, message.Chat.ID, message.MessageThreadID)
	default:
		return h.services.MessageManager.NotifyUnknownCommand(ctx, message.Chat.ID, message.MessageThreadID)
	}
}
