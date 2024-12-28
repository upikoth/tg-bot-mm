package users

import (
	"github.com/upikoth/tg-bot-mm/internal/pkg/logger"
	"github.com/upikoth/tg-bot-mm/internal/repositories/db/users"
)

type Users struct {
	logger    logger.Logger
	repoUsers *users.Users
}

func New(
	logger logger.Logger,
	repoUsers *users.Users,
) *Users {
	return &Users{
		logger:    logger,
		repoUsers: repoUsers,
	}
}
