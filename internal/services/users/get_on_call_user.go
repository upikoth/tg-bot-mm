package users

import (
	"context"
	"time"

	"github.com/upikoth/tg-bot-mm/internal/constants"
	"github.com/upikoth/tg-bot-mm/internal/models"
)

func (u *Users) GetOnCallUser(
	ctx context.Context,
) (*models.User, error) {
	users, err := u.repoUsers.GetList(ctx)

	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, constants.ErrEntityNotFound
	}

	onCallUser := users[time.Now().YearDay()%len(users)]

	return &onCallUser, nil
}
