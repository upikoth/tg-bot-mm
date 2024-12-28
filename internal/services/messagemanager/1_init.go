package messagemanager

import (
	"github.com/go-telegram/bot"
	"github.com/upikoth/tg-bot-mm/internal/pkg/logger"
	"github.com/upikoth/tg-bot-mm/internal/services/users"
)

type MessageManager struct {
	logger logger.Logger
	tgBot  *bot.Bot
	users  *users.Users
}

func New(
	logger logger.Logger,
	tgBot *bot.Bot,
	users *users.Users,
) *MessageManager {
	return &MessageManager{
		logger: logger,
		tgBot:  tgBot,
		users:  users,
	}
}
