package users

import (
	"context"

	"github.com/upikoth/tg-bot-mm/internal/models"
)

func (u *Users) GetList(
	ctx context.Context,
) ([]models.User, error) {
	return u.repoUsers.GetList(ctx)
}
