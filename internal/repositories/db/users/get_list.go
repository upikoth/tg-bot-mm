package users

import (
	"context"

	"github.com/pkg/errors"
	"github.com/upikoth/tg-bot-mm/internal/constants"
	"github.com/upikoth/tg-bot-mm/internal/models"
	"github.com/upikoth/tg-bot-mm/internal/repositories/db/dbmodels"
)

func (u *Users) GetList(
	ctx context.Context,
) (res []models.User, err error) {
	var users []dbmodels.User

	dbRes := u.db.
		WithContext(ctx).
		Find(&users)

	if dbRes.Error != nil {
		return nil, errors.WithStack(dbRes.Error)
	}

	if len(users) == 0 {
		return nil, constants.ErrEntityNotFound
	}

	var resUsers []models.User
	for _, user := range users {
		resUsers = append(resUsers, *user.FromYDBModel())
	}

	return resUsers, nil
}
