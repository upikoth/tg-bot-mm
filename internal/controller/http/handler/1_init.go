package handler

import (
	"github.com/upikoth/tg-bot-mm/internal/config"
	"github.com/upikoth/tg-bot-mm/internal/pkg/logger"
	"github.com/upikoth/tg-bot-mm/internal/services"
)

type Handler struct {
	logger   logger.Logger
	config   *config.Config
	services *services.Services
}

func New(
	logger logger.Logger,
	config *config.Config,
	services *services.Services,
) *Handler {
	return &Handler{
		logger:   logger,
		config:   config,
		services: services,
	}
}
