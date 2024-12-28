package users

import (
	"context"
	"github.com/upikoth/tg-bot-mm/internal/models"
)

func (u *Users) Create(
	ctx context.Context,
	telegramUsername string,
) error {
	_, err := u.repoUsers.Create(ctx, &models.User{
		TelegramUsername: telegramUsername,
	})

	return err
}
