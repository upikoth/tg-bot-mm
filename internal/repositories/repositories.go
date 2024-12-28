package repositories

import (
	"github.com/upikoth/tg-bot-mm/internal/config"
	"github.com/upikoth/tg-bot-mm/internal/pkg/logger"
	"github.com/upikoth/tg-bot-mm/internal/repositories/db"
)

type Repository struct {
	DB *db.DB
}

func New(
	log logger.Logger,
	cfg *config.Config,
) (*Repository, error) {
	dbInstance, err := db.New(log, cfg)

	if err != nil {
		return nil, err
	}

	return &Repository{
		DB: dbInstance,
	}, nil
}

func (r *Repository) Connect() error {
	return r.DB.Connect()
}

func (r *Repository) Disconnect() error {
	return r.DB.Disconnect()
}
