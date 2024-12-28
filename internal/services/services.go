package services

import (
	"github.com/go-telegram/bot"
	"github.com/upikoth/tg-bot-mm/internal/config"
	"github.com/upikoth/tg-bot-mm/internal/pkg/logger"
	"github.com/upikoth/tg-bot-mm/internal/repositories"
	"github.com/upikoth/tg-bot-mm/internal/services/messagemanager"
	"github.com/upikoth/tg-bot-mm/internal/services/users"
)

type Services struct {
	Users          *users.Users
	MessageManager *messagemanager.MessageManager
}

func New(
	log logger.Logger,
	cfg *config.Config,
	repo *repositories.Repository,
) (*Services, error) {
	srvs := &Services{}

	tgBot, err := bot.New(cfg.BotToken)
	if err != nil {
		return nil, err
	}

	srvs.Users = users.New(
		log,
		repo.DB.Users,
	)

	srvs.MessageManager = messagemanager.New(
		log,
		tgBot,
		srvs.Users,
	)

	return srvs, nil
}
