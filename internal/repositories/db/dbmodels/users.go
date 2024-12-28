package dbmodels

import "github.com/upikoth/tg-bot-mm/internal/models"

type User struct {
	TelegramUsername string `gorm:"primarykey"`
}

func NewYDBUserModel(user *models.User) *User {
	return &User{
		TelegramUsername: user.TelegramUsername,
	}
}

func (u *User) FromYDBModel() *models.User {
	return &models.User{
		TelegramUsername: u.TelegramUsername,
	}
}
