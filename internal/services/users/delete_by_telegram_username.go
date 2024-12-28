package users

import (
	"context"
)

func (u *Users) DeleteByTelegramUsername(
	ctx context.Context,
	telegramUsername string,
) error {
	return u.repoUsers.DeleteByTelegramUsername(ctx, telegramUsername)
}
