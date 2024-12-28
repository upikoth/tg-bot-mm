package users

import (
	"context"

	"github.com/pkg/errors"
	"github.com/upikoth/tg-bot-mm/internal/repositories/db/dbmodels"
)

func (u *Users) DeleteByTelegramUsername(
	ctx context.Context,
	telegramUsername string,
) (err error) {
	dbRes := u.db.
		WithContext(ctx).
		Delete(dbmodels.User{TelegramUsername: telegramUsername})

	if dbRes.Error != nil {
		return errors.WithStack(dbRes.Error)
	}

	return nil
}
