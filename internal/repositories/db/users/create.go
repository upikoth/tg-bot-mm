package users

import (
	"context"

	"github.com/pkg/errors"
	"github.com/upikoth/tg-bot-mm/internal/models"
	"github.com/upikoth/tg-bot-mm/internal/repositories/db/dbmodels"
)

func (u *Users) Create(
	ctx context.Context,
	userToCreate *models.User,
) (res *models.User, err error) {
	user := dbmodels.NewYDBUserModel(userToCreate)

	dbRes := u.db.WithContext(ctx).Create(&user)
	createdUser := user.FromYDBModel()

	if dbRes.Error != nil {
		return nil, errors.WithStack(dbRes.Error)
	}

	return createdUser, nil
}
